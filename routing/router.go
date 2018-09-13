package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/controllers"
	"github.com/kamalraimi/recruiter-api/middlewares"
)

func Routes(route *gin.Engine) {

	//route.Use(middlewares.CORSMiddleware()) // CORS

	route.POST("/recruiter-api/login", middlewares.GinJwtMiddlewareHandler().LoginHandler)

	route.GET("/recruiter-api/menu/:id", controllers.GetMenusByUser)

	roleRoutes := route.Group("/recruiter-api/roles")

	{

		roleRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetRoles)

		roleRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetRole)
		roleRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostRole)
		roleRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutRole)
		roleRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteRole)
	}

	userRoutes := route.Group("/recruiter-api/users")
	{
		//userRoutes.GET("/", middlewares.AuthMiddleware(), controllers.GetUsers)
		userRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetUsers)
		userRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetUser)
		userRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostUser)
		userRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutUser)
		userRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteUser)
	}

	menuRoutes := route.Group("/recruiter-api/menus")
	{
		menuRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetMenus)
		menuRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetMenu)
		menuRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostMenu)
		menuRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutMenu)
		menuRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteMenu)
	}

	customerRoutes := route.Group("/recruiter-api/customers")
	{
		customerRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetCustomers)
		customerRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetCustomer)
		customerRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostCustomer)
		customerRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutCustomer)
		customerRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteCustomer)
	}

	collaboraterRoutes := route.Group("/recruiter-api/collaboraters")
	{
		collaboraterRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetCollaboraters)
		collaboraterRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetCollaborater)
		collaboraterRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostCollaborater)
		collaboraterRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutCollaborater)
		collaboraterRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteCollaborater)
	}

	positionRoutes := route.Group("/recruiter-api/positions")
	{
		positionRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetPositions)
		positionRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetPosition)
		positionRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostPosition)
		positionRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutPosition)
		positionRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeletePosition)
	}

	applicationRoutes := route.Group("/recruiter-api/applications")
	{
		applicationRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetApplications)
		applicationRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetApplication)
		applicationRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostApplication)
		applicationRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutApplication)
		applicationRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteApplication)
	}

	relocationRoutes := route.Group("/recruiter-api/relocations")
	{
		relocationRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetRelocations)
		relocationRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetRelocation)
		relocationRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostRelocation)
		relocationRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutRelocation)
		relocationRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteRelocation)
	}

	immigrationRoutes := route.Group("/recruiter-api/immigrations")
	{
		immigrationRoutes.GET("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetImmigrations)
		immigrationRoutes.GET("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.GetImmigration)
		immigrationRoutes.POST("/", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PostImmigration)
		immigrationRoutes.PUT("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.PutImmigration)
		immigrationRoutes.DELETE("/:id", middlewares.GinJwtMiddlewareHandler().MiddlewareFunc(), controllers.DeleteImmigration)
	}

}
