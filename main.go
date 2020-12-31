package main

import (
	"calendar-backend/graph"
	"calendar-backend/graph/generated"
	"net/http"
	"os"
	"log"

	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

type Schedule struct {
  gorm.Model
	UserId int
	StartAt time.Time
	EndAt time.Time
}

var (
  db  *gorm.DB
  err error
)

func gormConnectTest() () {
	dsn := "root:pass@tcp(mysql:3306)/calendar?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	schedule := Schedule{UserId: 1, StartAt: time.Now(), EndAt: time.Now()}

	result := db.Create(&schedule)

	log.Println(result.Error)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// gormConnectTest()
}
