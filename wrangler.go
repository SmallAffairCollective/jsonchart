package main

import "strconv"

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
