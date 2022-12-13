
// func main() {
// 	fmt.Println("Hello, World!")
// 	// var x, y int = 12, 14
// 	// var y int = 13
// 	x, y := 44, 12
// 	// y := 12
// 	z := x + y
// 	fmt.Println(z)
// }

package main
import ("fmt")

type Person struct {
  name string
  age int
  job string
  salary int
}
func main() {
	var pers1 Person
	var pers2 Person
  
	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
  
	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500
  
	fmt.Println("Pers1 specification")
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)
  
	fmt.Println("Pers2 specification")
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
  }
