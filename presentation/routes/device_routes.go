package routes

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

func ApplyDeviceHealthCheckRoutes(router *gin.Engine) {
	configData, err := os.ReadFile("env.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err)
	}

	ctx := context.TODO()
	option := options.Client().ApplyURI(config.App.MongoDB.URI)

	client, err := mongo.Connect(ctx, option)
	if err != nil {
		panic(err)
	}

	devicesCollection := client.Database("FLOWECH_DEVICES_DB").Collection("devices")

	deviceRoutes := router.Group("/device")
	{
		deviceRoutes.GET("/list", func(c *gin.Context) {
			var deviceList []bson.M
			cur, err := devicesCollection.Find(context.TODO(), bson.D{{}})
			if err != nil {
				panic(err)
			}
			defer cur.Close(context.TODO())
			for cur.Next(context.TODO()) {
				var result bson.M
				err := cur.Decode(&result)
				if err != nil {
					panic(err)
				}
				deviceList = append(deviceList, result)
			}

			c.JSON(200, gin.H{
				"deviceList": deviceList,
			})
		})

		deviceRoutes.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")

			var deviceById bson.M
			filter := bson.D{{"id", id}}
			err = devicesCollection.FindOne(context.TODO(), filter).Decode(&deviceById)
			if err == mongo.ErrNoDocuments {
				fmt.Println("No documents found")
				return
			}
			if err != nil {
				panic(err)
			}

			c.JSON(200, gin.H{
				"device": deviceById,
			})
		})
	}
}
