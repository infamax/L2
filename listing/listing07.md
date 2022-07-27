Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
В любом порядке выведется последовательность чисел
от 1 до 8. А затем бесконечное число нулей
```

Объяснение: Так как конкурентные вычисления не гарантируют порядок их вывода, то
мы получим последовательность чисел от 1 до 8 в случайном порядке. Поскольку в функции
merge организован бесконечный цикл, то дальше будет получение нулевых значений из закрытых каналов. Нулевое
значение типа int это ноль
