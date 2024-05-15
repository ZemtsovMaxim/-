package pattern

/*
	Реализовать паттерн «стратегия».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
	после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
	Паттерн Стратегия предлагает определить семейство схожих алгоритмов, которые часто изменяются или расширяются, и вынести их в собственные классы, называемые стратегиями.

	+
	Горячая замена алгоритмов на лету.
 	Изолирует код и данные алгоритмов от остальных классов.
 	Уход от наследования к делегированию.
 	Реализует принцип открытости/закрытости.

	-
	Усложняет программу за счёт дополнительных классов.
 	Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

import "fmt"

// Strategy определяет общий интерфейс для всех стратегий
type Strategy interface {
	DoOperation(int, int) int
}

// ConcreteStrategyAdd - конкретная стратегия, реализующая операцию сложения
type ConcreteStrategyAdd struct{}

func (csa *ConcreteStrategyAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

// ConcreteStrategySubtract - конкретная стратегия, реализующая операцию вычитания
type ConcreteStrategySubtract struct{}

func (css *ConcreteStrategySubtract) DoOperation(num1, num2 int) int {
	return num1 - num2
}

// Context - контекст, который использует стратегию
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

// Пример использования

func ExampleUsage() {
	// Создаем контекст с конкретной стратегией сложения
	context := NewContext(&ConcreteStrategyAdd{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Println("Результат сложения:", result)

	// Меняем стратегию на вычитание
	context = NewContext(&ConcreteStrategySubtract{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Println("Результат вычитания:", result)
}
