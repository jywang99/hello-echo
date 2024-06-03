package main

import "github.com/labstack/echo/v4"

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Age int `json:"age"`
    Region string `json:"region"`
    Gender string
}

type UserData struct {
    Users []User
}

func newUsers() UserData {
    return UserData{
        Users: []User{
            {
                Name: "John",
                Email: "john@example.com",
                Age: 20,
                Region: "US",
                Gender: "Male",
            },
            {
                Name: "Jonathan",
                Email: "jonathan@example.com",
                Age: 25,
                Region: "US",
                Gender: "Male",
            },
            {
                Name: "Clara",
                Email: "clara@example.com",
                Age: 25,
                Region: "UK",
                Gender: "Female",
            },
            {
                Name: "Eva",
                Email: "eva@example.com",
                Age: 20,
                Region: "UK",
                Gender: "Female",
            },
        },
    }
}

type UserParmams struct {
    Age int `query:"age"`
    Region string `query:"region"`
    Gender string `query:"gender"`
}

func (ud UserData) find(params UserParmams) []User {
    if params.Age == 0 && params.Region == "" && params.Gender == "" {
        return ud.Users
    }

    users := make([]User, 0)
    for _, user := range ud.Users {
        if params.Age != 0 && user.Age != params.Age {
            continue
        }
        if params.Region != "" && user.Region != params.Region {
            continue
        }
        if params.Gender != "" && user.Gender != params.Gender {
            continue
        }
        users = append(users, user)
    }
    return users
}

func SetupUserQuery(e *echo.Echo) {
    e.GET("/users", func(c echo.Context) error {
        users := newUsers()
        var params UserParmams
        if err := c.Bind(&params); err != nil {
            return err
        }
        return c.JSONPretty(200, users.find(params), "  ")
    })
}

