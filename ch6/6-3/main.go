package main

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}
