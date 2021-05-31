package controllers

import (
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
