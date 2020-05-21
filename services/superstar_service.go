package services

import (
	"microservice/dao"
	"microservice/datasource"
	model "microservice/models"
)


type SuperstaerService interface {
	GetAll() []model.User
	GetById(id int)(model.User )
	Delete(id int)bool
	Update(user *model.User) error
	Search(country string)[]model.User
}

type superstaerService struct {
	dao *dao.SuperstarDao
}

func NewSuperstaerService() *superstaerService {
	return &superstaerService{
		dao: dao.NewSuperstarDao(datasource.InstanceMasterEngin()),
	}
}


func (d *superstaerService)GetAll() []model.User{
	return d.dao.GetAll()
}

func (d *superstaerService)GetById(id int) *model.User {

	return d.dao.Get(id)
}

func (d *superstaerService)Delete(id int) error{

	return d.dao.DeleltUser(id)
}

func (d *superstaerService)Update(user model.User,columns []string) error{

	return d.dao.UpdateUser(user,columns)
}
func (d * superstaerService)Search(country string)[]model.User{
	return d.dao.Search(country)
}