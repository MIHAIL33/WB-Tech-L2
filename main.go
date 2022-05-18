package main

import (
	"fmt"

	"github.com/MIHAIL33/WB-TECH-L2/pattern"
)

func main() {

	//Facade pattern
	fmt.Println("///////////////////////////////Facade")
	airplaneFacade := pattern.NewAirPlaneFacade()

	airplaneFacade.Up()
	airplaneFacade.Down()

	//Builder pattern
	fmt.Println("///////////////////////////////Builder")
	director1 := &pattern.Director{ Builder: pattern.NewSkodaBuilder() }
	avto1 := director1.Assemble()
	fmt.Println(avto1.String())

	director2 := &pattern.Director{ Builder: pattern.NewAudiBuilder() }
	avto2 := director2.Assemble()
	fmt.Println(avto2.String())

	//Visitor pattern
	fmt.Println("///////////////////////////////Visitor")
	square := &pattern.Square{ Side: 2 }
	circle := &pattern.Circle{ Radius: 3 }
	rectangle := &pattern.Rectangle{ Height: 3, Width: 4 }

	areaCalculator := &pattern.AreaCalculator{}
	square.Accept(areaCalculator)
	fmt.Println(areaCalculator.Area)
	circle.Accept(areaCalculator)
	fmt.Println(areaCalculator.Area)
	rectangle.Accept(areaCalculator)
	fmt.Println(areaCalculator.Area)

	//Command pattern
	fmt.Println("///////////////////////////////Command")
	
}