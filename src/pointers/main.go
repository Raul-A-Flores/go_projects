package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func createPointer() *int {
	x := 42
	return &x
}

func main() {

	user := User{ID: 42, Name: "Raul"}
	println(user.Name)

	userId := 43

	println(userId)

	println(&userId)

	anotherUserId := &userId
	println(anotherUserId)

	userId = 100
	println(anotherUserId)
	print(*anotherUserId)

	var age *int

	ages := &userId

	p := createPointer()
	fmt.Println(*p)

	println(age)
	println(*ages)

	update(ages, 42)
	println(*ages)
}

func update(val *int, to int) {
	*val = to

}
