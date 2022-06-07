package main

import (
	"fmt"
	"os"

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

	tv := &pattern.TV{}
	onCommand := &pattern.OnCommand{Device: tv}
	offCommand := &pattern.OffCommand{Device: tv}

	onButton := &pattern.Button{Command: onCommand}
	onButton.Press()
	offButton := &pattern.Button{Command: offCommand}
	offButton.Press()

	//Chain of resp pattern
	fmt.Println("///////////////////////////////Chain of resp")

	brokenCar := pattern.NewBrokenCar("Lada", true, false, true)

	engineMaster := &pattern.EngineMaster{}

	wiringMaster := &pattern.WiringMaster{}
	wiringMaster.SetNext(engineMaster)

	wheelsMaster := &pattern.WheelsMaster{}
	wheelsMaster.SetNext(wiringMaster)

	wheelsMaster.Execute(brokenCar)

	//Factory method
	fmt.Println("///////////////////////////////Factory method")

	bmw, err := pattern.GetAvto("BMW")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Println(bmw.GetName(), bmw.GetEngine())

	shkoda, err := pattern.GetAvto("Shkoda")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Println(shkoda.GetName(), shkoda.GetEngine())

	//Strategy
	fmt.Println("///////////////////////////////Strategy")

	sell30 := func() float64 {
		return 0.7
	}

	sell50 := func() float64 {
		return 0.5
	}

	price := pattern.Price{Amount: 100, Discount: sell30}
	price.Sell()
	fmt.Println(price.FinalPrice)
	price.SetStrategy(sell50)
	price.Sell()
	fmt.Println(price.FinalPrice)
	

	//State
	fmt.Println("///////////////////////////////State")

	context := pattern.Context{State: new(pattern.StateA)}
	context.Request()

	context.SetState(new(pattern.StateB))
	context.Request()
}	