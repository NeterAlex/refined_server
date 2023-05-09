namespace go comment

enum Code{
    Success = 0
    Failed = 1
    ParamInvalid = 2
    DbError = 3
}

struct Comment{
    1:required i64 id
    2:required string author
    3:required string content
    4:required i64 postID
}

struct CreateCommentRequest{
    1: string author (api.body="author",api.form="author",api.vd="(len($)>0)")
    2: string content (api.body="content",api.form="content",api.vd="(len($)>0)")
    3: i64 postID (api.body="postID",api.form="postID")
}

struct CreateCommentResponse{
    1: Code code
    2: string msg
}

struct QueryCommentRequest{
    1: optional string id (api.query="id",api.body="id", api.form="id",api.query="id")
    2: i64 page (api.query="page",api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
    3: i64 page_size (api.query="page_size",api.body="page_size", api.form="page_size",api.query="page_size",api.vd="$ > 0")
}

struct QueryCommentResponse{
    1: Code code
    2: string msg
    3: list<Comment> comments
    4: i64 total
}

struct DeleteCommentRequest{
    1: i64 id (api.path="id",api.vd="$>0")
}

struct DeleteCommentResponse{
   1: Code code
   2: string msg
}

struct UpdateCommentRequest{
    1: string author (api.body="author",api.form="author",api.vd="(len($)>0)")
    2: string content (api.body="content",api.form="content",api.vd="(len($)>0)")
    3: i64 id (api.body="id",api.form="id",api.path="id")
    4: i64 cid (api.body="cid",api.form="cid")
}

struct UpdateCommentResponse{
    1: Code code
    2: string msg
}

service CommentService{
    CreateCommentResponse CreateComment(1:CreateCommentRequest req)(api.post="/v1/comment/create/")
    UpdateCommentResponse UpdateComment(1:UpdateCommentRequest req)(api.put="/v1/comment/update/:id")
    DeleteCommentResponse DeleteComment(1:DeleteCommentRequest req)(api.delete="/v1/comment/delete/:id")
    QueryCommentResponse  QueryComment(1:QueryCommentRequest req)(api.get="/v1/comment/query/")
}