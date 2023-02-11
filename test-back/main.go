package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"draconia/ent"
)

// func GetTest(client *ent.Client) ([]*ent.User, error) {
// 	// client.User.Create().SetLast()
// 	users, err := client.User.Query().All(context.Background())
// 	if err != nil {
// 		return nil, fmt.Errorf("nique ta mere")
// 	}

// 	return users, nil
// }

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=game password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
