package server

import (
	"github.com/g41797/sputnik/sidecar"
	"github.com/memphisdev/memphis.go"
)

type BrokerConnConf struct {
	BROKER_HOST  string
	BROKER_USER  string
	BROKER_PSWRD string
}

type brokerConnector struct {
	cnf *BrokerConnConf
}

func newBrokerConnector() (bc *brokerConnector, err error) {
	bc = new(brokerConnector)

	bc.cnf, err = getConf()
	if err != nil {
		return nil, err
	}

	return bc, nil
}

func (bc *brokerConnector) connect() (*memphis.Conn, error) {

	mc, err := memphis.Connect(bc.cnf.BROKER_HOST, bc.cnf.BROKER_USER, memphis.Password(bc.cnf.BROKER_PSWRD))

	return mc, err
}

func getConf() (*BrokerConnConf, error) {
	cf, err := sidecar.ConfFolder()

	if err != nil {
		return nil, err
	}

	result := new(BrokerConnConf)

	err = sidecar.ConfigFactory(cf)("brokerconn", result)

	return result, err
}
