package mysql

import (
	"fmt"
	"github.com/Minggyyds/framework/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const ip = "127.0.0.1"
const p = 8848

var DB *gorm.DB

type mysqlConfig struct {
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Dbname string `yaml:"dbname"`
}

func InitMysql(serviceName string) error {
	type Val struct {
		Mysql mysqlConfig `yaml:"mysql"`
	}
	mysqlConfigVal := Val{}
	content, err := config.GetConfig("DEFAULT_GROUP", serviceName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(content), &mysqlConfigVal)
	if err != nil {
		fmt.Println("**********errr")
		return err
	}
	fmt.Println(content)
	fmt.Println(mysqlConfigVal)
	configM := mysqlConfigVal.Mysql
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configM.User,
		configM.Pwd,
		configM.Host,
		configM.Port,
		configM.Dbname,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func WithTX(txFc func(tx *gorm.DB) error) {
	var err error
	tx := DB.Begin()
	err = txFc(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
