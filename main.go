package main

import (
"fmt"
"log"
"net/http"
"os"
mgo "gopkg.in/mgo.v2"
)

type App struct {
    Session *mgo.Session
    Server string
    Database string
    Collection string
}

var a App

// MAIN.GO
func main(){
	fmt.Printf("Running...\n")
    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        // os.Getenv("APP_DB_NAME"),
        "mattertime", //todo pas as config
        "time",//todo pas as config
    )
    a.Connect()
    a.Run(":8000")    
}

func (a *App) Initialize(user, password, dbname string, collectionname string) {
	a.Database = dbname
	a.Collection = collectionname	
}

func (a *App) Run(addr string) { 
	router := NewRouter()
	log.Fatal(http.ListenAndServe(addr, router))
}

func (a *App) Connect() {	
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	a.Session = session
	fmt.Printf("Connected...\n")
}