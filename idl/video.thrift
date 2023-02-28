namespace go video

struct StatusResp{
    1:required i32 status_code
    2:optional string status_msg
}

struct PubReq{
    1: required string title(vt.min_size = "1")
    2: required i64 author_id(vt.gt = "0")
    3: required i64 datetime(vt.gt = "0")
    4: required string play_url(vt.min_size = "0")
}

struct PubResp{
    1: required StatusResp status_resp
}

struct ListReq{
    1: required i64 user_id(vt.gt = "0")
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
//  3: optional string avatar//用户头像
//  4: optional string background_image//用户个人页顶部大图
//  5: optional i64 work_count//作品数量
//  6: required bool is_follow // true-已关注，false-未关注
}

struct ListResp{
    1: required StatusResp status_resp
    2: required list<Video> video_list
}

struct FeedReq{
    1: required i64 latest_time // 返回的视频的最新投稿时间戳
}

struct FeedResp {
  1: required i32 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3:required list<Video>  video_list // 视频列表
  4: optional i64  next_time // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct ListLikeReq{
    1: required i64 user_id
}

struct ListLikeResp{
    1: required StatusResp status_resp
    2: list<Video> video_list
}

service VideoService{
    PubResp PublishVideo(1:PubReq req) // 投稿
    ListResp ListVideos(1: ListReq req) // 列出登录用户的投稿视频
    FeedResp ListFeeds(1: FeedReq req) // 未登录用户也可查看
    ListLikeResp ListLikes(1: ListLikeReq req) // 列出登录用户点赞的视频
}
