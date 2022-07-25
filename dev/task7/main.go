package main

import (
	"fmt"
	"reflect"
	"time"
)

func or(ch ...<-chan interface{}) (r chan interface{}) {
	res := make(chan interface{})
	var s []reflect.SelectCase
	for _, c := range ch {
		//Если Dir выбран по умолчанию, регистр представляет регистр по умолчанию.
		//Если выбран параметр Dir, регистр представляет операцию отправки.
		//Если Dir имеет значение SelectRecv, регистр представляет операцию приема.
		s = append(s, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		})
	}
	go func() {
		//Select выполняет операцию выбора, описанную в списке обращений. Как и оператор Go select,
		//он блокируется до тех пор, пока не будет выполнено хотя бы одно из обращений,
		//выполняется равномерный псевдослучайный выбор, а затем выполняется это обращение.
		reflect.Select(s)
		close(res)
	}()
	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(10*time.Second),
	)

	fmt.Printf("fone after %v\n", time.Since(start))

}
