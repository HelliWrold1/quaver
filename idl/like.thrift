namespace go like

include "video.thrift"

struct StatusResp{
    1:required i32 status_code
    2:optional string status_msg
}

struct Video {
  1: required i64 video_id // 视频唯一标识
  2: required User author // 视频作者信息
  3: required string play_url // 视频播放地址
  4: required string cover_url  // 视频封面地址
  5: required i64 favorite_count // 视频的点赞总数
  6: required i64 comment_count // 视频的评论总数
  7: required bool is_favorite // true-已点赞，false-未点赞
  8: required string title // 视频标题
}

struct User {
  1: required i64 id// 用户id
  2: required string name// 用户名称
  3: optional string avatar//用户头像
  4: optional string background_image//用户个人页顶部大图
  5: optional i64 work_count//作品数量
}

struct LikeReq {
    1: required i64 user_id(vt.gt = "0")// 用户id
    2: required i64 video_id(vt.gt = "0") // 视频id
}

struct LikeResp{
    1: required StatusResp status_resp
}

struct DeleteReq{
    1: required i64 user_id(vt.gt = "0")// 用户id
    2: required i64 video_id(vt.gt = "0")// 视频id
}

struct DeleteResp{
    1: required StatusResp status_resp
}

struct ListReq {
    1: required i64 user_id(v.gt = "0") // 用户id
}

struct ListResp {
    1: StatusResp status_resp
    2:required list<Video> video_list // 用户点赞视频列表
}

service LikeService{
    LikeResp LikeVideo(1: LikeReq req)
    DeleteResp DeleteLike(1: DeleteReq req)
    ListResp ListLikes(1: ListReq req)
}
