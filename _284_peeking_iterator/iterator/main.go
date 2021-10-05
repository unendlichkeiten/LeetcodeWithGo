// iterator demo
package main

import "fmt"

func main() {
	l := List{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	i := l.Iterator()
	for i.HasNext() {
		x := i.Value().(int)
		fmt.Println(x)
		i.Next()
	}
}
