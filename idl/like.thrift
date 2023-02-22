namespace go like

include "video.thrift"

struct StatusResp{
    1:required i32 status_code
    2:optional string status_msg
}
struct LikeReq {
    1: required i64 user_id // 用户id
    2: required i64 video_id // 视频id
}

struct LikeResp{
    1: required StatusResp status_resp
}

struct ListLikeReq {
    1: required i64 user_id = 1; // 用户id
}

struct ListLikeResp {
    1: StatusResp status_resp
    2:required list<Video> video_list // 用户点赞视频列表
}
