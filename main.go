package main

import (
	"context"
	"fmt"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/dal/query"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:root@(localhost:3306)/quaver?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	query.SetDefault(db)
	//q := query.Q
	//ctx := context.Background()
	//qc := q.WithContext(ctx)

	// 增
	//insert(ctx, q)

	user, err := query.User.Where(query.User.ID.Eq(1)).First()
	if err != nil {
		return
	}
	fmt.Println(user)
	fmt.Println("Done!")
}

func insert(ctx context.Context, q *query.Query) {
	qc := q.WithContext(ctx)

	// 插入数据
	users := []*model.User{
		{
			Name: "Hi",
			Pwd:  "123",
		},
	}
	err := qc.User.Create(users...)
	if err != nil {
		log.Fatal(err)
	}
}
