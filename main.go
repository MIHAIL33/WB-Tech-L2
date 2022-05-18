package main

import "github.com/MIHAIL33/WB-TECH-L2/pattern"

func main() {

	//Facade pattern
	airplaneFacade := pattern.NewAirPlaneFacade()

	airplaneFacade.Up()
	airplaneFacade.Down()

}