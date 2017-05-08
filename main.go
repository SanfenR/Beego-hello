package main

import (
	_ "Beego-hello/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"Beego-hello/model"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func main() {
	o := orm.NewOrm()
	o.Using("default")


	profile := new(model.Profile)
	profile.Age = 30

	beego.Run()

}

