package gocmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	expected := Person{
		Name: "Dennis",
		Age:  37,
	}

	result := CreatePerson("Dennis", 37)

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}

	// if diff := cmp.Diff(expected, result); diff != "" {
	// 	t.Error(diff)
	// }
}
