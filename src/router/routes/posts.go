package routes

import (
	"net/http"
	"portfolio-api/src/router/controllers"
)

var postsRoutes = []Route{
	{
		Uri:      "/post",
		Method:   http.MethodGet,
		Function: controllers.SearchPosts,
		Auth:     false,
	},
	{
		Uri:      "/post/{postId}",
		Method:   http.MethodGet,
		Function: controllers.SearchPostById,
		Auth:     false,
	},
	{
		Uri:      "/post/create",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		Auth:     true,
	},
	{
		Uri:      "/post/{postId}",
		Method:   http.MethodPost,
		Function: controllers.UpdatePost,
		Auth:     true,
	},
	{
		Uri:      "/post/{postId}",
		Method:   http.MethodPost,
		Function: controllers.DeletePost,
		Auth:     true,
	},
}
