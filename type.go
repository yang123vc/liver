package main

import "time"

type dbConfigType struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type coolqConfigType struct {
	API   string `json:"api"`
	Token string
}

type liverConfigType struct {
	Admin   int64
	Group   []int64
	Special map[string]string
}

type configType struct {
	Host  string
	DB    dbConfigType `json:"db"`
	Coolq coolqConfigType
	Liver liverConfigType
}

type user struct {
	ID    uint      `gorm:"type:int(10);primary_key;auto_increment"`
	QQ    int64     `gorm:"type:bigint(20);column:qq"`
	Group int64     `gorm:"type:bigint(20);column:group"`
	Grade float64   `gorm:"type:double;column:grade"`
	Next  time.Time `gorm:"type:datetime;column:next"`
	Ban   bool      `gorm:"type:tinyint(1);column:ban"`
}

type memberType struct {
	QQ       int64 `json:"user_id"`
	Nickname string
	Card     string
}

type messageType struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

type postDataType struct {
	PostType    string        `json:"post_type"`
	MessageType string        `json:"message_type"`
	SubType     string        `json:"sub_type"`
	GroupID     int64         `json:"group_id"`
	UserID      int64         `json:"user_id"`
	Message     []messageType `json:"message"`
	RawMessage  string        `json:"raw_message"`
}

type rankType struct {
	QQ    int64   `gorm:"column:qq"`
	Group int64   `gorm:"column:group"`
	Grade float64 `gorm:"column:grade"`
	Rank  int     `gorm:"column:rank"`
}
