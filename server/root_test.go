package server

import (
	"embed"
	"os"
	"testing"

	"github.com/g41797/sputnik/sidecar"
)

//go:embed _configuration
var embconf embed.FS

func TestMain(m *testing.M) {
	ecode := testMain(m)
	os.Exit(ecode)
}

func testMain(m *testing.M) int {
	cleanUp, _ := sidecar.UseEmbeddedConfiguration(&embconf)
	defer cleanUp()

	stop, err := sidecar.StartServices()
	if err != nil {
		return -1
	}
	defer stop()

	ecode := m.Run()
	return ecode
}
