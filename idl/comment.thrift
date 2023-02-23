namespace go comment

struct StatusResp{
    1: required i32 status_code
    2: optional string status_msg
}

struct Comment{
    required i64 comment_id
    required i64 user_id
    required string user_name
    required string msg
    required string date // 格式为mm-dd
}

struct PubReq{
    1: required string msg(vt.min_size = "1")
    2: required i64 author_id(vt.gt = "0")
    3: required i64 video_id(vt.gt = "0")
}

struct PubResp{
    1: required StatusResp status_resp
    2: required list<Comment> comment_list
}

struct ListReq{
    1: required i64 video_id(vt.gt = "0")
}

struct ListResp{
    1: required StatusResp status_resp
    2: required list<Comment> comment_list
}

struct DeleteReq{
    1: required i64 video_id
    2: required i64 comment_id
}

struct DeleteResp{
    1: required StatusResp status_resp
}

struct CountReq{
    1: required i64 video_id
}

struct CountResp{
    1: required StatusResp status_resp
    2: required i64 count
}

service CommentService{
    PubResp PublishComment(1:PubReq req)
    ListResp ListComment(1:ListReq req)
    DeleteResp DeleteComment(1: DeleteReq req)
    CountResp CountComments(1: CountReq req)
}