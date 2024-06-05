package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"testBetStarters/internal/db"
	"testBetStarters/internal/routes"
)

func main() {
	runMigrations := flag.Bool("migrate", false, "run migrations")
	runSeeds := flag.Bool("seed", false, "run seeds")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
	}

	DB, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	if *runMigrations {
		if err = db.Migrate(DB); err != nil {
			log.Fatal(err)
		}
	}

	if *runSeeds {
		if err = db.Seed(DB); err != nil {
			log.Fatal(err)
		}
	}

	e := echo.New()
	e.Use(middleware.Logger())

	routes.InitRoutes(e, DB)

	e.Logger.Fatal(e.Start(":8080"))
}
