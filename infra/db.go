package infra

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Addr     string `yaml:"addr"`
	Usernm   string `yaml:"usernm"`
	Passwd   string `yaml:"passwd"`
	Database string `yaml:"db"`
}

type Config struct {
	Mysqlconf DBConfig   `yaml:"mysql"`
	Mailconf  MailConfig `yaml:"gmail"`
}

func LoadConfig(path string) Config {
	var config Config
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func DBConnect() *gorm.DB {
	conn := LoadConfig("infra/config.yaml")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", conn.Mysqlconf.Usernm, conn.Mysqlconf.Passwd, conn.Mysqlconf.Addr, conn.Mysqlconf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
