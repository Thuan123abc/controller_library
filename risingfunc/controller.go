package risingfunc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s Student) InputStudent() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Nhap ID cua sinh vien")
	id, _ := consoleReader.ReadString('\n')
	id = strings.Replace(id, "\n", "", -1)
	idInt, _ := strconv.ParseInt(id, 10, 64)
	s.IDStu = idInt

	fmt.Println("Nhap ten cua sinh vien")
	name, _ := consoleReader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	s.NameStu = name

	fmt.Println("Nhap tuoi cua sinh vien")
	age, _ := consoleReader.ReadString('\n')
	age = strings.Replace(age, "\n", "", -1)
	ageInt, _ := strconv.ParseInt(age, 10, 64)
	s.AgeStu = ageInt

	fmt.Println("Nhap lop cua sinh vien")
	class, _ := consoleReader.ReadString('\n')
	class = strings.Replace(class, "\n", "", -1)
	s.ClassStu = class
}

func (c CardLibra) InputCard() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Nhap ID cua sinh vien muon sach")
	id, _ := consoleReader.ReadString('\n')
	id = strings.Replace(id, "\n", "", -1)
	idInt, _ := strconv.ParseInt(id, 10, 64)
	c.IDStu = idInt

	fmt.Println("Nhap ID cua card")
	idc, _ := consoleReader.ReadString('\n')
	idc = strings.Replace(idc, "\n", "", -1)
	idcInt, _ := strconv.ParseInt(idc, 10, 64)
	c.IDCard = idcInt

	fmt.Println("Nhap ngay sinh vien muon sach")
	day, _ := consoleReader.ReadString('\n')
	day = strings.Replace(day, "\n", "", -1)
	dayInt, _ := strconv.ParseInt(day, 10, 64)
	c.LoanDay = dayInt

	fmt.Println("Nhap so ngay sinh vien muon sach")
	time, _ := consoleReader.ReadString('\n')
	time = strings.Replace(day, "\n", "", -1)
	timeInt, _ := strconv.ParseInt(time, 10, 64)
	c.LoanTime = timeInt

	fmt.Println("Nhap ma so sach sinh vien muon")
	ms, _ := consoleReader.ReadString('\n')
	ms = strings.Replace(day, "\n", "", -1)
	c.IDBook = ms
}