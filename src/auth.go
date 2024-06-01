package main

type User struct {
    Id int
    Email string
    Password string
}

func newUser(email, password string) User {
    return User{
        Id: 0,
        Email: email,
        Password: password,
    }
}

var users = []User{
    newUser("taro@example.com", "password"),
}

func (u *User) authenticate(email, password string) bool {
    for _, user := range users {
        if user.Email == email && user.Password == password {
            return true
        }
    }
    return false
}

