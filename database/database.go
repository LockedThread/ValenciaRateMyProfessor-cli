package database

import (
	"cli/model/rate_my_professor"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var (
	Client mongo.Client
)

func Connect() *mongo.Client {
	fmt.Printf("Connecting to database\n")
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Connected to database\n")
	return client
}

func InsertScrapeData(schoolId int, professors rate_my_professor.Professors) []interface{} {
	collection := getCollection(fmt.Sprintf("%s.%d", "professors", schoolId))
	var documents []interface{}

	for _, professor := range professors.Professors {
		p := *professor
		documents = append(documents, bson.M{
			"firstName":     p.FirstName,
			"middleName":    p.MiddleName,
			"lastName":      p.LastName,
			"teacherId":     p.TeacherID,
			"department":    p.Department,
			"ratingCount":   p.RatingsCount,
			"ratingClass":   p.RatingClass,
			"overallRating": p.OverallRating,
		})
	}

	count, err := collection.EstimatedDocumentCount(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	if count > 0 {
		err = collection.Drop(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		_, err = collection.Indexes().DropAll(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
	}
	_, err = collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: "firstName",
	})
	if err != nil {
		log.Fatalln(err)
	}
	_, err = collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: "lastName",
	})
	if err != nil {
		log.Fatalln(err)
	}

	result, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		log.Fatalln(err)
	}
	return result.InsertedIDs

}

// Utility methods
func getDatabase(database string) *mongo.Database {
	return Client.Database(database)
}

func getDatabaseFromDefault() *mongo.Database {
	return getDatabase("valencia-rate-my-professor")
}

func getCollection(collection string) mongo.Collection {
	return *getDatabaseFromDefault().Collection(collection)
}
