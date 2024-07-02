package ticket

import (
	"backend-hex-arch/pkg/ports"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	collection *mongo.Collection
}

func NewRepository() ports.TicketRepository {
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetBSONOptions(bsonOpts).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
	return &repository{
		collection: client.Database("hexdeez").Collection("tickets"),
	}
}
