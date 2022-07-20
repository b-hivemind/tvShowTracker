package db

import (
	"context"
	"fmt"
	"log"
	"time"

	api "github.com/b-hivemind/preparer/pkg/tvmazeapi"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DatabaseName = "tvShowTrackerLive"

const (
	URI            = "mongodb://database:27017"
	collectionName = "allShows"
)

func connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ListDbs() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connect()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if databases, err := client.ListDatabaseNames(ctx, bson.M{}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(databases)
	}
}

func InsertShow(show api.Show) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connect()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
		return false
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Check if show episodes are populated
	if show.Episodes.All == nil {
		// try to populate shows
		show.PopulateEpisodes()
		if show.Episodes.All == nil {
			log.Printf("No episodes detected for %v", show)
			return false
		}
		log.Printf("Populated Show: %v", show)
	}

	coll := client.Database(DatabaseName).Collection(collectionName)
	if data, err := bson.Marshal(show); err != nil {
		log.Fatal(err)
		return false
	} else {
		result, err := coll.InsertOne(context.TODO(), data)
		if err != nil {
			log.Fatal(err)
			return false
		}
		list_id := result.InsertedID
		fmt.Println(list_id)
		return true
	}
}

func GetShowFromID(showID int) (api.Show, error) {
	filter := bson.D{bson.E{Key: "_id", Value: showID}}
	var showObj api.Show
	// var result bson.D

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connect()

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
		return showObj, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	coll := client.Database(DatabaseName).Collection(collectionName)
	err := coll.FindOne(context.TODO(), filter).Decode(&showObj)
	if err != nil {
		return showObj, err
	}
	return showObj, nil
}

func GetAllShows() ([]api.Show, error) {
	filter := bson.D{}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "name", Value: 1}, {Key: "image", Value: 1}, {Key: "status", Value: 1}}
	opts := options.Find().SetProjection(projection)
	var showObjs []api.Show
	// var result bson.D

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connect()

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
		return showObjs, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	coll := client.Database(DatabaseName).Collection(collectionName)
	cursor, err := coll.Find(context.TODO(), filter, opts)
	cursor.All(context.TODO(), &showObjs)

	if err != nil {
		return showObjs, err
	}
	return showObjs, nil
}

func SetEpisodes(showID int, episodes map[int]bool) error {
	filter := bson.D{bson.E{Key: "_id", Value: showID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "episodes.all", Value: episodes}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedDoc bson.D

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connect()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	coll := client.Database(DatabaseName).Collection(collectionName)
	err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// fmt.Println(updatedDoc)
	return nil
}
