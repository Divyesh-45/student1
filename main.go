package student1

import "fmt"

type Student struct {
	Name    string
	City    string
	PinCode uint
}

func S_Detail() {
	res := Student{Name: "divyesh", City: "ahmedabad", PinCode: 33066}
	fmt.Println(res)
}