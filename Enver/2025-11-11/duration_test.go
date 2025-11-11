package Datentyp

import "fmt"

func ExampleTime_conversions() {
	a := FromSekunden(60)
	b := FromMinuten(60)
	c := FromStunden(60)

	fmt.Println(a.Minuten())
	fmt.Println(b.Minuten())
	fmt.Println(c.Minuten())
	fmt.Println(a.Sekunden())
	// Output:
	//
	//
	//
	//
}
