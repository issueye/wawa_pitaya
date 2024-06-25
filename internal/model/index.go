package model

import "time"

type ProjectList struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey;comment:项目ID"`        // 项目id
	Topic       string    `json:"topic" gorm:"column:topic;comment:主题"`               // 主题
	Description string    `json:"description" gorm:"column:description;comment:项目描述"` // 项目描述
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;comment:项目创建时间"` // 项目创建时间
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;comment:项目更新时间"` // 项目更新时间
}

func (ProjectList) TableName() string {
	return "project_list"
}
