package database

import (
	"testing"

	"github.com/liteseed/bungo/schema"
	"github.com/stretchr/testify/assert"
)

func TestNewWdb(t *testing.T) {
	// dsn := "root@tcp(127.0.0.1:3306)/arseed?charset=utf8mb4&parseTime=True&loc=Local"
	// wdb := NewWdb(dsn)
	// wdb.Migrate(false,true)
	//
	// itemId := "57uEQ3iLKxhmWhdcJJh-22mou0xVTCJq8LN9ad20Bcg"
	// exist := wdb.ExistPaidOrd(itemId)
	// t.Log(exist)
	// expiredTime := int64(1661333743)
	// latest := wdb.IsLatestUnpaidOrd(itemId,expiredTime)
	// t.Log(latest)
}

func TestSqlite(t *testing.T) {
	db := NewSqliteDatabase("testSqlite")
	err := db.Migrate()
	assert.NoError(t, err)
	ord := &schema.Order{}
	err = db.DB.First(ord).Error
	assert.NoError(t, err)
	t.Log(ord)
}

func TestNew(t *testing.T) {
	db := NewMysqlDatabase("root@tcp(127.0.0.1:3306)/arseed?charset=utf8mb4&parseTime=True&loc=Local")
	err := db.Migrate()
	assert.NoError(t, err)
}
