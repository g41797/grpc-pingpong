package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"embed"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/g41797/sputnik/sidecar"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// Ensure running tests in ready environment
func PrepareTestEnvironment(embconf *embed.FS) (cleanFunc func(), err error) {

	cleaner := new(sidecar.Cleaner)

	// Create configuration and env files from embedded ones
	cleanUp, err := sidecar.UseEmbeddedConfiguration(embconf)
	if err != nil {
		return cleaner.Clean, err
	}

	cleaner.Push(cleanUp)

	// Setup environment variables from env files
	// Replaces vscode setting in launch.json file
	// Reason - running tests in github without vscode
	// see url of docker-compose.yml in test.env
	err = sidecar.LoadEnv()
	if err != nil {
		return cleaner.Clean, err
	}

	// Get url of docker-compose.yml from environment
	prefix := "TEST_"
	base := "COMPOSEURL"
	url := os.Getenv(prefix + base)
	if len(url) == 0 {
		url = os.Getenv(base)
	}

	// Download  docker-compose.yml to config folder
	err = sidecar.LoadDockerComposeFile(url)
	if err != nil {
		return cleaner.Clean, err
	}

	// Start broker and additional servers using docker compose file
	stop, err := sidecar.StartServices()
	if err != nil {
		return cleaner.Clean, err
	}
	cleaner.Push(stop)

	// Connect to broker - ensure running tests in ready environment
	conn := tryConnect()
	cleaner.Push(func() { conn.Close() })

	return cleaner.Clean, nil

}

func startServerConnectClient(server *grpc.Server) (conn *grpc.ClientConn, err error) {
	ctx := context.Background()

	conn, err = grpc.DialContext(ctx, "",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(dialer(server)),
		grpc.WithBlock())

	return conn, err
}

func dialer(server *grpc.Server) func(context.Context, string) (net.Conn, error) {

	listener := bufconn.Listen(1024 * 1024)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func tryConnect() *nats.Conn {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			conn, err := connect()
			if err == nil {
				ticker.Stop()
				return conn
			}
		}
	}
}

type CONFTEST struct {
	VERSION                        string
	MEMPHIS_HOST                   string
	MEMPHIS_PORT                   int
	MEMPHIS_CLIENT                 string
	USER_PASS_BASED_AUTH           bool
	ROOT_PASSWORD                  string
	ROOT_USER                      string
	CONNECTION_TOKEN               string
	CLIENT_CERT_PATH               string
	CLIENT_KEY_PATH                string
	ROOT_CA_PATH                   string
	JWT_EXPIRES_IN_MINUTES         int
	REFRESH_JWT_EXPIRES_IN_MINUTES int
	REST_GW_UPDATES_SUBJ           string
	JWT_SECRET                     string
	REFRESH_JWT_SECRET             string
	HTTP_PORT                      string
	DEV_ENV                        string
	DEBUG                          bool
	CLOUD_ENV                      bool
}

func getconf() CONFTEST {
	var cncnf CONFTEST
	cf, _ := sidecar.ConfFolder()
	sidecar.ConfigFactory(cf)("connector", &cncnf)
	return cncnf
}

func connect() (*nats.Conn, error) {
	configuration := getconf()
	creds := configuration.CONNECTION_TOKEN
	username := configuration.ROOT_USER
	if configuration.USER_PASS_BASED_AUTH {
		username = "$$memphis"
		creds = configuration.CONNECTION_TOKEN + "_" + configuration.ROOT_PASSWORD
		if !configuration.CLOUD_ENV {
			creds = configuration.ROOT_PASSWORD
		}
	}

	return connect2(configuration.MEMPHIS_HOST, username, creds, &configuration)
}

func connect2(hostname, username, creds string, configuration *CONFTEST) (*nats.Conn, error) {
	var nc *nats.Conn
	var err error

	natsOpts := nats.Options{
		Url:            hostname + ":" + strconv.Itoa(configuration.MEMPHIS_PORT),
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  3 * time.Second,
		Name:           configuration.MEMPHIS_CLIENT,
	}

	if configuration.USER_PASS_BASED_AUTH {
		natsOpts.Password = creds
		natsOpts.User = username
	} else {
		natsOpts.Token = username + "::" + creds
	}

	if configuration.CLIENT_CERT_PATH != "" && configuration.CLIENT_KEY_PATH != "" && configuration.ROOT_CA_PATH != "" {
		cert, err := tls.LoadX509KeyPair(configuration.CLIENT_CERT_PATH, configuration.CLIENT_KEY_PATH)
		if err != nil {
			return nil, err
		}
		cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
		if err != nil {
			return nil, err
		}
		TLSConfig := &tls.Config{MinVersion: tls.VersionTLS12}
		TLSConfig.Certificates = []tls.Certificate{cert}
		certs := x509.NewCertPool()

		pemData, err := os.ReadFile(configuration.ROOT_CA_PATH)
		if err != nil {
			return nil, err
		}
		certs.AppendCertsFromPEM(pemData)
		TLSConfig.RootCAs = certs
		natsOpts.TLSConfig = TLSConfig
	}

	nc, err = natsOpts.Connect()
	if err != nil {
		return nil, err
	}

	return nc, nil
}
