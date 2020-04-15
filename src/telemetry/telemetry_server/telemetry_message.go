package telemetry_server

type TelemetryMessage struct {
	id string
	sensor_type string
	unit string
	value float64
	timestamp uint64
}

