package main

import (
	"html/template"
	"io"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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

var page = newPage()

func getContacts(c echo.Context) error {
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
}

func deleteContact(c echo.Context) error {
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
}

