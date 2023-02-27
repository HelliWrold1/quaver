namespace go api

struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct User{
    1:i64 Id
    2:string name
    3:string password
}

struct Like{
    1:i64 Id
    2:i64 user_Id
    3:i64 video_Id
    4:i8 cancel
}

struct Video{
    1: i64 Id
    2: i64 AuthorId // 作者id
    3: string play_url  // 播放url
    4: string ConverUrl // 封面url
    5: i64 PublishTime // 发布时间
    6: string Title //视频名
}

struct Comment{
    1:i64 id
    2:i64 user_id
    3:i64 video_id
    4:string comment_text
    5:i64 create_time
    6:i8 cancel
}

struct UserRegisterRequest{
    1:string username(api.query="username", api.vd="len($) > 0")
    2:string password(api.query="password", api.vd="len($) > 0")
}

struct UserRegisterResponse{
    1:BaseResp baseresp
    2:i64 user_id
    3:string token
}

struct UserLoginRequest{
    1:string username
    2:string password
}

struct UserLoginResponse{
    1:BaseResp baseresp
    2:i64 user_id
    3:i64 token
}

struct UserInfoRequest{
    1:i64 user_id(api.query="user_id")
}

struct UserInfoResponse{
    1:BaseResp baseresp
    2:User user
}

struct FavouriteActionRequest{
    1:i64 user_id
    2:i64 video_id(api.query="video_id", api.vd="len($) > 0")
    3:i64 action_type(api.query="action_type", api.vd="len($) > 0")
}

struct FavouriteActionResponse{
    1:BaseResp baseresp
}

struct FavouriteListRequest{
    1:i64 user_id(api.query="user_id", api.vd="len($) > 0")
    2:i64 cur_user_id
}

struct FavouriteListResponse{
    1:BaseResp baseresp
    2:list<Video> videos
}

struct CommentActionRequest{
    1:i64 video_id(api.query="video_id", api.vd="len($) > 0")
    2:i64 action_type(api.query="action_type", api.vd="len($) > 0")
    3:string comment_text(api.query="comment_text", api.vd="len($) > 0")
    4:optional string comment_id(api.query="comment_id",api.vd="len($) > 0")
}

struct CommentActionResponse{
    1:BaseResp baseresp
    2:Comment comment
}

struct CommentListRequest{
    1:i64 cur_user_id
    2:i64 video_id(api.query="video_id", api.vd="len($) > 0")
}

struct CommentListResponse{
    1:BaseResp baseresp
    2:list<Comment> commentList
}

struct VideoFeedRequest{
    1: optional i64 latest_time(api.query="latest_time", api.vd="len($) > 0")
}

struct VideoFeedResponse{
    1:BaseResp baseresp
    2:list<Video> videos
    3:i64 next_time
}

struct VideoPublishResponse{
    1:BaseResp baseresp
}

struct VideoPublishListRequest{
    1:i64 user_id(api.query="user_id",api.vd="len($) > 0")
}

struct VideoPublishListResponse{
    1:BaseResp baseresp
    2:list<Video> VideoList
}

service Apiservice{
    UserRegisterResponse  UserRegister(1:UserRegisterRequest req) (api.post="/douyin/user/register")
    UserLoginResponse UserLogin(1:UserLoginRequest req) (api.post="/douyin/user/login")
    UserInfoResponse UserInfo(1:UserInfoRequest req) (api.get="douyin/user/")
    FavouriteActionResponse LikeAction(1:FavouriteActionRequest req) (api.post="/douyin/favorite/action/")
    FavouriteListResponse FavouriteList(1:FavouriteListRequest req) (api.get="/douyin/favorite/list/")
    CommentActionResponse CommentAction(1:CommentActionRequest req) (api.post="/douyin/comment/action/")
    CommentListResponse CommentList(1:CommentListRequest req) (api.get="/douyin/comment/list/")
    VideoFeedResponse VideoFeed(1:VideoFeedRequest req) (api.get="/douyin/feed/")
    VideoPublishResponse VideoPublish()(api.post="/douyin/publish/action/")
    VideoPublishListRequest VideoPublishList(1:VideoPublishListRequest req)(api.get="/douyin/publish/list/")
}