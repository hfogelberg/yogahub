package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	mongoDbHost := os.Getenv("YOGA_DB_HOST")
	fmt.Printf("Attempting to connect to %s\n", mongoDbHost)
	session, err := mgo.Dial(mongoDbHost)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB %s %s\n", mongoDbHost, err.Error())
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	connection := Connection{session.DB(MongoDb)}
	fmt.Println("Connected to Db")

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/admin", adminHandler).Methods("GET")
	router.HandleFunc("/admin/poses", connection.posesHandler).Methods("GET")
	router.HandleFunc("/admin/poses/create", createPoseHandler).Methods("GET")
	router.HandleFunc("/admin/poses/create", connection.postCreatePoseHandler).Methods("POST")
	router.HandleFunc("/favicon.ico", faviconHandler)

	mux := http.NewServeMux()
	mux.Handle("/", router)

	static := http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	router.PathPrefix("/public").Handler(static)

	n := negroni.Classic()
	n.UseHandler(mux)
	port := ":" + getEnv("PORT", "80")

	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(port, n)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
