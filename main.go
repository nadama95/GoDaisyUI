package main

import "fmt"

func main() {
	p := ButtonParams{Variant: Variant(Primary)}
	btn := Button(p)

	fmt.Println(btn.Render())
}
