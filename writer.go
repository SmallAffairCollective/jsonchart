package main

import (
	"html"
	"io/ioutil"
)

func writeGChartHtml() {
	const s = `<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script><div id="chart_div"></div>`
	content := []byte(html.UnescapeString(s))

	err := ioutil.WriteFile("www/index.html", content, 0644)
	check(err)
}

func writeGChartJs(title string, yLabel string, data map[string][]float64) {

}
