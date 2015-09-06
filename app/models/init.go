/*
 * init.go
 */

package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/robfig/config"
	//"strings"
	//"time"
)

var (
	Engine *xorm.Engine
)

func init() {
	revel.OnAppStart(Init)
}

func Init() {
	var err error
	var c *config.Config
	c, _ = config.ReadDefault(revel.BasePath + "/conf/db.conf")
	dbDriver, _ := c.String("db", "db.driver")
	dbUrl, _ := c.String("db", "db.url")

	Engine, err = xorm.NewEngine(dbDriver, dbUrl)
	// defer Engine.Close()
	if err != nil {
		panic(err)
	}

}
