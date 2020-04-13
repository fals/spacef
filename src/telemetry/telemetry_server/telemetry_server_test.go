package telemetry_server
	
import (
    "testing"
)

func Test_Should_Be_Able_To_Connect(t *testing.T)  {
    server, err := Connect(1)
    server.Close();
    
    if err != nil {
        t.Errorf("Can't connect %s", err.Error())
    }
}