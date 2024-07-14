package views

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Page interface {
	Render()
}
type Data struct {
	Nodes []opts.GraphNode
	Links []opts.GraphLink
}
type Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
type Experience struct {
	Name     string `json:"name"`
	Practice int    `json:"practice"`
	Theory   int    `json:"theory"`
	Link     []Link `json:"links"`
}
type Experiences struct {
	Experiences []Experience `json:"experiences"`
}
type ExpWebGraph struct{}

func fileToExp(input string) Experiences {
	bytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	var exps Experiences

	json.Unmarshal(bytes, &exps)

	//assign the Name of the current Experience as the Source in each of it's Links
	for e := range exps.Experiences {
		for l := range exps.Experiences[e].Link {
			exps.Experiences[e].Link[l].Source = exps.Experiences[e].Name
		}
	}

	return exps
}

func base() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Experience Web"}),
	)

	exp := fileToExp("tmp/experience.json")
	data := expToNode(exp)
	fmt.Printf("data: %v\n", data)

	// var data Data
	// loadDataFromFile(&data, "fixtures/experience-plottable.json")

	//add file data to graph
	graph.AddSeries("graph", data.Nodes, data.Links).
		SetSeriesOptions(
			//layout/gravity rules
			charts.WithGraphChartOpts(opts.GraphChart{
				Layout: "force",
				Force: &opts.GraphForce{
					InitLayout: "",
					Repulsion:  2000,
					Gravity:    0.5,
					EdgeLength: 0,
				},
				Roam:               opts.Bool(true),
				EdgeSymbol:         nil,
				EdgeSymbolSize:     nil,
				Draggable:          new(bool),
				FocusNodeAdjacency: opts.Bool(true),
				Categories:         []*opts.GraphCategory{},
				EdgeLabel: &opts.EdgeLabel{
					Show:          new(bool),
					Position:      "",
					Color:         "",
					FontStyle:     "",
					FontWeight:    nil,
					FontSize:      0,
					Align:         "",
					VerticalAlign: "",
					Padding:       nil,
					Width:         0,
					Height:        0,
					Formatter:     "",
				},
				SymbolKeepAspect: new(bool),
			}),
			//line defs
			charts.WithLineStyleOpts(opts.LineStyle{
				Color:     "",
				Width:     0,
				Type:      "",
				Opacity:   0,
				Curveness: 0,
			}),
			//labels
			charts.WithLabelOpts(
				opts.Label{Show: opts.Bool(true)},
			),
		)
	return graph
}

func expToNode(exp Experiences) Data {
	var data Data
	for e := range exp.Experiences {
		data.Nodes = append(
			data.Nodes,
			opts.GraphNode{
				Name:       exp.Experiences[e].Name,
				Value:      0,
				Symbol:     "circle",
				SymbolSize: sizeFromPractice(exp.Experiences[e].Practice),
				ItemStyle: &opts.ItemStyle{
					Color: colorFromTheory(exp.Experiences[e].Theory),
				},
				Tooltip: &opts.Tooltip{
					Show: opts.Bool(true),
				},
			},
		)
		for l := range exp.Experiences[e].Link {
			data.Links = append(
				data.Links,
				opts.GraphLink{
					Source: exp.Experiences[e].Link[l].Source,
					Target: exp.Experiences[e].Link[l].Target,
					Value:  0,
					Label:  &opts.EdgeLabel{},
				},
			)

		}
	}
	return data
}

func sizeFromPractice(i int) float32 {
	return (float32(i)) * 2
}

func colorFromTheory(i int) string {
	fmt.Printf("Theory->color(%d)", i)
	red := 255.0
	green := 0.0
	step := 7.5

	for range i {
		if green < 255 {
			green += step
			if green > 255.0 {
				fmt.Println("forcing green to 255")
				green = 255.0
			}
		}
		if red > 0 {
			red -= step
			if red < 0 {
				red = 0
				fmt.Println("forcing red to 0")
			}
		}
	}

	fmt.Printf("red: %f, green: %f, blue: %f", red, green, 0.0)
	return opts.RGBAColor(uint16(red), uint16(green), 0, 1)
}

func loadDataFromFile(data *Data, file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(f, &data); err != nil {
		fmt.Println(err)
	}
}

func (ExpWebGraph) Render() {
	page := components.NewPage()
	page.AddCharts(
		base(),
	)

	f, err := os.Create("examples/html/sample.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}
