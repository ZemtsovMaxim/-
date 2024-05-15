package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

	Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
	Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

	+
 	Позволяет создавать продукты пошагово.
 	Позволяет использовать один и тот же код для создания различных продуктов.
 	Изолирует сложный код сборки продукта от его основной бизнес-логики.

	-
	Усложняет код программы из-за введения дополнительных классов.
 	Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.

*/

// Product представляет объект, который мы собираемся создать
type Product struct {
	Part1 string
	Part2 int
	Part3 bool
}

// Builder определяет интерфейс для создания объекта Product
type Builder interface {
	BuildPart1() Builder
	BuildPart2() Builder
	BuildPart3() Builder
	GetProduct() Product
}

// ConcreteBuilder реализует интерфейс Builder и строит объект Product
type ConcreteBuilder struct {
	product Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{}
}

func (cb *ConcreteBuilder) BuildPart1() Builder {
	cb.product.Part1 = "part1"
	return cb
}

func (cb *ConcreteBuilder) BuildPart2() Builder {
	cb.product.Part2 = 42
	return cb
}

func (cb *ConcreteBuilder) BuildPart3() Builder {
	cb.product.Part3 = true
	return cb
}

func (cb *ConcreteBuilder) GetProduct() Product {
	return cb.product
}

// Director управляет процессом построения объекта Product
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() Product {
	return d.builder.BuildPart1().BuildPart2().BuildPart3().GetProduct()
}
