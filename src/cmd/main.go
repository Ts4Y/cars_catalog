package main

import (
	"cars_catalog/external"
	"cars_catalog/internal/repository/httpclient"
	"cars_catalog/internal/repository/postgresql"
	"cars_catalog/internal/repository/transaction"
	"cars_catalog/internal/usecase"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "cars_catalog/cmd/docs"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lmittmann/tint"
)

// @title Cars Catalog API
// @version 1.0
// @description API server for Cars Catalog
// @contact.name API Support

// @host localhost:8080
// @BasePath /api/v1

func main() {

	if err := godotenv.Load("vars.env"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(os.Getenv("PG_URL"))

	log := slog.New(tint.NewHandler(os.Stderr, nil))

	db, err := sqlx.Connect("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Error(err.Error())
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err.Error())
		return
	}

	peopleRepo := postgresql.NewPeopleRepo()
	carsRepo := postgresql.NewCarsRepo()
	httpClient := httpclient.NewHttpClientRepo()

	sm := transaction.NewSQLSessionManager(db)

	peopleUsecase := usecase.NewPeopleUsecase(log, sm, peopleRepo)
	carsUsecase := usecase.NewCarsUsecase(log, sm, carsRepo, httpClient)

	server := external.New(peopleUsecase, carsUsecase)
	server.Run()

}
