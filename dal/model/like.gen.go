// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLike = "like"

// Like mapped from table <like>
type Like struct {
	ID      int64 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	VideoID int64 `gorm:"column:video_id;type:int unsigned;not null" json:"video_id"`
	LikerID int64 `gorm:"column:liker_id;type:int unsigned;not null" json:"liker_id"`
	Like    int64 `gorm:"column:like;type:tinyint;not null;default:1" json:"like"`
}

// TableName Like's table name
func (*Like) TableName() string {
	return TableNameLike
}
