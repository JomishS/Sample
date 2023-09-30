package inject

import (
	"example/Project3/database"
	"example/Project3/internal/controllers"
	"example/Project3/internal/repository"
	 "example/Project3/internal/routes"
 	"example/Project3/internal/service"
	"fmt"
	//  "gorm.io/gorm"
)

func GetRouteMapper() Routes.RouteMapper {
	// con := database.ConnectFunc()
	// db := con.GetDB()
	db := database.GetDB()

	fmt.Println("inside GetRouteMapper")
	fmt.Println(db)
	user := repository.UserRepFunc(db)
	serviceUser := Service.UserServFunc(user)
	controllerUser := controllers.UserConFunc(serviceUser)

	document := repository.DocumentRepFunc(db)
	serviceDoc := Service.DocumentServFunc(document)
	controllerDoc := controllers.DocumentConFunc(serviceDoc)

	routeMapper := Routes.NewRouteMapper(controllerUser, controllerDoc)
	return routeMapper
}
