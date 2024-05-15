package pattern

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

	Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
	Фасад — это простой интерфейс для работы со сложной подсистемой, содержащей множество классов. Фасад может иметь урезанный интерфейс, не имеющий 100% функциональности,
	которой можно достичь, используя сложную подсистему напрямую. Но он предоставляет именно те фичи, которые нужны клиенту, и скрывает все остальные.
	Фасад предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы. Фасад определяет интерфейс более высокого уровня, который упрощает использование подсистемы.

	 + Изолирует клиентов от компонентов сложной подсистемы.
	 - Фасад рискует стать божественным объектом, привязанным ко всем классам программы.

*/

import "fmt"

type subsystemA struct{}

func (s *subsystemA) operationA() {
	fmt.Println("Subsystem A: operation")
}

type subsystemB struct{}

func (s *subsystemB) operationB() {
	fmt.Println("Subsystem B: operation")
}

// Фасад
type facade struct {
	subA *subsystemA
	subB *subsystemB
}

func newFacade() *facade {
	return &facade{
		subA: &subsystemA{},
		subB: &subsystemB{},
	}
}

// Методы Фасада делегируют вызовы к соответствующим методам подсистем
func (f *facade) operation1() {
	fmt.Println("Facade: operation 1")
	f.subA.operationA()
	f.subB.operationB()
}

func (f *facade) operation2() {
	fmt.Println("Facade: operation 2")
	f.subB.operationB()
}

// Пример использования
func HowToUseFacade() {
	f := newFacade()

	f.operation1()
	f.operation2()
}
