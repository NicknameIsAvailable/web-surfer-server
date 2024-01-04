package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"webSurfer/db"
	"webSurfer/routers"
)

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func checkConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{Message: "Hello", Success: true}

	json.NewEncoder(w).Encode(response)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURI := os.Getenv("MONGODB_URI")
	if err := db.ConnectDB(dbURI); err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	result, err := db.GetDocument("mycollection", bson.D{{"key", "value"}})
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return
	}

	fmt.Print(result)

	router := gin.Default()

	api := router.Group("/api")
	{
		routers.SetupUsersRouter(api)
	}

	router.Run(":8080")
}
