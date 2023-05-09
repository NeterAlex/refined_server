include "comment.thrift"
namespace go post

enum Code{
    Success = 0
    Failed = 1
    ParamInvalid = 2
    DbError = 3
}

struct Post{
    1:required i64 id
    2:required string title
    3:required string content
    4:required string author
    5:required string date
    6:required string tags
    7:required string image_url
    8:optional i64 viewed
    9:optional list<comment.Comment> comments
    10:required i64 userID
}

struct CreatePostRequest{
    1: string title (api.body="title",api.form="title",api.vd="(len($)>0)")
    2: string content (api.body="content",api.form="content")
    3: string author (api.body="author",api.form="author",api.vd="(len($)>0)")
    4: string date (api.body="date",api.form="date",api.vd="(len($)>0)")
    5: string tags (api.body="tags",api.form="tags")
    6: string image_url (api.body="image_url",api.form="image_url")
}

struct CreatePostResponse{
    1: Code code
    2: string msg
}

struct QueryPostRequest{
    1: optional string id (api.query="id",api.body="id", api.form="id",api.path="id")
    2: i64 page (api.query="page",api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
    3: i64 page_size (api.query="page_size",api.body="page_size", api.form="page_size",api.query="page_size",api.vd="$ > 0")
}

struct QueryPostResponse{
    1: Code code
    2: string msg
    3: list<Post> posts
    4: i64 total
}

struct DeletePostRequest{
    1: i64 id (api.path="id",api.vd="$>0")
}

struct DeletePostResponse{
   1: Code code
   2: string msg
}

struct UpdatePostRequest{
    1: string title (api.body="title",api.form="title",api.vd="(len($)>0)")
    2: string content (api.body="content",api.form="content")
    3: string author (api.body="author",api.form="author",api.vd="(len($)>0)")
    4: string date (api.body="date",api.form="date",api.vd="(len($)>0)")
    5: string tags (api.body="tags",api.form="tags")
    6: string image_url (api.body="image_url",api.form="image_url")
}

struct UpdatePostResponse{
    1: Code code
    2: string msg
}

struct ViewPostRequest{
    1: required i64 id (api.path="id")
}

struct ViewPostResponse{
    1: Code code
    2: string msg
}

service PostService{
    UpdatePostResponse UpdatePost(1:UpdatePostRequest req)(api.put="/v1/post/update/:id")
    DeletePostResponse DeletePost(1:DeletePostRequest req)(api.delete="/v1/post/delete/:id")
    QueryPostResponse  QueryPost(1:QueryPostRequest req)(api.get="/v1/post/query/")
    CreatePostResponse CreatePost(1:CreatePostRequest req)(api.post="/v1/post/create/")
    ViewPostResponse ViewPost(1:ViewPostRequest req)(api.get="/v1/post/view/:id")
}