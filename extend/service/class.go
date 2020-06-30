package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/beatrice950201/araneid/extend/model/spider"
)

type DefaultClassService struct{}

/** 根据爬虫文档Id获取一条数据 **/
func (service *DefaultClassService) One(id int) spider.Class {
	var maps spider.Class
	_ = orm.NewOrm().QueryTable(new(spider.Class)).Filter("id", id).One(&maps)
	return maps
}