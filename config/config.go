package config

import (
	"io/ioutil"
	"os"
	"regexp"
	"log"
	"encoding/json"
	"sujor.com/leo/sujor-api/utils"
	"fmt"
)

//const DriverName = "mysql"
//const DataSourceName = "root:password@tcp(rm-uf65a5rr68dk8434n.mysql.rds.aliyuncs.com:3306)/testdb?charset=utf8"
//const SecretKey = "sujor"
const ConnectingError = "数据库连接错误!"
const PingError = "数据库无法连接!"
const SqlError = "SQL错误或没有数据!"
const ParamError = "非法参数!"
const RequestSuccess = "请求成功！"
const SignUpSuccess = "注册成功！"
const SignUpError = "用户名已存在！"
const SignInSuccess = "登录成功！"
const SignInError = "用户名或密码错误！"
const SignOutSuccess = "注销成功！"
const SignOutError = "注销失败！"
const InternalServerError = "服务器内部错误"

type DbConfig struct {
	Dialect			string
	Database		string
	User			string
	Password		string
	Host			string
	Port			int
	Charset			string
	URL				string
	MaxIdleConns	int
	MaxOpenConns	int
}

var DBConfig DbConfig

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("read file error: ", err.Error())
		os.Exit(-1)
	}

	// 去注释'/**/' 并验证json
	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		log.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

func initDB() {
	utils.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}

func init() {
	initJSON()
	initDB()
}