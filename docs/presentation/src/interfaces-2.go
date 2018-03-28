package main

import "fmt"

type AvisiWriter int

func (a *AvisiWriter) Write(p []byte) (n int, err error) {
	wrote := len(p)
	*a = AvisiWriter(int(*a) + wrote)
	return n, nil
}

func main() {
	var writer AvisiWriter = 0

	fmt.Fprintf(&writer, "This is something random")
	fmt.Fprintf(&writer, "Even something more")
	fmt.Fprintf(&writer, "🙈")

	fmt.Println(writer)
}
