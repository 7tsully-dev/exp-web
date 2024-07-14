package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var nodes = []opts.GraphNode{
	{Name: "Node1", SymbolSize: 100, X: 0, Y: 0,
		ItemStyle: &opts.ItemStyle{
			Color:    "green",
			GapWidth: 0.5,
		},
	},
	{Name: "Node2", SymbolSize: 10},
	{Name: "Node3", SymbolSize: 10},
	{Name: "Node4", SymbolSize: 10},
	{Name: "Node5", SymbolSize: 10},
	{Name: "Node6", SymbolSize: 10},
	{Name: "Node7", SymbolSize: 10},
	{Name: "Node8", SymbolSize: 10},
}

func base() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "web graph example"}),
	)

	f, err := os.ReadFile("fixtures/sample.json")
	if err != nil {
		panic(err)
	}

	type Data struct {
		Nodes []opts.GraphNode
		Links []opts.GraphLink
	}

	//load from file
	var data Data
	if err := json.Unmarshal(f, &data); err != nil {
		fmt.Println(err)
	}

	//add file data to graph
	graph.AddSeries("graph", data.Nodes, data.Links).
		SetSeriesOptions(
			//layout/gravity rules
			charts.WithGraphChartOpts(opts.GraphChart{
				Layout:             "circular",
				Roam:               opts.Bool(true),
				FocusNodeAdjacency: opts.Bool(true),
				Force:              &opts.GraphForce{Repulsion: 1000, Gravity: 0.5},
			}),
			//line defs
			charts.WithLineStyleOpts(opts.LineStyle{
				Curveness: 0.3,
			}),
			//labels
			charts.WithLabelOpts(
				opts.Label{Show: opts.Bool(true)},
			),
		)
	return graph
}

type WebGraphExamples struct{}

func (WebGraphExamples) Examples() {
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
