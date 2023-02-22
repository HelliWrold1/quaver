package db

import (
	"fmt"
	"github.com/HelliWrold1/quaver/dal/model"
)

func migration() {
	// 设置存储引擎为InnoDB
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		// 自动迁移，建表
		AutoMigrate(
			&model.Comment{},
		)
	if err != nil {
		fmt.Println("migrate err", err)
	}
}
