package services

import (
	"time"

	"github.com/BankEx/madhack/config"

	"github.com/BankEx/bankex-tokensale-common/ratestates"
)

type ExchangeReader struct {
	conf *config.Config
}

func New(conf *config.Config) (service *ExchangeReader, err error) {
	service = new(ExchangeReader)
	service.conf = conf
	//cli, err := rpcclient.New(service.conf.BTCConn, nil)
	//service.btcClient = cli

	if err != nil {
		return nil, err
	}
	return service, nil
}



func (r *ExchangeReader) Start(cache chan ratestates.RateState) {
	for {
		//r.readSingle(cache)
		time.Sleep(time.Duration(r.conf.Interval) * time.Second)
	}
}
