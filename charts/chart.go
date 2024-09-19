package chart

import (
	objects "toharia/model"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	//Charting objects
	xVals = []string{"Seeds", "Food", "Water", "Level"}
)

func HealthGauge(p *objects.Player) *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Health"}),
	)

	gaugeTitle := p.Name + "'s Health"
	gauge.AddSeries(gaugeTitle, []opts.GaugeData{{Name: "Current Health", Value: p.Health}})

	return gauge
}

func InventoryBar(p *objects.Player) *charts.Bar {
	// create a new line instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Player Inventory", Subtitle: "This is how much stuff you have."}),
	)

	seeds := make([]opts.BarData, len(xVals))
	seeds[0] = opts.BarData{Value: p.Seeds}

	food := make([]opts.BarData, len(xVals))
	food[1] = opts.BarData{Value: p.Food}

	water := make([]opts.BarData, len(xVals))
	water[2] = opts.BarData{Value: p.Water}

	level := make([]opts.BarData, len(xVals))
	level[3] = opts.BarData{Value: p.Level}

	bar.SetXAxis(xVals).
		AddSeries("Seeds", seeds).
		AddSeries("Food", food).
		AddSeries("Water", water).
		AddSeries("Level", level)

	return bar
}
