package telemetry_server

import 
(
	"net"
)

type ITelemetryConnHandler interface {
	HandleClient(conn net.Conn)
}

type TelemetryConnHandler struct {

}

func CreateHandler() (*TelemetryConnHandler) {
	return &TelemetryConnHandler{}
}

func (h *TelemetryConnHandler)HandleClient(conn net.Conn) {
    defer conn.Close()

    var buf [512]byte
    for {
        _, err := conn.Read(buf[0:])
        if err != nil {
            return
		}
    }
}