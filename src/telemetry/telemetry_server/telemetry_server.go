package telemetry_server

import 
(
	"net"
	"strconv"
)

func Connect(port int) error {
	service := ":" + strconv.Itoa(port)
	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	
	if err != nil {
		return err
	}

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
		}
		
        go handleClient(conn)
	}
	
	return nil
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }

        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }
}