package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"vkHackaton/internal/cfg"
	"vkHackaton/internal/controllers"
)

func main() {

	setUpDatabase()
	e := echo.New()
	setUpRoutes(e)
	e.Logger.Fatal(e.Start(":8081"))

}

func setUpDatabase() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Couldn't load env file: %s", err)
	}

	dbUsername := os.Getenv("db_username")
	dbPassword := os.Getenv("db_password")
	//cfgString := dbUsername + ":" + dbPassword + "@tcp(docker.for.mac.localhost:3306)/nft_tickets_app"
	cfgString := dbUsername + ":" + dbPassword + "@/nft_tickets_app"
	cfg.CFG.DB, err = sql.Open("mysql", cfgString)
	if err != nil {
		log.Fatal("Couldn't connect to the database")
	}
}

func setUpRoutes(e *echo.Echo) {
	e.GET("/getevent/:title", controllers.GetEvent)
	e.GET("/getallevents", controllers.GetAllEvents)
	e.POST("/createevent", controllers.SaveEvent)
	e.PUT("/edit_event", controllers.EditEvent)
	e.POST("/registerallet", controllers.RegisterWaller)
	e.GET("/getwallet", controllers.GetWallet)
}
