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
			"firstName":       p.FirstName,
			"middleName":      p.MiddleName,
			"lastName":        p.LastName,
			"teacherId":       p.TeacherID,
			"department":      p.Department,
			"ratingCount":     p.RatingsCount,
			"ratingClass":     p.RatingClass,
			"overallRating":   p.OverallRating,
			"institutionName": p.InstitutionName,
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
	}
	collection = getCollection(fmt.Sprintf("%s.%d", "professors", schoolId))

	result, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{{"lastName", "text"}, {"firstName", "text"}},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	return result.InsertedIDs

}

func FindProfessor(schoolId int, name string) (professor rate_my_professor.Professor) {
	collection := getCollection(fmt.Sprintf("%s.%d", "professors", schoolId))
	find := collection.FindOne(context.Background(), bson.M{
		"$text": bson.M{
			"$search": name,
		},
	})
	if find.Err() != nil {
		log.Fatalln(find.Err())
	}

	err := find.Decode(&professor)
	if err != nil {
		log.Fatalln(err)
	}
	return professor
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
