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

type AgeMap struct {
	Teen int
	Adult int
	Elderly int
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
	user := User{Username: name, Password: psw, LoginTime: time.Now(), RegTime: time.Now(), Gender: gender, Dob: date}
	_, err := o.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

func UserExists(o orm.Ormer, tableName string, name string) bool {
	return o.QueryTable(tableName).Filter("Username", name).Exist()
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

func GetUserInfo(o orm.Ormer, name string) *User {
	var users []User
	_, err := o.QueryTable("user").Filter("Username", name).All(&users)
	if err != nil {
		return nil
	}
	return &users[0]
}

/*
The following methods are for the purpose of data visualization.
 */
func GetNumber(o orm.Ormer, colName string, colValue string) int64 {
	var users []User
	num, err := o.QueryTable("user").Filter(colName, colValue).All(&users)
	if err != nil {
		err.Error()
		return 0
	}
	return num
}

func GetAges(o orm.Ormer) []int {
	var users []User
	//var year int
	num, err := o.QueryTable("user").All(&users)
	agelist := make([]int, num)
	var maps []orm.Params
	_, err = o.QueryTable("user").Values(&maps)
	if err != nil {
		err.Error()
	}
	for i, m := range maps {
		if date, ok := m["Dob"].(time.Time); ok {
			year := date.Year()
			age := time.Now().Year() - year
			agelist[i] = age
		}
	}
	return agelist
}

func (m *AgeMap) CreateAgeMap(agelist []int) {
	for _, v := range agelist {
		if v <= 18 {
			m.Teen+=1
		} else if v <= 65 {
			m.Adult+=1
		} else {
			m.Elderly+=1
		}
	}
}














