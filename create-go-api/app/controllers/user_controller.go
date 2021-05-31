package controllers

import "github.com/revel/revel"

type UserController struct {
	*revel.Controller
}

type user struct {
	ID        string
	FirstName string
	LastName  string
}

type JsonResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func (c UserController) GetUsers() revel.Result {

	users := []user{
		{ID: "3333", FirstName: "bob", LastName: "builder"},
		{ID: "1111", FirstName: "carly", LastName: "wanda"},
		{ID: "2222", FirstName: "jenny", LastName: "clarkson"},
	}

	response := JsonResponse{Success: true, Data: users}

	return c.RenderJSON(response)
}
