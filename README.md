# Golang Interface

Golang Interface might be confusing for a lot of people that just started doing programming, including Me. I was quite struggling trying to understand it. Until.. I think of it in simple way. Let's talk about it.

## Interface as Profession

I found it easy to understand Interface by look at it as a Profession. For example a Teacher, Soccer Player, Racer, Boxer, etc. Why? Because both Interface and Profession require some set of criterias. A Boxer need to be able to throw Jab, Hook, Upprecut, etc.

Let's turn Boxer to an Interface.

```go
type Boxer interface{}
```

We just create a Boxer interface. Now, let's adds some set of criteria.

```go
type Boxer interface{
    jab()
    hook()
    uppercut()
}
```

What's all of those things mean? It means, for someone to be able or qualify to fill the role of a Boxer, they need to able to jab, hook, and uppercut.

They can't fill the role if they lack or not be able to do all of those criterias.

## Why do we need Interface?

We need interface to make things more versatile and not strict to one entity. For example if we make a storage module and we only take MySQL, when there's changes needed we need to update the code. If we use an interface, we don't need to do that. Because interface only specify the criteria, not the entity itself. Let's take a look at the example.

```go
type MySQL struct

func storage(datastore MySQL) {

}
```

The storage function above only works for datastore with type `MySQL`. So if you got datastore with type `PostgreSQL` you cannot use that function, you would probably need to create another storage function specific for `PostgreSQL`. Pretty annoying right?

```go
type Databse interface{
    Create()
    Read()
    Update()
    Delete()
}


type MySQL struct
func(m MySQL) Create(){}
func(m MySQL) Read(){}
func(m MySQL) Update(){}
func(m MySQL) Delete(){}

type PostgreSQL struct
func(p PostgreSQL) Create(){}
func(p PostgreSQL) Read(){}
func(p PostgreSQL) Update(){}
func(p PostgreSQL) Delete(){}

func storage(datastore Database) {

}
```

The example code above shows the storage function using interface. By doing this way, you could use any type of datastore as an argument to storage function, as long as they implement or satisfy the criteria of `Database` interface.

Let's write some code, and checks if it works or not.

```go
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
```

It works, you can use 2 different type of datastore on the same storage function thanks to interface feature.
