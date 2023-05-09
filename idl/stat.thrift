namespace go stat

enum Code{
    Success = 0
    Failed = 1
    ParamInvalid = 2
    DbError = 3
}

struct Stat{
    1:required i64 post_count
    2:required i64 comment_count
    3:required i64 user_count
    4:required i64 total_viewed
}

struct QueryStatRequest{

}

struct QueryStatResponse{
    1:Code code
    2:string msg
    3:Stat stat
}

service StatService{
    QueryStatResponse QueryStat()(api.get="/v1/stat/query")
}