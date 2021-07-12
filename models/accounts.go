package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int `orm:"auto"`
	Username string `orm:"index"`
	Password orm.CharField `orm:"size(60)"`
	LoginTime time.Time
	RegTime time.Time
	Gender string
	Dob orm.DateField `orm:"null"`
}

type RegAccount struct {
	Username string `form:"userName"`
	Password string `form:"userPassword"`
	CfmPassword string `form:"cuserPassword"`
	Gender string `form:"userGender"`
	Dob time.Time `form:"userDob"`
}

type LoginAccount struct {
	Username string `form:"userName"`
	Password string `form:"userPassword"`
}

func fetchConfig(key string) string {
	config, _ := web.AppConfig.String(key)
	return config
}

func init() {
	driver := fetchConfig("driverName")
	sqlUser := fetchConfig("mysqlUser")
	sqlPass := fetchConfig("mysqlPass")
	sqlUrl := fetchConfig("mysqlUrl")
	dbname := fetchConfig("dbname")
	attr := fetchConfig("attr")
	config := sqlUser + ":" + sqlPass + "@tcp(" + sqlUrl + ")/" + dbname + attr
	err := orm.RegisterDataBase("default", driver, config)
	if err != nil {
		fmt.Println("failed to connect!")
		return
	}
	orm.RegisterModel(new(User))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		return
	}
}

func toDate(t time.Time) orm.DateField {
	date := orm.DateField(t)
	return date
}

func InsertUser(o orm.Ormer, name string, psw orm.CharField, gender string, dob time.Time) error {
	var date orm.DateField
	date = toDate(dob)
	user := User{Username:name, Password: psw, LoginTime: time.Now(), RegTime: time.Now(), Gender: gender, Dob: date}
	_, err := o.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

func UserExists(o orm.Ormer, name string) bool {
	return o.QueryTable("user").Filter("Username", name).Exist()
}

func GetPassword(o orm.Ormer, name string) string {
	var str string
	var maps []orm.Params
	_, err := o.QueryTable("user").Values(&maps)
	if err != nil {
		err.Error()
	}
	for _, m := range maps {
		if m["Username"] == name {
			str = fmt.Sprintf("%s", m["Password"])
			break
		}
	}
	return str
}

func UpdateLoginTime(o orm.Ormer, name string) error {
	_, err := o.QueryTable("user").Filter("username", name).Update(orm.Params{"login_time": time.Now()})
	return err
}












