package telemetry_server
	
import (
    "testing"
)

func Should_Be_Able_To_Connect(t *testing.T)  {
    err := Connect(15002)

    if err != nil {
        t.Errorf("Can't connect %s", err.Error())
    }
}