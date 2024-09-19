package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	chart "toharia/charts"
	objects "toharia/model"
	calc "toharia/utility"

	"github.com/go-echarts/go-echarts/v2/components"
)

var (
	timeMu sync.RWMutex
)
var (
	//Game variables
	p = objects.Player{Health: 100, Seeds: 0, Level: 1, Exp: 0}
	g = objects.Game{FallowLand: 10, FertileLand: 0, NaturalWater: calc.RandRange(1, 10)}

	//Charting objects
	xVals = []string{"Seeds", "Food", "Water", "Level"}
)

func main() {

	var n string
	fmt.Println("Welcome to Toharia, a lost wilderness!")
	fmt.Println("What is your name?")
	//Get user input
	fmt.Scan(&n)
	p.UpdateName(n)

	p.Print()

	go runDataLoop(&p, &g)

	//Generate initial graphs
	http.HandleFunc("/", HttpServer)
	http.ListenAndServe(":8081", nil)

}

func runDataLoop(p *objects.Player, g *objects.Game) {
	repeat := true
	//Create main game loop
	for repeat {
		//Get user input
		repeat = updateGameState(p, g) //pass in dereferenced pointers
	}
}

func HttpServer(w http.ResponseWriter, _ *http.Request) {
	timeMu.RLock()
	defer timeMu.RUnlock()
	page := components.NewPage()

	//Health Guage
	gauge := chart.HealthGauge(&p)

	// create a new line instance
	bar := chart.InventoryBar(&p)

	// Add charts to page
	page.AddCharts(
		gauge,
		bar,
	)

	// Render charts to http server
	page.Render(w)
}

func updateGameState(p *objects.Player, g *objects.Game) bool {
	fmt.Println("==========================")
	fmt.Println("You have ", p.Seeds, " Seeds.")
	fmt.Println("You have ", p.Food, " Food.")
	fmt.Println("You have ", p.Water, " Water.")
	fmt.Println("You have ", g.FallowLand, " fallow land.")
	fmt.Println("You have ", g.FertileLand, " fertile land.")
	fmt.Println("You have ", p.Health, " Health.")
	fmt.Println("You have ", len(g.Crops), " Crops.")
	fmt.Println("==========================")

	fmt.Println("What would you like to do?")
	fmt.Println("1. Plant Seeds")
	fmt.Println("2. Harvest Crops")
	fmt.Println("3. Eat Crops")
	fmt.Println("4. Gather Seeds")
	fmt.Println("5. Gather Water")
	fmt.Println("6. Water Fields")
	fmt.Println("==========================")

	//Get user input
	var choice string
	fmt.Scan(&choice)

	keepPlaying := true

	fmt.Println("**************************")
	//Lock data structure
	timeMu.Lock()

	switch userChoice := choice; userChoice {
	case "1":
		keepPlaying = plantSeeds(p, g)
	case "2":
		keepPlaying = harvestCrops(p, g)
	case "3":
		keepPlaying = eatFood(p)
	case "4":
		keepPlaying = gatherSeeds(p)
	case "5":
		keepPlaying = gatherWater(p, g)
	case "6":
		keepPlaying = WaterFields(p, g)
	case "/lvl":
		fmt.Println("Level", p.Level)
		return true
	case "/xp":
		fmt.Println("Current XP:", p.Exp, "/100")
		return true
	case "/help":
		printHelp()
		return true
	case "/quit":
		os.Exit(1)
	default:
		fmt.Println("Invalid choice")
		return true
	}

	fmt.Println("**************************")

	//If there are Crops grow them a little
	g.IncrementAllCropGrowth()

	//Check for random world event
	randWorldEvent(g)

	//Health check
	p.SubtractHealth(1)

	if p.Health <= 0 {
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxx")
		fmt.Println("x", p.Name, "has died! x")
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxx")
		keepPlaying = false
	}

	//Unlock data structure to allow for update
	timeMu.Unlock()

	return keepPlaying
}

// Logic for planting Seeds
func plantSeeds(p *objects.Player, g *objects.Game) bool {
	if p.Seeds < 1 {
		fmt.Println("You don't have enough Seeds!")
		return true
	}
	p.SubtractSeeds(1)
	if g.FertileLand < 1 {
		fmt.Println("You don't have enough fertile land!")
		return true
	}
	g.SubtractFertileLand(1)
	g.AddCrop()
	fmt.Println("You used 1 seed and 1 fertile land!")
	return true
}

