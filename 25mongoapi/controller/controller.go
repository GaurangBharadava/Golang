package controller

import (
	"context"
	"encoding/json"
	
	"fmt"
	"log"
	"net/http"

	"github.com/GaurangBharadava/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://gaurang:gaurang@myatlasclusteredu.m9e3tef.mongodb.net/?retryWrites=true&w=majority&appName=myAtlasClusterEDU"
const dbName = "netflix"
const colName = "watchlist"

// most imp
var collection *mongo.Collection

// connect with mongoDB
func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption) //TODO is type of context we also can use Background()

	if err != nil {
		// panic(err)
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance is ready")
}

//mongo helpers -file

//insert 1 record
func insertOneMovie(movie model.Netflix) { //this method is in smaller letter because it will not exported
	inserted, err := collection.InsertOne(context.Background(),movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted one movie with id",inserted.InsertedID)
}

//update one record
func updateOneMovie(movieId string) {
	id,err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id":id} // we can use bson.D also if order of element matters.
	update := bson.M{"$set":bson.M{"watched":true}}

	result,err := collection.UpdateOne(context.Background(),filter,update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modifiesd count:",result.ModifiedCount)
}

//deletion of the record
func deleteOneMovie(movieId string) {
	id,_ := primitive.ObjectIDFromHex(movieId)

	filter :=  bson.M{"_id":id}
	deleteCount,err := collection.DeleteOne(context.Background(),filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie delete count:",deleteCount)
}

func deleteAllTheMovie() {
	filter := bson.D{{}}
	delete, err := collection.DeleteMany(context.Background(),filter,nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleterd:",delete.DeletedCount)
}

//read all movies
func readAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies,movie)
		
	}
	defer cursor.Close(context.Background())

	return movies
}

// actual controller

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	allMovies := readAllMovies() 

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	deleteAllTheMovie()
}