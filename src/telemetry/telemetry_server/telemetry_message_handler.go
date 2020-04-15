package telemetry_server

import (
	"strings"
	"strconv"
)

func HandleMessage(bytes []byte) (*TelemetryMessage) {
	parsedBytes := string(bytes)
	messageParts := strings.Split(parsedBytes, "&")

	return &TelemetryMessage {
		getStringValue(messageParts[0]), 
		getStringValue(messageParts[1]), 
		getStringValue(messageParts[2]), 
		getFloar64Value(messageParts[3]), 
		getUint64Value(messageParts[4])}
}

func getStringValue(part string) (string) {
	return part[3:len(part)]
}

func getFloar64Value(part string) (float64) {
	temp := part[3:len(part)]
	v, _ := strconv.ParseFloat(temp, 64)
	
	return v
}

func getUint64Value(part string) (uint64) {
	temp := part[3:len(part)]
	v, _ := strconv.ParseUint(temp, 0, 64)

	return v
}