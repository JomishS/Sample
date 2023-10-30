package Routes

import (
	"example/Project3/internal/controllers"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouteMapper interface {
	SetupRoutes()
	GetHandler() http.Handler
}

type RouteMapperDummy struct{}

func (r *RouteMapperDummy) SetupRoutes() {}

func (r *RouteMapperDummy) GetHandler() http.Handler {
	return http.Server{}.Handler
}

type routeMapper struct {
	engine         *gin.Engine
	userController controllers.Usercontroller
	docController  controllers.Documentcontroller
}


func NewRouteMapper(
	userController controllers.Usercontroller,
	docController controllers.Documentcontroller,
) RouteMapper {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} 
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.AllowCredentials = true

	router.Use(cors.New(config))
	return &routeMapper{
		engine:         router,
		userController: userController,
		docController:  docController,
	}
}

func (r *routeMapper) SetupRoutes() {

	r.SetupDocRoutes()
	r.SetupUserRoutes()

}
func (r *routeMapper) GetHandler() http.Handler {
	return r.engine
}
