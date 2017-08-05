package main

import (
	"strconv"
	"time"

	"github.com/mediocregopher/radix.v2/redis"
)

func getMetrics(jsonMap map[string]interface{}) map[string]float64 {

	metrics := make(map[string]float64)

	for key, item := range jsonMap {
		if item != nil {
			if value, ok := item.(float64); ok {
				metrics[key] = value
			} else if value, ok := item.(string); ok {
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
