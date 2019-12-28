package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"./controllers"
	"./handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres-dev"
	password = "dev"
	dbname   = "dev"
)

func main() {
	fmt.Println("Start app")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Cannot connected!")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connected!")
		panic(err)
	}
	fmt.Println("Successfully connected!")

	env := &handler.Env{
		DB:   db,
		Port: string(port),
		Host: host,
	}

	router := mux.NewRouter()
	router.Handle("/register", handler.Handler{env, controllers.RegisterHandler}).Methods("POST")
	router.Handle("/login", handler.Handler{env, controllers.LoginHandler}).Methods("POST")

	apiRouter := router.PathPrefix("/api").Subrouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	apiRouter.Handle("/getweather", handler.Handler{env, controllers.GetWeather}).Methods("POST")
	apiRouter.Use(controllers.Middleware)

	handlWrapper := c.Handler(router)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":4000", handlWrapper))
}

// //https://blog.questionable.services/article/http-handler-error-handling-revisited/
