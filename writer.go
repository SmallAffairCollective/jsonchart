package main

import (
	"html"
	"io/ioutil"
	"os"
	"strconv"
)

func writeGChartHTML() {
	const s = `<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script><script type="text/javascript" src="chart.js"></script><div id="chart_div" style="width: 1200px; height: 800px"></div>`
	content := []byte(html.UnescapeString(s))

	_ = os.Mkdir("www", 0755)
	err := ioutil.WriteFile("www/index.html", content, 0644)
	check(err)
}

func writeGChartJs(title string, delay int, iterations int, data map[string][]float64) {

	const s = `google.charts.load('current', {packages: ['corechart', 'line']});
google.charts.setOnLoadCallback(drawCurveTypes);

function drawCurveTypes() {
	var data = new google.visualization.DataTable();
	data.addColumn('number', 'X');`

	content := []byte(html.UnescapeString(s))

	dataRows := [][]float64{}
	counter := 0
	delayTime := delay - 1
	for counter < iterations {
		row := []float64{float64(delayTime)}
		for field := range data {
			row = append(row, float64(data[field][counter]))
		}
		dataRows = append(dataRows, row)
		delayTime += delay
		counter++
	}

	dataStr := ""
	for _, row := range dataRows {
		dataStr += "["
		for _, value := range row {
			dataStr += strconv.FormatFloat(value, 'f', -1, 64) + ", "
		}
		dataStr += "], "
	}
	dataStr = dataStr[:len(dataStr)-1]

	columnData := ""
	for field := range data {
		columnData += "\n\tdata.addColumn('number', '" + field + "');"
	}

	columnData += "\n"

	// write header stuff
	_ = os.Mkdir("www", 0755)
	err := ioutil.WriteFile("www/chart.js", content, 0644)
	check(err)

	// open up file again for appending more stuff
	f, err := os.OpenFile("www/chart.js", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	// write field columns
	_, err = f.WriteString(columnData)
	check(err)

	// write actual value data
	values := "\n\n    data.addRows([" + dataStr + "])\n\n"
	_, err = f.WriteString(values)
	check(err)

	// write footer stuff
	const footer = `    var options = {
        title: '`
	const mfooter = `',
        hAxis: {
          title: 'Time'
        },
        vAxis: {
          title: ''
        },
        series: {
          1: {curveType: 'function'}
        }
      };

      var chart = new google.visualization.LineChart(document.getElementById('chart_div'));
      chart.draw(data, options);
	}`

	// what the...
	_, err = f.WriteString(footer)
	check(err)
	_, err = f.WriteString(title)
	check(err)
	_, err = f.WriteString(mfooter)
	check(err)
}
