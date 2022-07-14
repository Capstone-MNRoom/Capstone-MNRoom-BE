package main

import (
	"be9/mnroom/config"
	"be9/mnroom/factory"
	_middlewares "be9/mnroom/middlewares"
	"be9/mnroom/migration"
	"be9/mnroom/routes"
)

func main() {
	dbConn := config.InitDB()
	migration.InitMigrate(dbConn)
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	_middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
