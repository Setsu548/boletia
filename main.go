package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error to load .env file")
	}

	elapsed, _ := strconv.Atoi(os.Getenv("ELAPSED"))

	ticker := time.NewTicker(time.Duration(elapsed) * time.Minute)

	// DB Configuration

	// echo
	e := echo.New()

	// routes

	e.Start(":8000")
}
