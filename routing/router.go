package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/controllers"
)

func Routes(route *gin.Engine) {

	roleRoutes := route.Group("/roles")
	{
		roleRoutes.GET("/", controllers.GetRoles)
		roleRoutes.GET("/:id", controllers.GetRole)
		roleRoutes.POST("/", controllers.PostRole)
		roleRoutes.PUT("/:id", controllers.PutRole)
		roleRoutes.DELETE("/:id", controllers.DeleteRole)
	}

	userRoutes := route.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.POST("/", controllers.PostUser)
		userRoutes.PUT("/:id", controllers.PutUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	menuRoutes := route.Group("/menus")
	{
		menuRoutes.GET("/", controllers.GetMenus)
		menuRoutes.GET("/:id", controllers.GetMenu)
		menuRoutes.POST("/", controllers.PostMenu)
		menuRoutes.PUT("/:id", controllers.PutMenu)
		menuRoutes.DELETE("/:id", controllers.DeleteMenu)
	}

	customerRoutes := route.Group("/customers")
	{
		customerRoutes.GET("/", controllers.GetCustomers)
		customerRoutes.GET("/:id", controllers.GetCustomer)
		customerRoutes.POST("/", controllers.PostCustomer)
		customerRoutes.PUT("/:id", controllers.PutCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomer)
	}

	collaboraterRoutes := route.Group("/collaboraters")
	{
		collaboraterRoutes.GET("/", controllers.GetCollaboraters)
		collaboraterRoutes.GET("/:id", controllers.GetCollaborater)
		collaboraterRoutes.POST("/", controllers.PostCollaborater)
		collaboraterRoutes.PUT("/:id", controllers.PutCollaborater)
		collaboraterRoutes.DELETE("/:id", controllers.DeleteCollaborater)
	}

	positionRoutes := route.Group("/positions")
	{
		positionRoutes.GET("/", controllers.GetPositions)
		positionRoutes.GET("/:id", controllers.GetPosition)
		positionRoutes.POST("/", controllers.PostPosition)
		positionRoutes.PUT("/:id", controllers.PutPosition)
		positionRoutes.DELETE("/:id", controllers.DeletePosition)
	}

	applicationRoutes := route.Group("/applications")
	{
		applicationRoutes.GET("/", controllers.GetApplications)
		applicationRoutes.GET("/:id", controllers.GetApplication)
		applicationRoutes.POST("/", controllers.PostApplication)
		applicationRoutes.PUT("/:id", controllers.PutApplication)
		applicationRoutes.DELETE("/:id", controllers.DeleteApplication)
	}

	relocationRoutes := route.Group("/relocations")
	{
		relocationRoutes.GET("/", controllers.GetRelocations)
		relocationRoutes.GET("/:id", controllers.GetRelocation)
		relocationRoutes.POST("/", controllers.PostRelocation)
		relocationRoutes.PUT("/:id", controllers.PutRelocation)
		relocationRoutes.DELETE("/:id", controllers.DeleteRelocation)
	}

	immigrationRoutes := route.Group("/immigrations")
	{
		immigrationRoutes.GET("/", controllers.GetImmigrations)
		immigrationRoutes.GET("/:id", controllers.GetImmigration)
		immigrationRoutes.POST("/", controllers.PostImmigration)
		immigrationRoutes.PUT("/:id", controllers.PutImmigration)
		immigrationRoutes.DELETE("/:id", controllers.DeleteImmigration)
	}

}
