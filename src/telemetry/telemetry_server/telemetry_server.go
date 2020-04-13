package telemetry_server

import 
(
	"net"
	"strconv"
	"os"
	"fmt"
)

type TelemetryServer struct {
	address string
    listener *net.TCPListener
}

func Connect(port int) (*TelemetryServer, error) {
	address := ":" + strconv.Itoa(port)
	
	tcpAddr, err := net.ResolveTCPAddr("tcp", 	address)
	
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
		fmt.Println("Waiting connections...")
        conn, err := server.listener.Accept()
        if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
            continue
		}

		go handler.HandleClient(conn)
		fmt.Println("One connection handled closed")
	}
}

func (server *TelemetryServer) Close()  {
	server.listener.Close()
	os.Exit(0);
}
