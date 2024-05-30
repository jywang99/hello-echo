package main

import (
	"html/template"
	"io"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
    Templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
    return t.Templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates{
        Templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

var id = 0
type Contact struct {
    Id int
    Name string
    Email string
}

func newContact(name, email string) Contact {
    id ++
    return Contact{
        Id: id,
        Name: name,
        Email: email,
    }
}

type Contacts = []Contact

type Data struct {
    Contacts Contacts
}

func (d *Data) indexOf(id int) int {
    for i, contact := range d.Contacts {
        if contact.Id == id {
            return i
        }
    }
    return -1
}

func newData() Data {
    return Data{
        Contacts: Contacts{
            newContact("John", "jd@gmail.com"),
            newContact("Clara", "cd@gmail.com"),
        },
    }
}

func (d *Data) hasEmail(email string) bool {
    for _, contact := range d.Contacts {
        if contact.Email == email {
            return true
        }
    }
    return false
}

type FormData struct {
    Values map[string]string
    Errors map[string]string
}

func newFormData() FormData {
    return FormData{
        Values: make(map[string]string),
        Errors: make(map[string]string),
    }
}

type Page struct {
    Data Data
    Form FormData
}

func newPage() Page {
    return Page{
        Data: newData(),
        Form: newFormData(),
    }
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = newTemplate()
    e.Static("/static", "static")

    page := newPage()
    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", page)
    })

    e.POST("/contacts", func(c echo.Context) error {
        name := c.FormValue("name")
        email := c.FormValue("email")

        if page.Data.hasEmail(email) {
            data := newFormData()
            data.Values["name"] = name
            data.Values["email"] = email
            data.Errors["email"] = "Email already exists"

            return c.Render(422, "form", data)
        }

        contact := newContact(name, email)
        page.Data.Contacts = append(page.Data.Contacts, contact)

        c.Render(200, "form", newFormData())
        return c.Render(200, "oob-contact", contact)
    })

    e.DELETE("/contacts/:id", func(c echo.Context) error {
        time.Sleep(3 * time.Second)
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            return c.String(400, "Invalid ID")
        }

        idx := page.Data.indexOf(id)
        if idx == -1 {
            return c.String(404, "Contact not found")
        }

        page.Data.Contacts = append(page.Data.Contacts[:idx], page.Data.Contacts[idx+1:]...)
        return c.NoContent(200)
    })

    e.Logger.Fatal(e.Start(":42069"))
}

