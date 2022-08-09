package risingfunc

type Student struct {
	IDStu    int64 `gorm:"ForeignKey:IDStu"`
	NameStu  string
	AgeStu   int64
	ClassStu string
}
type CardLibra struct {
	IDStu    int64
	IDCard   int64
	LoanDay  int64
	LoanTime int64
	IDBook   string
}
