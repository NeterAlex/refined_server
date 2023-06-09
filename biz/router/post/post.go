// Code generated by hertz generator. DO NOT EDIT.

package post

import (
	post "Refined_service/biz/handler/post"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		{
			_post := _v1.Group("/post", _postMw()...)
			{
				_create := _post.Group("/create", _createMw()...)
				_create.POST("/", append(_createpostMw(), post.CreatePost)...)
			}
			{
				_delete := _post.Group("/delete", _deleteMw()...)
				_delete.DELETE("/:id", append(_deletepostMw(), post.DeletePost)...)
			}
			{
				_query := _post.Group("/query", _queryMw()...)
				_query.GET("/", append(_querypostMw(), post.QueryPost)...)
				_query.GET("/latest", append(_latestpostMw(), post.LatestPost)...)
			}
			{
				_update := _post.Group("/update", _updateMw()...)
				_update.PUT("/:id", append(_updatepostMw(), post.UpdatePost)...)
			}
			{
				_view := _post.Group("/view", _viewMw()...)
				_view.GET("/:id", append(_viewpostMw(), post.ViewPost)...)
			}
		}
	}
}
