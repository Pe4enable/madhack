package repositories

import (
	"strconv"
	"github.com/madhack/config"
)

type MongoRepository struct {
	conf       config.Mongo
	//session    *mgo.Session
	//collection *mgo.Collection
}

func New(config config.Mongo) (repository *MongoRepository, err error) {
	repository = new(MongoRepository)
	repository.conf = config

	//repository.session = mgoconn.DialLoop(repository.conf.Url)
	//mgoconn.TimeoutReconnectFunc = mgoconn.ReconnectFuncForSessions([]*mgo.Session{repository.session})
	//
	//db := repository.session.DB(repository.conf.DB)
	//repository.collection = db.C(repository.conf.Collection)

	return
}

//func (r *MongoRepository) Start(cache chan ratestates.RateState) (err error) {
//	//err = ratestates.MongoWriter(cache, r.collection)
//	return
//}

//func (r *MongoRepository) GetRate(from string, to string, date string) (result string, err error) {
//	var raw map[string]interface{}
//
//	findParams := bson.M{
//		"timestamp": bson.M{
//			"$lte": date,
//		},
//	}
//
//	err = mgoconn.DoResistant(
//		func() error {
//		    if date != "" {
//			    return r.collection.Find(findParams).Sort("-timestamp").One(&raw)
//			}
//			return  r.collection.Find(nil).Sort("-timestamp").One(&raw)
//		},
//	)
//	if err != nil {
//		return "", err
//	}
//
//	var rateState, er = ratestates.Unmap(raw)
//	if er != nil {
//		return "", er
//	}
//
//	fromRate := rateState.Rates()[from]
//	toRate := rateState.Rates()[to]
//	resultRate := toRate / fromRate
//	result = strconv.FormatFloat(resultRate, 'f', -1, 64)
//	return
//}
//
//func (r *MongoRepository) GetAllRates(from string, date string) (result map[string]float64, err error) {
//	var raw map[string]interface{}
//
//   // isoData:= ISODate(date)
//	findParams := bson.M{
//		//"timestamp": {$gte: isoDate, $lt: isoDate}}
//	}
//
//	err = mgoconn.DoResistant(
//		func() error {
//		    if date != "" {
//			    return r.collection.Find(findParams).Sort("-timestamp").One(&raw)
//			}
//			return  r.collection.Find(nil).Sort("-timestamp").One(&raw)
//		},
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	var rateState, er = ratestates.Unmap(raw)
//	if er != nil {
//		return nil, er
//	}
//
//	fromRate := rateState.Rates()[from]
//
//    result = make(map[string]float64)
//	for key := range rateState.Rates() {
//        result[key] = rateState.Rates()[key]/fromRate
//     }
//
//
//	return
//}
