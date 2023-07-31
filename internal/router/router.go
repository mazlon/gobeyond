package router

func (gbt *GbtServer) PrepareRoutest() {
	gbt.router.GET("/hello", gbt.HelloWorld)
	gbt.router.POST("/ask", gbt.AskQuestions)
	gbt.router.POST("/login", gbt.Login)

}
