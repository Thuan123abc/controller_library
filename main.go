package main

import (
	"con_troller_library/database"
	"fmt"
)

func main() {
	stu1 := database.Student{
		11,
		"Thuan",
		26,
		"21A1",
	}
	/*car1:=database.CardLibra{
		11,
		108,
		10,
		4,
		"litebook2112",
	}
	*/
	db := database.NewDB()
	stuRepo1 := database.NewStudentRepo(db)
	err := stuRepo1.CreateStudent(stu1)
	if err != nil {
		fmt.Println("fail\n%v", err)
		return
	}
	fmt.Println("Thanh cong")
}
