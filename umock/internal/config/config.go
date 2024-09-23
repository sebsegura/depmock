package config

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	LatencyMS float64
	ErrorRate float64
)

func LoadConfig() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	latencyStr := os.Getenv("LATENCY_MS")
	if latencyStr != "" {
		if latency, err := strconv.ParseFloat(latencyStr, 64); err == nil {
			LatencyMS = latency
		}
	}

	errorRateStr := os.Getenv("ERROR_RATE")
	if errorRateStr != "" {
		if rate, err := strconv.ParseFloat(errorRateStr, 64); err == nil {
			ErrorRate = rate
		}
	}
}
