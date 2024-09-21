package chart

import (
	objects "toharia/model"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	//Charting objects
	xVals = []string{"Seeds", "Food", "Water", "Level", "Fallow Land", "Fertile Land", "Crops"}
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

func InventoryBar(p *objects.Player, g *objects.Game) *charts.Bar {
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

	fallow := make([]opts.BarData, len(xVals))
	fallow[4] = opts.BarData{Value: g.FallowLand}

	fertile := make([]opts.BarData, len(xVals))
	fertile[5] = opts.BarData{Value: g.FertileLand}

	crops := make([]opts.BarData, len(xVals))
	crops[6] = opts.BarData{Value: len(g.Crops)}

	bar.SetXAxis(xVals).
		AddSeries("Seeds", seeds).
		AddSeries("Food", food).
		AddSeries("Water", water).
		AddSeries("Level", level).
		AddSeries("Fallow Land", fallow).
		AddSeries("Fertile Land", fertile).
		AddSeries("Crops", crops)

	return bar
}

func ExpPieRoseArea(p *objects.Player) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Experience",
		}),
	)

	values := make([]opts.PieData, 2)

	currentXP := p.Exp
	remainingXP := 100 - p.Exp

	values[0] = opts.PieData{Name: "Remaining XP", Value: remainingXP}
	values[1] = opts.PieData{Name: "Current XP", Value: currentXP}

	pie.AddSeries("Exp", values).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	return pie
}
