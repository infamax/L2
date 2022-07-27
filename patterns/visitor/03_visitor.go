package main

import (
	"fmt"
	"math"
)

/*

 */

type Shape interface {
	Name() string
	Accept(v Visitor)
	GetParams() []int
}

type Rectangle struct {
	a, b int
}

func (r *Rectangle) Name() string {
	return "rectangle"
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

func (r *Rectangle) GetParams() []int {
	return []int{r.a, r.b}
}

type Circle struct {
	r int
}

func (c *Circle) Name() string {
	return "circle"
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetParams() []int {
	return []int{c.r}
}

type Triangle struct {
	a, b, c int
}

func (t *Triangle) Name() string {
	return "triangle"
}

func (t *Triangle) Accept(v Visitor) {
	v.VisitForTriangle(t)
}

func (t *Triangle) GetParams() []int {
	return []int{t.a, t.b, t.c}
}

type Visitor interface {
	VisitForRectangle(r *Rectangle)
	VisitForCircle(c *Circle)
	VisitForTriangle(t *Triangle)
}

type SquareVisitor struct{}

func (sv *SquareVisitor) VisitForRectangle(r *Rectangle) {
	params := r.GetParams()
	fmt.Printf("square rectangle: %d\n", params[0]*params[1])
}

func (sv *SquareVisitor) VisitForCircle(c *Circle) {
	params := c.GetParams()
	fmt.Printf("square circle: %f\n", float64(params[0])*3.14)
}

func (sv *SquareVisitor) VisitForTriangle(t *Triangle) {
	params := t.GetParams()
	p := (params[0] + params[1] + params[2]) / 2
	s := math.Sqrt(float64(p) * float64(p-params[0]) * float64(p-params[1]) * float64(p-params[2]))
	fmt.Printf("square triangle: %f\n", s)
}

func main() {
	figures := []Shape{
		&Rectangle{
			a: 4,
			b: 5,
		},
		&Triangle{
			a: 3,
			b: 4,
			c: 5,
		},
		&Circle{
			r: 1,
		},
	}
	var sv *SquareVisitor
	for i := 0; i < len(figures); i++ {
		figures[i].Accept(sv)
	}
}
