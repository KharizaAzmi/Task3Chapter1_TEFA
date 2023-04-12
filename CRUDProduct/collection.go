package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for storing data
type Product struct {
	IDProduct int `json:"idproduct"`
	ProductTitle string `json:"producttitle"`
	Category string `json:"category"`
	Variant string `json:"variant"`
	Price float64 `json:"price"`
	Stock  int    `json:"stock"`
	Description string `json:"description"`
}

var productCollection = db().Database("belajar").Collection("product") // get collection "users" from db() which returns *mongo.Client


func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var data Product
	err := json.NewDecoder(r.Body).Decode(&data) // storing in product variable of type user
	if err != nil {
		fmt.Print(err)
	}

	ctx := context.Background()

	product := Product{
        IDProduct: 1,
        ProductTitle: "Sosis",
        Category: "Food",
		Variant: "Chicken",
		Price: 20000,
		Stock: 14,
		Description: "Sosis Rasa Ayama kemasan 200 gram",
    }

	insertResult, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the mongodb ID of generated document

}

// Get Profile of a particular product by ID

func getProductProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body Product
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := productCollection.FindOne(context.TODO(), bson.D{{Key: "IDProduct", Value: body.IDProduct}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

//Update Profile of User
func updateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		IDProduct int `bson:"idproduct"`
		Stock  int    `bson:"stock"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{Key: "idproduct", Value: body.IDProduct}} // converting value to BSON type
	after := options.After                // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "stock", Value: body.Stock}}}}
	updateResult := productCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of User

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := productCollection.DeleteOne(context.TODO(), bson.D{{Key: "IDProduct", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents
	cur, err := productCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}