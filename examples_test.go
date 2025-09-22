package boolmap_test

import (
	"fmt"

	"vimagination.zapto.org/boolmap"
)

func ExampleMap() {
	m := boolmap.NewMap()

	m.SetBool(42, true)
	m.SetBool(100, false)

	fmt.Println(m.GetBool(42))
	fmt.Println(m.GetBool(100))
	fmt.Println(m.GetBool(101))

	// Output:
	// true
	// false
	// false
}

func ExampleSlice() {
	m := boolmap.NewSlice()

	m.SetBool(42, true)
	m.SetBool(100, false)

	fmt.Println(m.GetBool(42))
	fmt.Println(m.GetBool(100))
	fmt.Println(m.GetBool(101))

	// Output:
	// true
	// false
	// false
}

func ExampleCrumbMap() {
	m := boolmap.NewCrumbMap()

	m.Set(42, 1)
	m.Set(100, 3)

	fmt.Println(m.Get(42))
	fmt.Println(m.Get(100))
	fmt.Println(m.Get(101))

	// Output:
	// 1
	// 3
	// 0
}

func ExampleCrumbSlice() {
	m := boolmap.NewCrumbSlice()

	m.Set(42, 1)
	m.Set(100, 3)

	fmt.Println(m.Get(42))
	fmt.Println(m.Get(100))
	fmt.Println(m.Get(101))

	// Output:
	// 1
	// 3
	// 0
}

func ExampleNibbleMap() {
	m := boolmap.NewNibbleMap()

	m.Set(42, 11)
	m.Set(100, 13)

	fmt.Println(m.Get(42))
	fmt.Println(m.Get(100))
	fmt.Println(m.Get(101))

	// Output:
	// 11
	// 13
	// 0
}

func ExampleNibbleSlice() {
	m := boolmap.NewNibbleSlice()

	m.Set(42, 11)
	m.Set(100, 13)

	fmt.Println(m.Get(42))
	fmt.Println(m.Get(100))
	fmt.Println(m.Get(101))

	// Output:
	// 11
	// 13
	// 0
}
