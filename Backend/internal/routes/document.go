package Routes


func (r *routeMapper) SetupDocRoutes(){
	docRouter:=r.engine.Group("/documents")
	docController:=r.docController

	docRouter.GET("",docController.GetAllDoc)
	docRouter.GET("/:id",docController.GetByIdDoc)

	docRouter.POST("",docController.CreateDoc)

	docRouter.PUT("/:id",docController.UpdateDoc)

	docRouter.DELETE("/:id",docController.DeleteDoc)

}