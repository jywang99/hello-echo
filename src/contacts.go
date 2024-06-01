package main

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

