package api

import (
	"net/http"

	"github.com/dattranman/todo/model/request"
	"github.com/dattranman/todo/model/response"
	"github.com/gin-gonic/gin"
)

func (api *API) InitTask() {
	api.BaseRouters.Tasks.POST("", api.createTask)
	api.BaseRouters.Tasks.GET("", api.getListTask)
	api.BaseRouters.Tasks.GET("/:id", api.getTaskByID)
	api.BaseRouters.Tasks.PUT("/:id", api.updateTask)
	api.BaseRouters.Tasks.DELETE("/:id", api.deleteTaskByID)
}

func (api *API) createTask(c *gin.Context) {
	var task request.Task
	err := c.ShouldBindJSON(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := api.App.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (api *API) getTaskByID(c *gin.Context) {
	var req request.IDParam
	err := req.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := api.App.GetTaskByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (api *API) getListTask(c *gin.Context) {
	var task request.Task
	err := c.ShouldBindJSON(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := api.App.GetListTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (api *API) updateTask(c *gin.Context) {
	var task request.Task
	err := c.ShouldBindJSON(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := api.App.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (api *API) deleteTaskByID(c *gin.Context) {
	var req request.IDParam
	err := req.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.UtilResponse{
			Status: response.StatusSuccess,
			Msg:    err.Error(),
		})
		return
	}

	resp, err := api.App.DeleteTaskByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
