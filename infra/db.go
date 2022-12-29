package infra

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
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

// func ConnectDatabase() *gorm.DB {
// 	conn := LoadConfig("infra/config.yaml")
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", conn.Mysqlconf.Usernm, conn.Mysqlconf.Passwd, conn.Mysqlconf.Addr, conn.Mysqlconf.Database)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetMaxOpenConns(100)

// 	defer sqlDB.Close()

// 	return db
// }

func GetConnector() *sql.DB {
	conn := LoadConfig("infra/config.yaml")
	cfg := mysql.Config{
		User:                 conn.Mysqlconf.Usernm,
		Passwd:               conn.Mysqlconf.Passwd,
		Net:                  "tcp",
		Addr:                 conn.Mysqlconf.Addr,
		Collation:            "utf8mb4_general_ci",
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               "users",
	}

	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	db := sql.OpenDB(connector)

	return db
}
