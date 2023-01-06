package api

import (
	"github.com/dattranman/todo/app"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	Root  *gin.Engine
	APIv1 *gin.RouterGroup

	Tasks *gin.RouterGroup
}

type API struct {
	App         *app.App
	BaseRouters *Routers
}

func Init(app *app.App, root *gin.Engine) *API {
	api := &API{
		App:         app,
		BaseRouters: &Routers{},
	}

	api.BaseRouters.Root = root
	api.BaseRouters.APIv1 = api.BaseRouters.Root.Group("/api/v1")
	api.BaseRouters.Tasks = api.BaseRouters.APIv1.Group("/tasks")

	api.HealthCheck()
	api.InitTask()
	return api
}

func (a *API) Run() error {
	err := a.BaseRouters.Root.Run(a.App.Config.ServiceSettings.Port)
	if err != nil {
		return err
	}
	return nil
}
