package router

func (gbt *GbtServer) PrepareRoutest() {
	authenticatedRoute := gbt.router.Group("/")
	authenticatedRoute.Use(AuthorizationMiddleWare())
	authenticatedRoute.GET("/hello", gbt.HelloWorld)
	authenticatedRoute.POST("/ask", gbt.AskQuestions)
	gbt.router.POST("/login", gbt.Login)

}
