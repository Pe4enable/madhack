package services

import (
	"log"
	"time"

	"github.com/madhack/config"

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

//func (r *ExchangeReader) readSingle(cache chan ratestates.RateState) {
//	bitfinexRates, err := bitfinex.Get(r.conf.BitfinexTickers)
//	log.Print("bitfinex", bitfinexRates)
//
//	if err != nil {
//		log.Printf("BITFINEX ERROR: %s", err)
//		err = nil
//	}
//
//	coinmarketcapRates, err := coinmarketcap.GetUSDs(r.conf.Coinmarketcap)
//	log.Print("coinmarketcap", coinmarketcapRates)
//	if err != nil {
//		log.Printf("COINMARKET ERROR: %s", err)
//		err = nil
//	}
//
//	rates, err := ratestates.MergeSlice([]ratestates.RateState{bitfinexRates, coinmarketcapRates})
//	if err != nil {
//		log.Printf("MERGE ERROR: %s", err)
//		return
//	}
//
//	allExists := rates.AllExist(r.conf.Must...)
//	if !allExists {
//		log.Printf("ERROR: NOT ALL EXISTS: %s", rates)
//		log.Fatalf("Must be % v", r.conf.Must)
//		return
//	}
//
//	convertedRates, err := rates.RateStatesByConvertedToBase(r.conf.Base)
//	if err != nil {
//		log.Fatalf("FATAL ERROR: CANNOT CONVERT %s TO BASE %s: ", rates, r.conf.Base, err)
//		return
//	}
//
//	// log.Printf("R: %s", rates)
//	// log.Printf("C: %s", convertedRates)
//	convertedRates = convertedRates.OnlyTickers(r.conf.Must)
//
//	// For debug purposes
//	convertedRates.AddTimestampDate(
//		r.conf.TimeOffset.Years,
//		r.conf.TimeOffset.Months,
//		r.conf.TimeOffset.Days,
//	)
//
//	cache <- convertedRates
//}

func (r *ExchangeReader) Start(cache chan ratestates.RateState) {
	for {
		//r.readSingle(cache)
		time.Sleep(time.Duration(r.conf.Interval) * time.Second)
	}
}
