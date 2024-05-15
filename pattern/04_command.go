package pattern

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

	Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов,
	ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

	+
	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
 	Позволяет реализовать простую отмену и повтор операций.
 	Позволяет реализовать отложенный запуск операций.
 	Позволяет собирать сложные команды из простых.
 	Реализует принцип открытости/закрытости.

	-
	Усложняет код программы из-за введения множества дополнительных классов.
*/

import "fmt"

// Command определяет общий интерфейс для всех команд
type Command interface {
	Execute()
}

// Receiver - получатель, который знает, как выполнять действие
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: выполнение действия")
}

// ConcreteCommand - конкретная команда, которая связывает операцию с получателем
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (cc *ConcreteCommand) Execute() {
	fmt.Println("ConcreteCommand: выполнение команды")
	cc.receiver.Action()
}

// Invoker - инициатор, который выполняет команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	fmt.Println("Invoker: выполнение команды")
	i.command.Execute()
}

// Пример использования

func ExampleUsage() {
	// Создаем получателя
	receiver := &Receiver{}

	// Создаем команду и связываем ее с получателем
	command := NewConcreteCommand(receiver)

	// Создаем инициатор и устанавливаем команду
	invoker := &Invoker{}
	invoker.SetCommand(command)

	// Инициатор запускает команду
	invoker.ExecuteCommand()
}
