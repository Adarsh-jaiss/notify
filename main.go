package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	listenAddr := flag.String("addr", ":3000", "HTTP server address")
	flag.Parse()
	
	http.HandleFunc("/notifications", makeApifunc(FetchNotificationsRepeated))
	fmt.Println("Starting server on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
	
}
