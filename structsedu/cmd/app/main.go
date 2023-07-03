package main

func main() {
	//person := Person{
	//	Name:    "John Doe",
	//	Age:     30,
	//	Address: "123 Main St",
	//}
	//car := Car{
	//	Brand: "Toyota",
	//	Model: "Camry",
	//	Year:  2020,
	//}
	//fmt.Println(person.Name)
	//person.Name = "Jack Doe"
	//fmt.Println(person.Name)
	//fmt.Println(car)

	/////////////////////////////

	//file, err := os.Open("personal.json")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//fileBytes, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(fileBytes), "----")

	/////////////////////////////

	//var p Person
	//err = json.Unmarshal(fileBytes, &p)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(p)

	/////////////////////////////

	//var p2 Person
	//p2.Name = "John Doe"
	//p2.Age = 30
	//p2.Address = "123 Main St"
	//d, err := json.Marshal(p2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//newFile, err := os.Create("personal2.json")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer newFile.Close()
	//_, err = newFile.Write(d)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Car struct {
	Brand string
	Model string
	Year  int
}
