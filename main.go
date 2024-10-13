package main

import "fmt"

type Database interface {
	Create()
	Read()
	Update()
	Delete()
}

type PostgreSQL struct {
	name string
	port int
}

func (p PostgreSQL) Create() {
	fmt.Printf("Create %v on port %v\n", p.name, p.port)
}

func (p PostgreSQL) Read() {
	fmt.Printf("Read %v on port %v\n", p.name, p.port)
}

func (p PostgreSQL) Update() {
	fmt.Printf("Update %v on port %v\n", p.name, p.port)
}

func (p PostgreSQL) Delete() {
	fmt.Printf("Delete %v on port %v\n", p.name, p.port)
}

type MySQL struct {
	name string
	port int
}

func (m MySQL) Create() {
	fmt.Printf("Create %v on port %v\n", m.name, m.port)
}

func (m MySQL) Read() {
	fmt.Printf("Read %v on port %v\n", m.name, m.port)
}

func (m MySQL) Update() {
	fmt.Printf("Update %v on port %v\n", m.name, m.port)
}

func (m MySQL) Delete() {
	fmt.Printf("Delete %v on port %v\n", m.name, m.port)
}

func storage(datastore Database) {
	datastore.Create()
	datastore.Read()
	datastore.Update()
	datastore.Delete()
}

func main() {
	pg := PostgreSQL{
		name: "connection 1",
		port: 5432,
	}

	ms := MySQL{
		name: "connection 2",
		port: 3306,
	}

	storage(pg)
	fmt.Println()
	fmt.Println()
	storage(ms)
}
