package main

import (
	"fmt"
	"time"
)

// testcommit
func main() {
	//a := 4
	//b := 5
	//c := plus(a, b)
	//fmt.Println(c, "plus")
	//
	//printSum(a, b)
	//
	//five := giveMeFive()
	//fmt.Println(five, "giveMeFive")
	//
	//negative, isNegative := isNegativeFind(-5)
	//fmt.Println(negative, isNegative, "isNegative")
	//
	//printCurrentTime()
	//
	//evenNumbers := returnEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	//fmt.Println(evenNumbers, "returnEvenNumbers")

	//////////////////////////////////////////////

	//coffeeMachine := CoffeeMachine{
	//	Water:  1000,
	//	Coffee: 1000,
	//	Milk:   1000,
	//}
	//
	//coffeeMachine.MakeEspresso()
	//fmt.Println(coffeeMachine, "MakeEspresso")
	//
	//coffeeMachine.MakeLatte()
	//fmt.Println(coffeeMachine, "MakeLatte")
	//
	//coffeeMachine.MakeCappuccino()
	//fmt.Println(coffeeMachine, "MakeCappuccino")
	//
	//status := coffeeMachine.Status()
	//fmt.Println(status, "Status")
	//
	//coffeeMachine.Refill(2000, 2000, 2000)
	//fmt.Println(coffeeMachine, "Refill")
	//
	//coffeeMachine.MakeEspresso()
	//fmt.Println(coffeeMachine, "MakeEspresso")
	//
	//status = coffeeMachine.Status()
	//fmt.Println(status, "Status")
	//
	//machine := coffeeMachine.GetCoffeeMachine()
	//fmt.Println(machine, "GetCoffeeMachine")
	//
	//coffeeMachine2 := CoffeeMachine{
	//	Water:  3000,
	//	Coffee: 3000,
	//	Milk:   3000,
	//}
	//
	//coffeeMachine2.MakeEspresso()
	//fmt.Println(coffeeMachine2, "MakeEspresso2")

	////////////////////////////////////////////////

	//var driver Driver
	//driver.Name = "John"
	//driver.Age = 18
	//driver.Experience = 1
	//driver.Car = Car{
	//	Brand: "BMW",
	//	Model: "X5",
	//	Year:  2020,
	//	Class: Class{
	//		Name:        "Premium",
	//		Coefficient: 3.5,
	//	},
	//}

	//fmt.Println(driver, "driver")

	//dr := AddDriver("Jack", 25, 5, driver.Car)
	//fmt.Println(dr, "AddDriver")
	//dr.Drive()
	//
	//dr.ChangeCar("Audi", "A6", 2019, Class{
	//	Name:        "Premium",
	//	Coefficient: 2.9,
	//})
	//fmt.Println(dr, "AddDriver 2")
	//
	//dr.ChangeCoefficient(6.5)
	//fmt.Println(dr, "AddDriver 3")

	//////////////////////anonymous func////////////////////////

	//func() {
	//	fmt.Println("Hello")
	//}()
	//
	//func(a, b int) {
	//	fmt.Println(a + b)
	//}(a, b)

	//for i := 0; i < 5; i++ {
	//	go fmt.Println(i, "==========i")
	//}
	//time.Sleep(time.Second * 1)

	//for i := 0; i < 5; i++ {
	//	go func() {
	//		fmt.Println(i, "------i")
	//	}()
	//}
	//time.Sleep(time.Second * 2)

	//for i := 0; i < 5; i++ {
	//	go func(i int) {
	//		fmt.Println(i, "*****************i")
	//	}(i)
	//}
	//time.Sleep(time.Second * 2)
}

func plus(a, b int) int {
	return a + b
}

func printSum(a, b int) {
	fmt.Println(plus(a, b))
}

func giveMeFive() int {
	return 5
}

func isNegativeFind(a int) (int, bool) {
	if a < 0 {
		return -a, true
	}
	return a, false
}

func printCurrentTime() {
	fmt.Println(time.Now())
}

func returnEvenNumbers(a []int) []int {
	var result []int
	for _, v := range a {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

///////////////////////////////////////////
///////////////////////////////////////////

type CoffeeMachine struct {
	Water  int
	Coffee int
	Milk   int
}

func (c *CoffeeMachine) MakeEspresso() {
	c.Water -= 250
	c.Coffee -= 16
}

func (c *CoffeeMachine) MakeLatte() {
	c.Water -= 350
	c.Coffee -= 20
	c.Milk -= 75
}

func (c *CoffeeMachine) MakeCappuccino() {
	c.Water -= 200
	c.Coffee -= 12
	c.Milk -= 100
}

func (c *CoffeeMachine) Status() string {
	return fmt.Sprintf("Water: %d, Coffee: %d, Milk: %d", c.Water, c.Coffee, c.Milk)
}

func (c *CoffeeMachine) Refill(water, coffee, milk int) {
	c.Water += water
	c.Coffee += coffee
	c.Milk += milk
}

func (c *CoffeeMachine) GetCoffeeMachine() CoffeeMachine {
	return *c
}

///////////////////////////////////////////
///////////////////////////////////////////

type Driver struct {
	Name       string
	Age        int
	Experience int
	Car        Car
}

type Car struct {
	Brand string
	Model string
	Year  int
	Class Class
}

type Class struct {
	Name        string
	Coefficient float64
}

func AddDriver(name string, age, experience int, car Car) Driver {
	return Driver{
		Name:       name,
		Age:        age,
		Car:        car,
		Experience: experience,
	}
}

func (d *Driver) Drive() {
	fmt.Printf("Driver %s is driving\n", d.Name)
}

func (d *Driver) ChangeCar(brand, model string, year int, class Class) {
	d.Car = Car{
		Brand: brand,
		Model: model,
		Year:  year,
		Class: class,
	}
}

func (d *Driver) ChangeCoefficient(coefficient float64) {
	d.Car.Class.Coefficient = coefficient
}
