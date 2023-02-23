namespace go user

enum ErrCode{
    SuccessCode = 0
    ServiceErrCode = 10001
    ParamErrCode = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct StatusResp{
    1:required i32 status_code
    2:optional string status_msg
}

struct RegisterReq{
    1: required string username(vt.min_size = "1", vt.max_size = "32")
    2: required string password(vt.min_size = "6", vt.max_size = "32")
}

struct RegisterResp{
    1: required StatusResp status_resp
    2: required i64 user_id
}

struct LoginReq{
    1: required string username(vt.min_size = "1")
    2: required string password(vt.min_size = "6")
}

struct LoginResp{
    1: required StatusResp status_resp
    2: required i64 user_id
}

struct InfoReq{
    1: required i64 user_id(vt.gt = "0")
}

struct InfoResp{
    1: required i64 user_id
    2: required string username
}


service UserService{
    RegisterResp UserRegister(1: RegisterReq req)
    LoginResp UserLogin(1: LoginReq req)
    InfoResp UserInfo(1: InfoReq req)
}