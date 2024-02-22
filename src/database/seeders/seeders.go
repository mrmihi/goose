package seeders

//
//import (
//	"context"
//	"fmt"
//	"goose/src/database"
//	"goose/src/modules/restaurants/api/v1/models"
//	"log"
//	"time"
//)
//
//func seeder() {
//	Connect()
//
//	defer func() {
//		if err := client.Disconnect(context.Background()); err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	// Access the database
//	db := client.Database(database.DatabaseName)
//
//	// Access the collection
//	collection := db.Collection("restaurants")
//
//	// Drop existing collection
//	err = collection.Drop(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Define restaurant data to be seeded
//	restaurants := []models.Restaurant{
//		{
//			Name:       "Restaurant 1",
//			Email:      "restaurant1@example.com",
//			Password:   "password123",
//			APIKey:     "apikey123",
//			LocationID: "location123",
//			CreatedAt:  time.Now().Format(time.RFC3339),
//			UpdatedAt:  time.Now().Format(time.RFC3339),
//		},
//		{
//			Name:       "Restaurant 2",
//			Email:      "restaurant2@example.com",
//			Password:   "password456",
//			APIKey:     "apikey456",
//			LocationID: "location456",
//			CreatedAt:  time.Now().Format(time.RFC3339),
//			UpdatedAt:  time.Now().Format(time.RFC3339),
//		},
//	}
//
//	// Insert restaurant data into the collection
//	for _, restaurant := range restaurants {
//		_, err := collection.InsertOne(context.Background(), restaurant)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	fmt.Println("Seed data inserted successfully!")
//}
