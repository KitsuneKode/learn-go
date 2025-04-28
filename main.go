package main

import (
	"fmt"
	"learnGo/exercise"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	i := []exercise.Item{
		{Name: "Sword", Type: "Weapon"},
		{Name: "Health Potion", Type: "potion"},
	}
	p := exercise.Player{Name: "John Doe", Inventory: i}

	fmt.Println("Player details", p)

	p.PickUpItem(exercise.Item{Name: "Health Potion", Type: "potion"})

	p.DropItem(exercise.Item{Name: "Sword", Type: "Weapon"})
	p.DropItem(exercise.Item{Name: "Shield", Type: "weapon"})

	p.UseItem(exercise.Item{Name: "Health Potion", Type: "potion"})
}

func learn2() {
	learn1()
	person := Person{Name: "Manish", Age: 27}
	fmt.Printf("This is our person %+v\n", person)

	employee := struct {
		id   int
		Name string
	}{
		Name: "Alice",
		id:   2034034,
	}

	fmt.Printf("This is our employee %+v\n", employee)

	type Address struct {
		city   string
		state  string
		street string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "John Doe",
		Address: Address{
			street: "123 Main St",
			city:   "New York",
			state:  "NY",
		},
		Phone: "123-456-7890",
	}

	fmt.Printf("This is our contact %+v\n", contact)
	fmt.Println("Before modifying", person)
	person.modifyPersonName("JOhn mitchell")

	fmt.Println("After modifying", person)
}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
}

func learn1() {
	fmt.Println("Hello, World!")

	age := 27
	fmt.Println("Age:", age)
	result := add(5, 3)
	fmt.Println("Result of addition:", result)
	sum, product := calculatProductandSum(5, 3)
	fmt.Println("Sum:", sum, product)
	conditonal(age)

	for i := 0; i < 3; i++ {
		fmt.Println("Loop iteration:", i)
	}

	for {
		fmt.Println("Infinite loop")
		if age < 30 {
			break
		}
	}

	numbers := [4]int{10, 20, 30, 40}
	fmt.Println("Array:", numbers)
	fmt.Println("Array length:", len(numbers))

	number2 := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("2D Array:", number2)

	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Slice:", fruits)

	fruits = append(fruits, "orange")
	fmt.Println("Slice after append:", fruits)

	morefruits := []string{"strawberry", "blueberry", "mango"}

	fruits = append(fruits, morefruits...)
	fmt.Println("Slice after appending another slice:", fruits)

	for _, value := range numbers {
		fmt.Println("Value:", value)
	}

	capitalCities := map[string]string{
		"USA":   "Washington",
		"India": "New Delhi",
		"UK":    "London",
	}

	fmt.Println("Map: ", capitalCities)

	capital, exists := capitalCities["France"]

	if exists {
		fmt.Println("Capital of India:", capital)
	} else {
		fmt.Println("Capital not found")
	}
}

func add(a int, b int) int {
	return a + b
}

func calculatProductandSum(a, b int) (int, int) {
	return a + b, a * b
}

func conditonal(age int) {
	if age < 30 {
		fmt.Println("Age is less than 30")
	}
}
