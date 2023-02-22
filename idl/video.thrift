namespace go video

struct StatusResp{
    1:required i32 status_code
    2:optional string status_msg
}

struct PubVideoReq{
    1: required byte data
    2: required string title
    3: required i64 author_id
}

struct PubVideoResp{
    1: required StatusResp status_resp
}

struct ListVideoReq{
    1: required i64 user_id
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
  1: required i64 id = 1; // 用户id
  2: required string name = 2; // 用户名称
  3: optional string avatar = 6; //用户头像
  4: optional string background_image = 7; //用户个人页顶部大图
  5: optional i64 work_count = 10; //作品数量
}

struct ListVideoResp{
    1: required StatusResp status_resp
}

struct FeedReq{
    1: optional i64 latest_time // 返回的视频的最新投稿时间戳
}

struct FeedResp {
  required i32 status_code // 状态码，0-成功，其他值-失败
  optional string status_msg // 返回状态描述
  required list<Video>  video_list // 视频列表
  optional i64  next_time = 4; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

service PublishVideo{
    PubVideoResp PublishVideo(1:PubVideoReq req)
    ListVideoReq ListVideos(1: ListVideoReq req) // 仅仅是列出该用户的视频
    FeedResp ListFeeds(1: FeedReq req)
}
