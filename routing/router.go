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

}
