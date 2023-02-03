package main

import (
	"fmt"
	"math"
)

// User represents a user in a social network
type User struct {
	Name  string
	Email string
}

// NewUser creates a new User instance with the given name and email.
// The function returns a pointer to the newly created User instance.
func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

// UpdateName updates the name of the user to the given name.
func (u *User) UpdateName(name string) {
	u.Name = name
}

// Notify sends a notification to the user by printing a message to the console.
// The message includes the user's name and email address.
func (u *User) Notify() {
	fmt.Printf("Notifying %s <%s>\n", u.Name, u.Email)
}

// Distance calculates the distance between two points in a two-dimensional plane.
// The method takes two Point instances as arguments and returns the calculated distance.
func (u *User) Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	user := NewUser("John Doe", "johndoe@example.com")
	user.UpdateName("Jane Doe")
	user.Notify()

	distance := user.Distance(0, 0, 3, 4)
	fmt.Printf("Distance between (0, 0) and (3, 4) is %f\n", distance)
}
