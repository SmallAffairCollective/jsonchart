package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/mediocregopher/radix.v2/redis"
	"github.com/tokuhirom/json_path_scanner"
)

func getMetrics(jsonMap map[string]interface{}) map[string]float64 {

	metrics := make(map[string]float64)

	ch := make(chan *json_path_scanner.PathValue)
	go func() {
		// TODO try/catch if fails check if it's flat, otherwise fail
		json_path_scanner.Scan(jsonMap, ch)
	}()

	for p := range ch {
		if p.Value != nil {
			key := p.Path[2:len(p.Path)]
			// format path to make it cleaner
			key = strings.Replace(key, "[", ".", -1)
			key = strings.Replace(key, "]", "", -1)
			key = strings.Replace(key, "'", "", -1)
			if value, ok := p.Value.(float64); ok {
				metrics[key] = value
			} else if value, ok := p.Value.(string); ok {
				if v, err := strconv.Atoi(value); err == nil {
					metrics[key] = float64(v)
				}
			}
		}
	}

	return metrics

}

func storeMetrics(url string, conn *redis.Client, metrics map[string]float64) bool {

	unixTime := time.Now().Unix()

	resp := conn.Cmd("SADD", url, strconv.FormatInt(unixTime, 10))
	if resp.Err != nil {
		return false
	}

	resp = conn.Cmd("HMSET", url+strconv.FormatInt(unixTime, 10), metrics)
	if resp.Err != nil {
		return false
	}

	return true

}

func getStoredMetricMatrix(url string, conn *redis.Client) map[string]map[string]float64 {

	resp := conn.Cmd("SMEMBERS", url)
	check(resp.Err)

	l, _ := resp.List()
	matrix := make(map[string]map[string]float64)

	for _, item := range l {
		resp = conn.Cmd("HGETALL", url+item)
		check(resp.Err)

		matrix[url+item] = make(map[string]float64)
		m, _ := resp.Map()
		for key := range m {
			value, _ := strconv.ParseFloat(m[key], 64)
			matrix[url+item][key] = value
		}

	}

	return matrix
}

func flattenMetricMatrix(url string, matrix map[string]map[string]float64) map[string]map[string][]float64 {

	flattenedMatrix := make(map[string]map[string][]float64)
	flattenedMatrix[url] = make(map[string][]float64)

	for item := range matrix {
		if strings.HasPrefix(item, url) {
			for field := range matrix[item] {
				flattenedMatrix[url][field] = append(flattenedMatrix[url][field], matrix[item][field])
			}
		}
	}
	return flattenedMatrix
}
