package database

import "gorm.io/gorm"

type StudentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *StudentRepo {
	db.AutoMigrate(&Student{})
	return &StudentRepo{
		db: db,
	}
}
func (s StudentRepo) CreateStudent(stu Student) error {
	err := s.db.Create(&stu).Error
	if err != nil {
		return err
	}
	return nil
}
func (s StudentRepo) UpdateStudent(stu Student) error {
	err := s.db.Where("id_stu=?", stu.IDStu).Omit("id_stu").Updates(stu).Error
	if err != nil {
		return err
	}
	return nil
}
func (s StudentRepo) DeleteStudent(stu Student) error {
	err := s.db.Delete(&stu).Error
	if err != nil {
		return err
	}
	return nil
}

func (s StudentRepo) GetListStudent() []Student {
	var ListStu []Student
	err := s.db.Find(ListStu)
	if err != nil {
		return nil
	}
	return ListStu
}

func (s StudentRepo) GetStudent(id int64) (stu Student) {
	err := s.db.Where("id_stu =?", id).First(&stu)
	if err != nil {
		return
	}
	return stu
}
