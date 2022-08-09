package database

type Student struct {
	IDStu    int64 `gorm:"primaryKey"`
	NameStu  string
	AgeStu   int64
	ClassStu string
}
type CardLibra struct {
	IDStu    int64 `gorm:"ForeignKey:IDStu"`
	IDCard   int64 `gorm:"primaryKey"`
	LoanDay  int64
	LoanTime int64
	IDBook   string
}

func (s Student) TableName() string {
	return "student"
}
func (c CardLibra) TableName() string {
	return "card"
}
