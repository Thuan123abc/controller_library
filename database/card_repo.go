package database

import (
	"fmt"
	"gorm.io/gorm"
)

type CardRepo struct {
	db *gorm.DB
}

func NewCardRepo(db *gorm.DB) *CardRepo {
	db.AutoMigrate(&CardLibra{})
	return &CardRepo{
		db: db,
	}
}
func (c CardRepo) CreateCard(car CardRepo) error {
	err := c.db.Create(&car).Error
	if err != nil {
		return err
	}
	return nil
}
func (c CardRepo) UpdateCard(car CardLibra) error {
	err := c.db.Where("id_card=?", car.IDCard).Omit("id_card").Updates(car).Error
	if err != nil {
		return err
	}
	return nil
}
func (c CardRepo) DeleteCard(car CardLibra) error {
	err := c.db.Delete(&car).Error
	if err != nil {
		return err
	}
	return nil
}

func (c CardRepo) GetListCard() []CardLibra {
	var ListCard []CardLibra
	err := c.db.Find(ListCard)
	if err != nil {
		return nil
	}
	return ListCard
}

func (c CardRepo) GetCard(id int64) (car CardLibra) {
	err := c.db.Where("id_card =?", id).First(&car)
	if err != nil {
		fmt.Println("khong co the thu vien nhu vay!")
		return
	}
	return car
}
