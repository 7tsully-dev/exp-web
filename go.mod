module github.com/go-echarts/examples

go 1.22.5

require (
	github.com/7tsully-dev/exp-web/views v0.0.0-00010101000000-000000000000
	github.com/go-echarts/go-echarts/v2 v2.4.0
)

replace github.com/7tsully-dev/exp-web/views => ./examples

// dev mode
//replace github.com/go-echarts/go-echarts/v2 => ../../go-echarts
//
//replace github.com/go-echarts/snapshot-chromedp => ../snapshot-chromedp
