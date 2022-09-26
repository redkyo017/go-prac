package main

func initializeRoutes() {

	// Handle the index route
	router.Use(setUserStatus())
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)

		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
	}
	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:article_id", getArticle)

		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)

		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
}
