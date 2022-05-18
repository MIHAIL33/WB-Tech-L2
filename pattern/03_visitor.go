package pattern

import "math"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}


type Rectangle struct {
	Height int
	Width int
}

func (r *Rectangle) Accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

type AreaCalculator struct {
	Area float64
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	a.Area = float64(s.Side * s.Side)
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	a.Area = float64(c.Radius * c.Radius) * math.Pi
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	a.Area = float64(r.Height * r.Width)
}

/*
	Использование:
	1) когда нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов (дерево)
	2) когда не нужно засорять классы несвязанными операциями
	3) когда новое поведение имеет смысл только для некоторых классов из существующей иерархии

	+:
	1) упрощает добавление операций, работающих со сложными структурами объектов
	2) объединение родственных операций в одном классе
	3) посетитель может накапливать состояние при обходе структуры элементов

	-:
	1) паттерн не оправдан, если иерархия элементов часто меняется
	2) может привести к нарушению инкапсуляции элементов
*/