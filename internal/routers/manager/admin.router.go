package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (pr *UserRouter) InitAdminRouter(Router *gin.RouterGroup) {

	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
		adminRouterPublic.POST("/change-password")
	}
}
