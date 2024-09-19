package objects

import (
	"fmt"
	calc "toharia/utility"
)

type Player struct {
	Name   string
	Seeds  int
	Food   int
	Health int
	Water  int
	Exp    int
	Level  int
}

type Game struct {
	FallowLand   int
	FertileLand  int
	NaturalWater int
	Crops        []Crop
}

type Crop struct {
	GrowthTime     int
	MaturationTime int
}

func (p *Player) UpdateName(newName string) {
	(*p).Name = newName
}

func (p *Player) AddSeeds(number int) {
	(*p).Seeds += number
}

func (p *Player) SubtractSeeds(number int) {
	(*p).Seeds -= number
}

func (p *Player) AddFood(number int) {
	(*p).Food += number
}

func (p *Player) SubtractFood(number int) {
	(*p).Food -= number
}

func (p *Player) AddHealth(number int) {
	(*p).Health += number
}

func (p *Player) SubtractHealth(number int) {
	(*p).Health -= number
}

func (p *Player) AddWater(number int) {
	(*p).Water += number
}

func (p *Player) SubtractWater(number int) {
	(*p).Water -= number
}

func (p *Player) AddLevel(number int) {
	(*p).Level += number
}

func (p *Player) AddExp(number int) {
	fmt.Println("Gained", number, "Experience points!")
	(*p).Exp += number
	if (*p).Exp >= 100 {
		(*p).SubtractExp(100)
		(*p).AddLevel(1)
		fmt.Println("Leveled up! You are now Level", (*p).Level, ".")
	}
}

func (p *Player) SubtractExp(number int) {
	(*p).Exp -= number
}

func (g *Game) AddFallowLand(number int) {
	(*g).FallowLand += number
}

func (g *Game) SubtractFallowLand(number int) {
	(*g).FallowLand -= number
}

func (g *Game) AddFertileLand(number int) {
	(*g).FertileLand += number
}

func (g *Game) SubtractFertileLand(number int) {
	(*g).FertileLand -= number
}

func (g *Game) AddNaturalWater(number int) {
	(*g).NaturalWater += number
}

func (g *Game) SubtractNaturalWater(number int) {
	(*g).NaturalWater -= number
}

func (g *Game) AddCrop() {
	//Add crop to Game
	(*g).Crops = append((*g).Crops, Crop{GrowthTime: 0, MaturationTime: calc.RandRange(5, 10)})
}

func (g *Game) RemoveCrop(index int) {
	(*g).Crops = append((*g).Crops[:index], (*g).Crops[index+1:]...)
}

func (g *Game) IncrementAllCropGrowth() {
	for i := 0; i < len((*g).Crops); i++ {
		(*g).Crops[i].AddCropGrowth(1)
	}
}

func (g *Game) GetIndicesOfCropsToHarvest() []int {
	var indices []int
	for i := 0; i < len((*g).Crops); i++ {
		if (*g).Crops[i].GrowthTime >= (*g).Crops[i].MaturationTime {
			indices = append(indices, i)
		}
	}
	return indices
}

func (c *Crop) AddCropGrowth(amount int) {
	(*c).GrowthTime += amount
}

func (c *Crop) SubtractCropGrowth(amount int) {
	(*c).GrowthTime -= amount
}

// receiver function
func (p Player) Print() {
	fmt.Printf("%+v", p)
}
