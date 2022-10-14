package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EloToJaa/cyoa"
)

func main() {
	port := flag.Int("port", 5000, "the port to start the server on")
	fileName := flag.String("file", "gopher.json", "JSON file for Choose Your Own Adventure story")
	flag.Parse()

	fmt.Printf("Using the story from file: %s\n", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	handler := cyoa.NewHandler(story)
	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(addr, handler))
}
