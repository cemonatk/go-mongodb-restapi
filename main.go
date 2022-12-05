package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Id      uint8  `json:"_id" bson:"_id"`
	Name    string `json:"name" bson:"name" binding:"required"`
	Address string `json:"address" bson:"address" binding:"required"`
	Age     int    `json:"age" bson:"age" binding:"required"`
	Gender  string `json:"gender" bson:"gender"`
}

func initDB() {
	// db.users.deleteMany({})

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("userdb").Collection("users")
}

func getUsers(c *gin.Context) {
	findOptions := options.Find()

	var results []*User
	cursor, err := collection.Find(context.Background(), findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.Background()) {
		var u User
		err := cursor.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &u)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.Background())
	if len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{"Result": "No results found"})
	}

	c.JSON(http.StatusOK, results)
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	var user []User

	findOptions := options.Find()

	searchResult, err := collection.Find(context.Background(), bson.M{"name": name}, findOptions)

	defer searchResult.Close(context.Background())

	if err = searchResult.All(context.Background(), &user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": ""})
		return
	}
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid Format"})
		return
	}

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "Invalid Format"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Value Inserted": user})
}

func removeUser(c *gin.Context) {
	name := c.Param("name")

	var deletedUser User

	err := collection.FindOneAndDelete(context.Background(), bson.M{"name": name}).Decode(&deletedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatal(err)
			c.JSON(http.StatusNotModified, gin.H{"ERROR": "ErrNoDocuments."})
			return

		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "Internal error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Deleted User": deletedUser})
}

func main() {

	initDB()

	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:name", getUser)
	router.POST("/users/create", createUser)
	router.DELETE("/users/:name/delete", removeUser)
	router.Run("localhost:8080")
}
