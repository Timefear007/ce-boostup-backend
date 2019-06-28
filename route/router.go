package route

import (
	"ce-boostup-backend/api"

	"github.com/labstack/echo"
)

//Init init a router for api
func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	//user handlers
	e.GET("/users", api.GetAllUsers)
	e.GET("/users/:id", api.GetUserWithID)
	e.POST("/users", api.CreateUser)
	e.PUT("/users/:id", api.UpdateUser)
	e.DELETE("/users", api.DeleteAllUsers)
	e.DELETE("/users/:id", api.DeleteUserWithSpecificID)

	//problem handlers
	e.GET("/problems", api.GetAllProblems)
	e.GET("/problems/:id", api.GetProblemWithID)
	e.POST("/problems", api.CreateProblem)
	e.DELETE("/problems", api.DeleteAllProblems)
	e.DELETE("problems/:id", api.DeleteProblemWithSpecificID)

	return e
}
