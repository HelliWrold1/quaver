package api

/*
* @Author: pzqu
* @Date:
* @Description:
 */

type CommentPub struct {
	Id         int64  `json:"id""`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}
type CommentListResp struct {
	BaseResp    BaseResp
	CommentList []*CommentPub `json:"comment_list"`
}
