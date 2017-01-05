package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p) - 1, nil
}

func main() {
	var c ByteCounter
	count, err := c.Write([]byte("hello"))
	if err != nil {
		fmt.Println("no way! ", count)
	}
	fmt.Println(c)

	c = 0
	fmt.Println(c)
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

}
