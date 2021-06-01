package controllers

import (
	"create-go-api/app/models"
	"create-go-api/app/repositories"

	"github.com/revel/revel"
)

type UserController struct {
	*revel.Controller
}

type JsonResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func (c UserController) GetUsers() revel.Result {

	users := repositories.GetAllUsers()

	response := JsonResponse{Success: true, Data: users}

	return c.RenderJSON(response)
}

func (c UserController) CreateUser() revel.Result {
	user := models.User{}
	c.Params.BindJSON(&user)
	revel.AppLog.Debugf("User object Received %v", user)
	repositories.CreateUser(&user)
	response := JsonResponse{Success: true, Data: user}
	return c.RenderJSON(response)
}
