package repositories

import (
	"context"
	"github.com/BankEx/madhack/config"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/BankEx/madhack/models"
)

type MongoRepository struct {
	conf       config.Mongo
	client *mongo.Client
	db *mongo.Database
	Items *mongo.Collection
	//session    *mgo.Session
	//collection *mgo.Collection
}

func New(config config.Mongo, ctx context.Context) (repository *MongoRepository, err error) {
	repository = new(MongoRepository)
	repository.conf = config

	repository.client, err = mongo.NewClient(repository.conf.Url)
	if err != nil {
		log.Fatal(err)
	}

	err = repository.client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer repository.client.Disconnect(nil)

	repository.db = repository.client.Database(repository.conf.DB)
	repository.Items = repository.db.Collection(repository.conf.Collection)

	return repository, nil
}


func (r *MongoRepository) AddItem(item models.Item) (err error){
	_, err = r.Items.InsertOne(context.Background(), item)
	return
}


func (r *MongoRepository) GetAllItems() (result []models.Item, err error){
	cursor, err := r.Items.Find(context.Background(), nil,nil...)
	if err == nil {
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			item := models.Item{}
			err := cursor.Decode(&item)
			if err == nil {
				result = append(result, item)
			}
		}

		if err := cursor.Err(); err != nil {
			return nil, err
		}
	}
	return result, nil
}