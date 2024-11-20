package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/whatsauth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var IteungIPAddress string = os.Getenv("ITEUNGBEV1")

var MongoString string = os.Getenv("MONGOSTRING")

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "petapedia",
}

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)

var Usertables = [4]whatsauth.LoginInfo{mhs, dosen, user, user1}

var mhs = whatsauth.LoginInfo{
	Userid:   "MhswID",
	Password: "Password",
	Phone:    "Telepon",
	Username: "Login",
	Uuid:     "simak_mst_mahasiswa",
	Login:    "2md5",
}

var dosen = whatsauth.LoginInfo{
	Userid:   "NIDN",
	Password: "Password",
	Phone:    "Handphone",
	Username: "Login",
	Uuid:     "simak_mst_dosen",
	Login:    "2md5",
}

var user = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "phone",
	Username: "user_name",
	Uuid:     "simak_besan_users",
	Login:    "2md5",
}

var user1 = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "user_phone",
	Username: "user_name",
	Uuid:     "besan_users",
	Login:    "2md5",
}

func InitMongoDBConnection() {
	// Example of MongoDB connection options
	clientOptions := options.Client().ApplyURI(MongoString).
		SetMaxPoolSize(50).               // Set max pool size to 50
		SetMinPoolSize(10).               // Set min pool size to 10
		SetConnectTimeout(10 * time.Second) // Timeout for connection

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check if the connection is established
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	// Assign to global MongoDB connection variable
	Ulbimongoconn = client.Database("petapedia")

	// Create a geospatial index on the "geometry" field for the "jalan" collection
	roadsCollection := Ulbimongoconn.Collection("jalan")
	_, err = roadsCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{{Key: "geometry", Value: "2dsphere"}}, // Create a 2dsphere index
	})
	if err != nil {
		log.Fatalf("Error creating geospatial index: %v", err)
	} else {
		log.Println("Geospatial index on 'geometry' field created successfully.")
	}
}
