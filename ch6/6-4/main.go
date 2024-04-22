package main

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}

	return f, nil
}
