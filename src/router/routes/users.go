package routes

import (
	"net/http"
	"portfolio-api/src/router/controllers"
)

var userRoutes = []Route{
	{
		Uri:      "/user/create",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     true,
	},
	{
		Uri:      "/user",
		Method:   http.MethodGet,
		Function: controllers.SearchUsers,
		Auth:     true,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodGet,
		Function: controllers.SearchUserById,
		Auth:     true,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     true,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		Auth:     true,
	},
}
