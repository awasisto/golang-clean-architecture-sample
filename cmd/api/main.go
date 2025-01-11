package main

import (
	"go-clean-architecture-sample/api/controllers"
	"go-clean-architecture-sample/api/routers"
	"go-clean-architecture-sample/infrastructure/avatarprovider"
	"go-clean-architecture-sample/infrastructure/database"
	"net/http"
)

func main() {
	employeeRepository := database.NewInMemoryEmployeeDatabase()
	avatarProvider := avatarprovider.NewGithubAvatarProvider()
	employeeController := controllers.NewEmployeeController(employeeRepository, avatarProvider)
	employeeRouter := routers.NewEmployeeRouter(*employeeController)
	http.ListenAndServe("0.0.0.0:8080", &employeeRouter.MuxRouter)
}
