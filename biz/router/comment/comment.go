// Code generated by hertz generator. DO NOT EDIT.

package comment

import (
	comment "Refined_service/biz/handler/comment"
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
			_comment := _v1.Group("/comment", _commentMw()...)
			{
				_delete := _comment.Group("/delete", _deleteMw()...)
				_delete.DELETE("/:id", append(_deletecommentMw(), comment.DeleteComment)...)
			}
			{
				_query := _comment.Group("/query", _queryMw()...)
				_query.GET("/", append(_querycommentMw(), comment.QueryComment)...)
			}
			{
				_update := _comment.Group("/update", _updateMw()...)
				_update.PUT("/:id", append(_updatecommentMw(), comment.UpdateComment)...)
			}
		}
	}
}