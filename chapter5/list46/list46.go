package main

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	//构造器？
	duration(15).pretty()
	//.\list46.go:13:14: cannot call pointer method on duration(15)
	//.\list46.go:13:14: cannot take the address of duration(15)
}

// Values  Methods Receivers
// T        (t T)
// *T       (t T) and (t *T)

// Methods Receivers        Values
// (t T)                    T and *T
// (t *T)                   *T
