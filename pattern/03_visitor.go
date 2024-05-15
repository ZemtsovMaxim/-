package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Посетитель — это поведенческий паттерн проектирования,который позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.

	Паттерн Посетитель предлагает разместить новое поведение в отдельном классе, вместо того чтобы множить его сразу в нескольких классах.
	Объекты, с которыми должно было быть связано поведение, не будут выполнять его самостоятельно. Вместо этого вы будете передавать эти объекты в методы посетителя.

	+
	 Упрощает добавление операций, работающих со сложными структурами объектов.
 	Объединяет родственные операции в одном классе.
 	Посетитель может накапливать состояние при обходе структуры элементов.

	-
 	Паттерн не оправдан, если иерархия элементов часто меняется.
 	Может привести к нарушению инкапсуляции элементов.
*/

import "fmt"

// Element интерфейс, который должны реализовывать все элементы, доступные для посещения
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElement1 и ConcreteElement2 - конкретные элементы, реализующие интерфейс Element
type ConcreteElement1 struct{}

func (ce *ConcreteElement1) Accept(visitor Visitor) {
	visitor.VisitConcreteElement1(ce)
}

type ConcreteElement2 struct{}

func (ce *ConcreteElement2) Accept(visitor Visitor) {
	visitor.VisitConcreteElement2(ce)
}

// Visitor интерфейс, определяющий методы для посещения каждого типа элемента
type Visitor interface {
	VisitConcreteElement1(ce *ConcreteElement1)
	VisitConcreteElement2(ce *ConcreteElement2)
}

// ConcreteVisitor реализует интерфейс Visitor и содержит реализации посещения каждого элемента
type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitConcreteElement1(ce *ConcreteElement1) {
	fmt.Println("Visit ConcreteElement1")
}

func (cv *ConcreteVisitor) VisitConcreteElement2(ce *ConcreteElement2) {
	fmt.Println("Visit ConcreteElement2")
}

// Пример использования

func ExampleUsage() {
	// Создаем экземпляры элементов
	element1 := &ConcreteElement1{}
	element2 := &ConcreteElement2{}

	// Создаем экземпляр посетителя
	visitor := &ConcreteVisitor{}

	// Элементы принимают посетителя
	element1.Accept(visitor)
	element2.Accept(visitor)
}
