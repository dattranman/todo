package app

import (
	"github.com/dattranman/todo/model/request"
	"github.com/dattranman/todo/model/response"
)

func (a *App) CreateTask(task request.Task) (resp response.Task, err error) {
	taskData, err := task.ParseToSchema()
	if err != nil {
		return
	}
	err = a.Store.Task().Create(&taskData)
	if err != nil {
		return
	}

	resp, err = taskData.ParseToResponse()
	return
}

func (a *App) GetTaskByID(id string) (resp response.Task, err error) {
	taskData, err := a.Store.Task().GetByID(id)
	if err != nil {
		return
	}

	resp, err = taskData.ParseToResponse()
	return
}

func (a *App) GetListTasks() (resp response.ListTasks, err error) {
	list, _, err := a.Store.Task().GetList()
	if err != nil {
		return
	}
	for _, item := range list {
		respTask, err := item.ParseToResponse()
		if err != nil {
			continue
		}
		resp.List = append(resp.List, respTask)
	}
	return
}

func (a *App) UpdateTask(req request.Task) (resp response.Task, err error) {
	taskData, err := req.ParseToSchema()
	if err != nil {
		return
	}
	err = a.Store.Task().Update(&taskData)
	if err != nil {
		return
	}

	resp, err = taskData.ParseToResponse()
	return
}

func (a *App) DeleteTaskByID(id string) (resp response.UtilResponse, err error) {
	err = a.Store.Task().Delete(id)
	if err != nil {
		resp = response.UtilResponse{
			Status: response.StatusFail,
		}
		return
	}
	resp = response.UtilResponse{
		Status: response.StatusSuccess,
	}
	return
}
