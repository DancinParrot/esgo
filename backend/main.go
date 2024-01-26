package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
)

const INDEX = "kibana_sample_data_flights"

var client *elasticsearch.TypedClient

type AllFields struct {
	Fields []string `json:"fields"`
}

func getAllFields(c *gin.Context) {
	res, err := client.Indices.GetMapping().
		Index(INDEX).
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	fields := res[INDEX].Mappings.Properties
	keys := make([]string, 0, len(fields))

	for key := range fields {
		keys = append(keys, key)
	}

	c.IndentedJSON(http.StatusOK, AllFields{
		Fields: keys,
	})
}

func getAllDocuments(c *gin.Context) {
	res, err := client.Search().
		Index(INDEX).
		Request(
			&search.Request{
				Query: &types.Query{
					MatchAll: &types.MatchAllQuery{},
				},
			},
		).
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	c.IndentedJSON(http.StatusOK, res)
}

// func getDocument(c *gin.Context) {
// 	query := `{
// 	"query": {
// 		"match": {
// 				"Carrier": "ES-Air"
// 			}
// 		}
// 	}`
//
// 	res, err := client.Search(
// 		client.Search.WithIndex("kibana_sample_data_flights"),
// 		client.Search.WithBody(strings.NewReader(query)),
// 	)
// 	if err != nil {
// 		log.Fatalf("Error getting response: %s", err)
// 	}
// 	defer res.Body.Close()
// 	// Check response status
// 	if res.IsError() {
// 		log.Fatalf("Error: %s", res.String())
// 	}
//
// 	var r map[string]interface{}
// 	// Deserialize the response into a map.
// 	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
// 		log.Fatalf("Error parsing the response body: %s", err)
// 	}
//
// 	c.IndentedJSON(http.StatusOK, r)
// }

func getHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ok")
}

func corsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	cert, _ := os.ReadFile("./ca.crt")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
			"https://localhost:9201",
		},
		Username: "elastic",
		Password: "changeme",
		CACert:   cert,
	}

	// https://github.com/elastic/go-elasticsearch/blob/main/_examples/main.go
	var err error
	client, err = elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("elasticsearch.NewClient: %v", err)
	}

	// res, err := client.Info()
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }
	// defer res.Body.Close()
	// // Check response status
	// if res.IsError() {
	// 	log.Fatalf("Error: %s", res.String())
	// }
	//
	// var r map[string]interface{}
	// // Deserialize the response into a map.
	// if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	// 	log.Fatalf("Error parsing the response body: %s", err)
	// }
	// // Print client and server version numbers.
	// log.Printf("Client: %s", elasticsearch.Version)
	// log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	// log.Println(strings.Repeat("~", 37))

	router := gin.Default()
	router.Use(corsConfig())
	router.GET("/hello-world", getHelloWorld)
	// router.GET("/doc", getDocument)
	// router.GET("/docs", getAllDocuments)
	router.GET("/fields", getAllFields)

	router.Run("localhost:8080")
}
