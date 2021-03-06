package telemetry_server
	
import (
    "testing"
    "os"
)


var  messageInBytes = []byte("id=sensor_number&tp=temperature&un=C&vl=23.4&ts=1586983044290")
func TestMain(m *testing.M) {
    os.Exit(m.Run())
}

func Test_Message_Handle_Should_Return_Message_With_Id(t *testing.T)  {
    message := HandleMessage(messageInBytes)

    if message.id != "sensor_number" {
        t.Errorf("Wrong Parsed Value for id %s", message.id)
        os.Exit(-1)
    }
}

func Test_Message_Handle_Should_Return_Message_With_Type(t *testing.T)  {
    message := HandleMessage(messageInBytes)

    if message.sensor_type != "temperature" {
        t.Errorf("Wrong Parsed Value for sensor_type %s", message.sensor_type)
        os.Exit(-1)
    }
}

func Test_Message_Handle_Should_Return_Message_With_Unit(t *testing.T)  {
    message := HandleMessage(messageInBytes)

    if message.unit != "C" {
        t.Errorf("Wrong Parsed Value for unit %s", message.unit)
        os.Exit(-1)
    }
}

func Test_Message_Handle_Should_Return_Message_With_Value(t *testing.T)  {
    message := HandleMessage(messageInBytes)

    if message.value != 23.4 {
        t.Errorf("Wrong Parsed Value for value %f", message.value)
        os.Exit(-1)
    }
}

func Test_Message_Handle_Should_Return_Message_With_Timestamp(t *testing.T)  {
    message := HandleMessage(messageInBytes)

    if message.timestamp != 1586983044290 {
        t.Errorf("Wrong Parsed Value for timestamp %d", message.timestamp)
        os.Exit(-1)
    }
}