// Logic for harvesting Crops
func harvestCrops(p *objects.Player, g *objects.Game) bool {
	total := len(g.GetIndicesOfCropsToHarvest())
	if total > 0 {
		for _, index := range g.GetIndicesOfCropsToHarvest() {
			g.RemoveCrop(index)
			p.AddFood(1)
			fmt.Println("You harvested 1 crop! You now have", p.Food, "available!")

			g.AddFallowLand(1)
			fmt.Println("You regained 1 fallow land!", g.FallowLand, "is now available!")

			//Chance to get Seeds from Crops
			if calc.RandPercentageChance(50) {
				extraSeed := calc.RandRange(1, 3)
				p.AddSeeds(extraSeed)
				fmt.Println("You got", extraSeed, "seed(s) from your Crops!")
			}
		}
		fmt.Println("You harvested", total, "Crops total!")
		p.AddExp(5 * total)
	} else {
		fmt.Println("You don't have any Crops to harvest!")
	}

	return true
}

// Logic for eating
func eatFood(p *objects.Player) bool {
	if p.Food < 1 {
		fmt.Println("You don't have any Crops to eat!")
		return true
	}
	p.SubtractFood(1)
	p.AddHealth(calc.RandRange(1, 4))

	fmt.Println("You ate 1 crop! YUM! You now have", p.Food, "available!")
	fmt.Println("You now have ", p.Health, " Health!")

	return true
}

// Logic for gathering Seeds
func gatherSeeds(p *objects.Player) bool {
	s := calc.RandRange(0, 3)
	if s == 0 {
		fmt.Println("You didn't gather any Seeds!")
	} else {
		p.AddSeeds(s)
		fmt.Println("You gathered", s, "seed(s)! You now have", p.Seeds, "available!")
	}
	return true
}

// Logic for gathering Water from environment
func gatherWater(p *objects.Player, g *objects.Game) bool {
	if g.NaturalWater < 1 {
		fmt.Println("The land is in a drought! You don't have any Water to gather!")
		return true
	}
	//gather Water
	g.SubtractNaturalWater(1)
	p.AddWater(1)

	fmt.Println("You gathered 1 Water! You now have", p.Water, "available!")
	fmt.Println("There seems to be about", g.NaturalWater, "natural Water left in the environment!")

	return true
}

// Logic for Watering fields
func WaterFields(p *objects.Player, g *objects.Game) bool {
	if p.Water < 1 {
		fmt.Println("You don't have any Water to Water the fields!")
		return true
	}

	if g.FallowLand < 1 {
		fmt.Println("You don't have any fallow land to Water!")
		return true
	}

	p.SubtractWater(1)
	g.AddFertileLand(1)
	g.SubtractFallowLand(1)

	fmt.Println("You Watered 1 field! You now have", g.FertileLand, "fertile land available!")
	fmt.Println("You now have", p.Water, "Water available!")

	return true
}

// Logic for random world events
func randWorldEvent(g *objects.Game) {
	// Check if it rains
	if calc.RandPercentageChance(5) {
		fmt.Println("-----------It rained!-----------")

		g.AddNaturalWater(calc.RandRange(1, 5))

		fertilized := calc.RandRange(1, 3)
		g.AddFertileLand(fertilized)
		g.SubtractFallowLand(fertilized)
	}

	//Check if locusts
	if calc.RandPercentageChance(1) {
		fmt.Println("-----------A Swarm of Locusts decimate your Crops!!-----------")
		size := len(g.Crops)
		loss := calc.RandRange(1, size)

		for i := 0; i < loss; i++ {
			g.RemoveCrop(i)
		}
		fmt.Println("-----------", len(g.Crops), "Crops remaining...-----------")
	}

	//Check if growth spurt
	if calc.RandPercentageChance(10) && len(g.Crops) > 0 {
		fmt.Println("-----------One of your Crops grew exceptionally well!-----------")
		size := len(g.Crops)
		if size > 1 {
			growIndex := calc.RandRange(0, size-1)
			g.Crops[growIndex].AddCropGrowth(1)
		} else {
			g.Crops[0].AddCropGrowth(1)
		}
	}

}

//Help Menu

func printHelp() {
	fmt.Println("/lvl: Show current Level")
	fmt.Print("/xp: Show current Experience")
	fmt.Println("/help: Show help menu")
	fmt.Println("/quit: Quit objects.Game")
}
