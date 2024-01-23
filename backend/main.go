package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func getDocument(c *gin.Context) {
	query := `{ 
	"query": { 
		"term": { 
			"Carrier": {
					"value":"ES-Air"
				} 
			} 
		} 
	}`

	res, err := client.Search(
		client.Search.WithIndex("kibana_sample_data_flights"),
		client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	c.IndentedJSON(http.StatusOK, res)
}

func getHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ok")
}

var client *elasticsearch.Client

func main() {
	cert, _ := os.ReadFile("./http_ca.crt")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
			"https://localhost:9201",
		},
		Username: "elastic",
		Password: "doIOriK3vhMP19rMMuTH",
		CACert:   cert,
	}

	// https://github.com/elastic/go-elasticsearch/blob/main/_examples/main.go
	var err error
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("elasticsearch.NewClient: %v", err)
	}

	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	var r map[string]interface{}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	router := gin.Default()
	router.GET("/hello-world", getHelloWorld)
	router.GET("/doc", getDocument)

	router.Run("localhost:8080")
}
