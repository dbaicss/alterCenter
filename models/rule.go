package models

import (
	orm "alterGateway/database"
)

type Rule struct {
	Metric    string `json:"metric" gorm:"type:varchar(128);"`         //负责人
	Tag     string `json:"tag" gorm:"type:varchar(11);"`           //手机
	Group     string `json:"group" gorm:"type:varchar(64);"`           //邮箱
	Express    string `json:"express" gorm:"type:int(1);"`            //状态
	Value    string `json:"value" gorm:"type:int(1);"`
	CreateBy  string `json:"createBy" gorm:"type:varchar(64);"`
	UpdateBy  string `json:"updateBy" gorm:"type:varchar(64);"`
	BaseModel
}

func (Rule) TableName() string {
	return "rule"
}

func (e *Rule) GetList() ([]Rule, error) {
	var doc []Rule

	table := orm.Eloquent.Table(e.TableName())
	if err := table.Order("updateBy").Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
