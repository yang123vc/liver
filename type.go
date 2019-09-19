package main

import "time"

type dbConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type coolqConfig struct {
	API   string `json:"api"`
	Token string
}

type liverConfig struct {
	Admin   int64
	Group   []int64
	Special map[string]string
}

type config struct {
	Host  string
	DB    dbConfig `json:"db"`
	Coolq coolqConfig
	Liver liverConfig
}

type user struct {
	ID    uint      `gorm:"type:int(10);primary_key;auto_increment"`
	QQ    int64     `gorm:"type:bigint(20);column:qq"`
	Group int64     `gorm:"type:bigint(20);column:group"`
	Grade float64   `gorm:"type:double;column:group"`
	Next  time.Time `gorm:"type:datetime;column:next"`
	Ban   bool      `gorm:"type:tinyint(1);column:ban"`
}

type memberType struct {
	QQ       int64 `json:"user_id"`
	Nickname string
	Card     string
}
