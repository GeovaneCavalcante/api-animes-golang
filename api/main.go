package main

import (
	"fmt"
	"github.com/GeovaneCavalcante/api-anime-golang/infrastructure/repository/nosql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/GeovaneCavalcante/api-anime-golang/api/handler"
	"github.com/GeovaneCavalcante/api-anime-golang/usecase/anime"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	connectionDatabase, err := nosql.ConnectMongoDB()

	if err != nil{
		fmt.Println("Error connect to database")
	}

	animeRepo := nosql.NewNoSqlRepo(connectionDatabase)

	animeService := anime.NewService(animeRepo)

	handler.MakeBookHandlers(r, *n, animeService)

	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(8000),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
