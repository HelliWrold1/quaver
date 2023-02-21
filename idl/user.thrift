namespace go user

enum ErrCode{
    SuccessCode = 0
    ServiceErrCode = 10001
    ParamErrCode = 1002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct StatusResp{
    1:required i32 status_code
    2:optional string status_msg
}

struct RegisterReq{
    1: required string username
    2: required string password
}

struct RegisterResp{
    1: required StatusResp status_resp
    2: required i64 user_id
    3: required string token
}

struct LoginReq{
    1: required string username
    2: required string password
}

struct LoginResp{
    1: required StatusResp status_resp
    2: required i64 user_id
    3: required string token
}

struct InfoReq{
    1: required i64 user_id
    2: required string token
}

struct InfoResp{
    1: required StatusResp status_resp
    2: required User user
}

struct User{
    1: required i64 id
    2: required string name
    3: optional i64 follow_count // 关注总数
    4: optional i64 follower_count // 粉丝总数
    5: required bool is_follow // 是否已关注
    6: optional string avatar // 头像
    7: optional string background_image // 主页背景图
    8: optional string signature // 个人简介
    9: optional i64 total_favorited // 获赞数量
    10: optional i64 work_count // 作品数量
    11: optional i64 favorite_count // 点赞数量
}

service UserService{
    RegisterResp UserRegister(1: RegisterReq req)
    LoginResp UserLogin(1: LoginReq req)
    InfoResp UserInfo(1: InfoReq req)
}