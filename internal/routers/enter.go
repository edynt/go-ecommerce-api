package routers

import (
	"github.com/edynt/go-ecommerce-api/internal/routers/manager"
	"github.com/edynt/go-ecommerce-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
