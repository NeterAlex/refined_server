include "post.thrift"
namespace go user

enum Code{
    Success = 0
    Failed = 1
    ParamInvalid = 2
    DbError = 3
}

struct User{
    1:required i64 id
    2:required string username (go.tag="gorm:\"unique\"")
    3:required string password
    4:required string status
    5:required string email (go.tag="gorm:\"unique\"")
    6:optional string phone
    7:required string nickname
    8:optional list<post.Post> posts
}

struct CreateUserRequest{
    1: string username (api.body="username",api.form="username",api.vd="(len($)>0 && len($)<20)")
    2: string password (api.body="password",api.form="password",api.vd="(len($)>6 && len($)<18)")
    3: string nickname (api.body="nickname",api.form="nickname",api.vd="(len($)>0 && len($)<20)")
    4: string email (api.body="email",api.form="email",api.vd="(len($)>0)")
    5: optional string phone (api.body="phone",api.form="phone")
}

struct CreateUserResponse{
    1: Code code
    2: string msg
}

struct QueryUserRequest{
    1: string id (api.query="id",api.body="id", api.form="id",api.query="id")
    2: i64 page (api.query="page",api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
    3: i64 page_size (api.query="page_size",api.body="page_size", api.form="page_size",api.query="page_size",api.vd="$ > 0")
}

struct QueryUserResponse{
    1: Code code
    2: string msg
    3: list<User> users
    4: i64 total
}

struct DeleteUserRequest{
    1: i64 id (api.path="id",api.vd="$>0")
}

struct DeleteUserResponse{
   1: Code code
   2: string msg
}

struct UpdateUserRequest{
    1: i64 id (api.path="id",api.vd="$>0")
    2: string username (api.body="username",api.form="username",api.vd="(len($)>0 && len($)<20)")
    3: optional string password (api.body="password",api.form="password",api.vd="(len($)<18)")
    4: string nickname (api.body="nickname",api.form="nickname",api.vd="(len($)>0 && len($)<20)")
    5: string email (api.body="email",api.form="email",api.vd="(len($)>0)")
    6: optional string phone (api.body="phone",api.form="phone")
}

struct UpdateUserResponse{
    1: Code code
    2: string msg
}

service UserService{
    UpdateUserResponse UpdateUser(1:UpdateUserRequest req)(api.put="/v1/user/update/:id")
    DeleteUserResponse DeleteUser(1:DeleteUserRequest req)(api.delete="/v1/user/delete/:id")
    QueryUserResponse  QueryUser(1:QueryUserRequest req)(api.get="/v1/user/query/")
    CreateUserResponse CreateUser(1:CreateUserRequest req)(api.post="/v1/user/create/")
}