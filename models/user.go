package models

type User struct {
	Id       int    `xorm:"not null INT(11)"`
	Name     string `xorm:"VARCHAR(100)"`
	Age      int    `xorm:"INT(11)"`
	Password string `xorm:"VARCHAR(100)"`
	State    int    `xorm:"default 0 INT(11)"`
}
