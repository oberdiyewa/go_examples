package examples

import "fmt"

type Finalsalary func(int, int) int

type Developer struct {
	name       string
	level      string
	age        int
	experience int
	salary     Finalsalary
}

func FuncType() {
	developer := Developer{
		name:       "Ezkahan",
		level:      "middle",
		age:        24,
		experience: 10,
		salary: func(first int, second int) int {
			return first * second
		},
	}

	fmt.Println("Name", developer.name)
	fmt.Println("Level", developer.level)
	fmt.Println("Salary", developer.salary(developer.age, developer.experience))
}
