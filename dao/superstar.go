package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	model "microservice/models"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

//NewSuperstarDao
func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao{
	return &SuperstarDao{
		engine: engine,
	}
}

func (d *SuperstarDao) Get(id int) *model.User{
	user :=&model.User{Id: id}
	ok,err := d.engine.Get(user)
	if ok && err == nil {
		return user
	}else {
		return nil
	}
}

func(d *SuperstarDao)GetAll()[]model.User{
	//userlist := make([]model.User,0)
	userlist := []model.User{}
	err := d.engine.Decr("id").Find(&userlist)
	if err !=nil{
		log.Panicln("SuperstarDao GetAll error",err)
		return nil
	}
	return userlist
}

func(d *SuperstarDao)DeleltUser(id int) error{
	user := &model.User{Id: id,State: 1}
	//
	//affected ,err := d.engine.Id(id).Delete(user)    //物理删除  硬删除
	//if affected && err = nil{
	//	return nil
	//}
	_,err := d.engine.Id(id).Update(user)
	log.Panicln("SuperstarDao DeleltUser error",err)
	return err
}
func(d *SuperstarDao)UpdateUser(user model.User,columns []string)error{
	//_,err :=d.engine.Id(user.Id).MustCols(columns...).Update(user)   //针对ID进行更新
	_,err :=d.engine.Where("id=?",user.Id).MustCols(columns...).Update(user)  //指定条件更新   MustCols 强制更新的字段
	return err
}
func(d *SuperstarDao)Search(name string) []model.User{
	userlist := []model.User{}
	err := d.engine.Where("name=?",name).Decr("id").Find(&userlist)
	if err != nil{
		log.Panicln("SuperstarDao Search error",err)
		return nil
	}
	return userlist
}
