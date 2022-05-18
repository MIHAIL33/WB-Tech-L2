package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Engine struct {}

func (e *Engine) Start() {
	fmt.Println("Engine is on")
}

func (e *Engine) Stop() {
	fmt.Println("Engine is off")
}

type Chassis struct {}

func (c *Chassis) Up() {
	fmt.Println("Chassis is up")
}

func (c *Chassis) Down() {
	fmt.Println("Chassis is down")
}

type AirPlaneFacade struct {
	engine *Engine
	chassis *Chassis
}

func NewAirPlaneFacade() *AirPlaneFacade {
	return &AirPlaneFacade{ new(Engine), new(Chassis) }
}

func (a *AirPlaneFacade) takkingOff() {
	fmt.Println("Airplane is taking off")
}

func (a *AirPlaneFacade) slowingDown() {
	fmt.Println("Airplane is slowing down")
}

func (a *AirPlaneFacade) Up() {
	a.engine.Start()
	a.takkingOff()
	a.chassis.Up()
}

func (a *AirPlaneFacade) Down() {
	a.chassis.Down()
	a.slowingDown()
	a.engine.Stop()
}

/*
	Использование:
	1) для доступа к сложной системе требуется простой интерфейс
	2) система очень сложна или трудна для понимания
	3) точка входа необходима для каждого уровня многоуровневого программного обеспечения
	4) абстракции и реализации подсистемы тесно свяаны

	+:
	1) изолирует клиентов от компонентов сложной подсистемы

	-:
	1) рискует стать антипаттерном "божественный объект"
*/