// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/HelliWrold1/quaver/dal/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Video{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Video{}) fail: %s", err)
	}
}

func Test_videoQuery(t *testing.T) {
	video := newVideo(db)
	video = *video.As(video.TableName())
	_do := video.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(video.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <video> fail:", err)
		return
	}

	_, ok := video.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from video success")
	}

	err = _do.Create(&model.Video{})
	if err != nil {
		t.Error("create item in table <video> fail:", err)
	}

	err = _do.Save(&model.Video{})
	if err != nil {
		t.Error("create item in table <video> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Video{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <video> fail:", err)
	}

	_, err = _do.Select(video.ALL).Take()
	if err != nil {
		t.Error("Take() on table <video> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <video> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <video> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <video> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Video{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <video> fail:", err)
	}

	_, err = _do.Select(video.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <video> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <video> fail:", err)
	}

	_, err = _do.Select(video.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <video> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <video> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <video> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <video> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Video{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <video> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <video> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <video> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <video> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <video> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <video> fail:", err)
	}
}

var VideoFilterWithNameAndRoleTestCase = []TestCase{}

func Test_video_FilterWithNameAndRole(t *testing.T) {
	video := newVideo(db)
	do := video.WithContext(context.Background()).Debug()

	for i, tt := range VideoFilterWithNameAndRoleTestCase {
		t.Run("FilterWithNameAndRole_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.FilterWithNameAndRole(tt.Input.Args[0].(string), tt.Input.Args[1].(string))
			assert(t, "FilterWithNameAndRole", res1, tt.Expectation.Ret[0])
			assert(t, "FilterWithNameAndRole", res2, tt.Expectation.Ret[1])
		})
	}
}
