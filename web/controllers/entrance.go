package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"microservice/services"

)
type IndexController struct {
	Ctx iris.Context
	Svervice services.SuperstaerService

}

func (c *IndexController) GetAll() mvc.Result{

	return nil
}

func (c *IndexController)GetUser(id int)mvc.Result{
	return nil
}

func (c *IndexController)Getsearch(country string)mvc.Result{
	return nil
}

func (c IndexController)DelectUser(id int)mvc.Result{
	return nil
}