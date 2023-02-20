// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID       int64     `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	AuthorID int64     `gorm:"column:author_id;type:int unsigned;not null;index:ctr_id,priority:1" json:"author_id"` // 评论者id
	VideoID  int64     `gorm:"column:video_id;type:int unsigned;not null;index:ctrv_id,priority:1" json:"video_id"`  // 被评论的视频id
	Msg      string    `gorm:"column:msg;type:varchar(255);not null" json:"msg"`                                     // 评论内容
	Datetime time.Time `gorm:"column:datetime;type:datetime;not null" json:"datetime"`                               // 发表日期
	Cancel   int64     `gorm:"column:cancel;type:tinyint;not null" json:"cancel"`                                    // 是否删除评论，默认为未删除：0
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
