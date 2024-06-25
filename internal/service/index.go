package service

import (
	"pitaya/internal/global"
	"pitaya/internal/model"
)

type ProjectList struct{}

// 获取列表
func (p *ProjectList) List() ([]*model.ProjectList, error) {
	pList := make([]*model.ProjectList, 0)
	err := global.DB.Model(&model.ProjectList{}).Find(&pList).Error
	return pList, err
}

// 添加
func (p *ProjectList) Add(pList *model.ProjectList) error {
	return global.DB.Create(pList).Error
}

// 删除
func (p *ProjectList) Delete(id uint) error {
	return global.DB.Delete(&model.ProjectList{}, id).Error
}

// 更新
func (p *ProjectList) Update(pList *model.ProjectList) error {
	return global.DB.Save(pList).Error
}
