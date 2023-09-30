package Routes

import (

	// "example/Project3/internal/inject"
	_ "github.com/lib/pq"
)

func (r *routeMapper) SetupUserRoutes(){
	userRouter:=r.engine.Group("/users")
	userController:=r.userController

	userRouter.GET("",userController.GetAllUser)
	userRouter.GET("/:id",userController.GetByIdUser)

	userRouter.POST("",userController.CreateUser)

	userRouter.PUT("/:id",userController.UpdateUser)

	userRouter.DELETE("/:id",userController.DeleteUser)

}