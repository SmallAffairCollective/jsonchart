package main

import (
	"strconv"

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

func storeMetrics(url string, redisHost string, metrics map[string]float64) bool {

	conn, er := redis.Dial("tcp", redisHost+":6379")
	check(er)

	defer conn.Close()

	resp := conn.Cmd("HMSET", url, metrics)

	if resp.Err != nil {
		return false
	}

	return true

}
