package main

import (
	"fmt"
)

type User struct {
	name string
	mail string
}

func main() {

	// value
	u1 := User{}
	u1.name, u1.mail = "name_u1", "u1@user.com"

	// reference
	u2 := &User{}
	u2.name, u2.mail = "name_u2", "u2@user.com"

	// reference
	u3 := new(User)
	u3.name, u3.mail = "name_u3", "u3@user.com"

	// reference and init: One advantage of using the &Type{} syntax for structs is that
	// we can specify initial field values as we did here when creating the augusta pointer
	u4 := &User{"name_u4", "u4@user.com"}

	fmt.Println("Hello struct ", u1, u2, u3, u4)
}
