package repository

import (
	"github.com/jinzhu/gorm"
	"myblog/database"
	"myblog/model"
)

type CategortRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategortRepository {
	return CategortRepository{
		DB: database.InitDB(),
	}
}

func (c CategortRepository) Create(name string) (*model.Category, error) {
	category := &model.Category{
		Name: name,
	}

	if err := c.DB.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategortRepository) Update(category model.Category, id uint) (*model.Category, error) {
	category.ID = id
	if err := c.DB.Model(&model.Category{}).Update(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil

}

func (c CategortRepository) Select(id uint) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return &category, err
	}
	return &category, nil
}

func (c CategortRepository) DeleteById(id int) error {
	if err := c.DB.Delete(model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
