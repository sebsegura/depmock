package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"math/rand"
	"sebsegura/umock/internal/config"
	"time"
)

func getOperationConfig(operation *openapi3.Operation) (float64, float64) {
	latencyMS := config.LatencyMS
	errorRate := config.ErrorRate

	if latencyExt, ok := operation.Extensions["x-latency-ms"]; ok {
		if latencyMSValue, ok := latencyExt.(float64); ok {
			latencyMS = latencyMSValue
		}
	}

	if errorRateExt, ok := operation.Extensions["x-error-rate"]; ok {
		if errorRateValue, ok := errorRateExt.(float64); ok {
			errorRate = errorRateValue
		}
	}

	return latencyMS, errorRate
}

func applyLatency(latencyMS float64) {
	if latencyMS > 0 {
		time.Sleep(time.Duration(latencyMS) * time.Millisecond)
	}
}

func shouldInjectError(errorRate float64) bool {
	if errorRate <= 0 {
		return false
	}
	return rand.Float64() < errorRate
}
