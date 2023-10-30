package inject

import (
	"example/Project3/database"
	"example/Project3/internal/controllers"
	"example/Project3/internal/repository"
	 "example/Project3/internal/routes"
 	"example/Project3/internal/service"
)

func GetRouteMapper() Routes.RouteMapper {

	db := database.GetDB()
	user := repository.UserRepFunc(db)
	serviceUser := Service.UserServFunc(user)
	controllerUser := controllers.UserConFunc(serviceUser)

	document := repository.DocumentRepFunc(db)
	serviceDoc := Service.DocumentServFunc(document)
	controllerDoc := controllers.DocumentConFunc(serviceDoc)

	routeMapper := Routes.NewRouteMapper(controllerUser, controllerDoc)
	return routeMapper
}
