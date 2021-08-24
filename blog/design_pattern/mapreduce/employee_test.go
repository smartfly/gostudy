package mapreduce

import (
	"fmt"
	"testing"
)

// TestEmployeeAgeMoreThan40 统计有多少员工大于40岁
func TestEmployeeAgeMoreThan40(t *testing.T) {
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)
}

func TestEmployeeSalaryMoreThan6000(t *testing.T) {
	highPayNum := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary people: %d\n", highPayNum)
}

func TestEmployeeNoVacation(t *testing.T) {
	noVacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no vacation: %v\n", noVacation) //People no vacation: [{Hao 44 0 8000} {Jack 26 0 4000} {Marry 29 0 6000}]
}

func TestEmployeeTotalSalaryValue(t *testing.T) {
	totalPay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})
	fmt.Printf("Total Salary: %d\n", totalPay) //Total Salary: 43500
}

func TestStatTotalSalaryByLessThan30(t *testing.T) {
	youngerPay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("younger Salary: %d\n", youngerPay) // 19000
}
