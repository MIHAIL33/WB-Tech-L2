package pattern

import "strconv"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Avto struct {
	name string
	price int
}

func (a *Avto) String() string {
	return "Avto [name = " + a.name + ", price = " + strconv.Itoa(a.price) + "]"
}

type Builder interface {
	SetName()
	SetPrice()
	GetAvto() * Avto
}

type Director struct {
	Builder Builder
}

func (d *Director) Assemble() *Avto {
	d.Builder.SetName()
	d.Builder.SetPrice()
	return d.Builder.GetAvto()
}

type SkodaBuilder struct {
	avto *Avto
}

func NewSkodaBuilder() *SkodaBuilder {
	return &SkodaBuilder{ new(Avto) }
}

func (s *SkodaBuilder) SetName() {
	s.avto.name = "skoda"
}

func (s *SkodaBuilder) SetPrice() {
	s.avto.price = 10000000
}

func (s *SkodaBuilder) GetAvto() *Avto {
	return s.avto
}

type AudiBuilder struct {
	avto *Avto
}

func NewAudiBuilder() *AudiBuilder {
	return &AudiBuilder{ new(Avto) }
}

func (a *AudiBuilder) SetName() {
	a.avto.name = "audi"
}

func (a *AudiBuilder) SetPrice() {
	a.avto.price = 15000000
}

func (a *AudiBuilder) GetAvto() *Avto {
	return a.avto
}

/*
	Использование:
	1) построение сложного объекта от его представления

	+:
	1) позволяет изменить внутреннее представление продукта
	2) инкапсулирует код для построения и представления
	3) обеспечивает контроль за этапами процесса строительства

	-:
	1) для каждого типа продукта должен быть создан отдельный строитель
	2) классы строителя должны быть изменяемыми
	3) может затруднить внедрение зависимостей
*/