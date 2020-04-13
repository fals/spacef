package telemetry_server

import 
(
	"net"
	"strconv"
	"os"
)

type TelemetryServer struct {
	address string
    listener *net.TCPListener
}

func Connect(port int) (*TelemetryServer, error) {
	address := ":" + strconv.Itoa(port)
	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", 	address)
	
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	
	if err != nil {
		return nil, err
	}

	return &TelemetryServer {address, listener}, nil
}

func (server *TelemetryServer) HandleConnections(handler ITelemetryConnHandler) {
	for {
        conn, err := server.listener.Accept()
        if err != nil {
            continue
		}
		
        go handler.HandleClient(conn)
	}
}

func (server *TelemetryServer) Close()  {
	server.listener.Close()
	os.Exit(0);
}