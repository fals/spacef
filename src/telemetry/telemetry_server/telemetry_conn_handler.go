package telemetry_server

import 
(
    "net"
    "fmt"
    "os"
    "bytes"
)

type ITelemetryConnHandler interface {
	HandleClient(conn net.Conn)
}

type TelemetryConnHandler struct {

}

func CreateHandler() (*TelemetryConnHandler) {
	return &TelemetryConnHandler{}
}

func (h *TelemetryConnHandler) HandleClient(conn net.Conn) {
    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err != nil {
            fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
            return
        }
        
        message := HandleMessage(result.Bytes());

        result.Reset()

        fmt.Println(message)
    }
}