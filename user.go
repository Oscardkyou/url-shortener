package main

import (
"fmt"
)

type ShortenerResponse struct {
ShortURL string `json:"shortened"`
}

type User struct {
name string // Dauren, Azamat, Sergey
}

func (u *User) GetName() string { // getter
return [u.name](http://u.name/)
}

func (u *User) SetName(n string) error { // setter mutator
if len(n) < 3 {
return fmt.Errorf("Name is too short")
}
if n == "Azamat" || n == "Dauren" || n == "Sergey" {
return fmt.Errorf("Name is not allowed")
}
[u.name](http://u.name/) = n
return nil
}

func main() {
var user User
user.SetName("Vasya") // Example of using the setter
fmt.Println(user.GetName()) // Example of using the getter
}
