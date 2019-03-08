package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rs/cors"
)

//Configuration used for global config varabiles
type Configuration struct {
	bucket         string
	pacTable       string
	renderJobTable string
	fpModelTable   string
}

var config = Configuration{}
var sess, _ = session.NewSession(&aws.Config{
	Region: aws.String("us-west-2"),
})

func main() {
	env := os.Getenv("ENV")
	if env == `dev` {
		config = devConfig
	} else {
		config = prodConfig
	}

	router := filmPacRouter()

	fmt.Println(`server listening on port :80`, config)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":80", handler))
}

var devConfig = Configuration{
	bucket:         "empty",
	pacTable:       "empty",
	renderJobTable: "empty",
	fpModelTable:   "empty",
}
var prodConfig = Configuration{
	bucket:         "empty",
	pacTable:       "empty",
	renderJobTable: "empty",
}
