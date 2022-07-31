package serial

import (
	"testing"
)

func TestReadWrite(t *testing.T) {

	config1 := Config{
		Address:  "COM3",
		BaudRate: 9600,
		DataBits: 8,
		Parity:   "N",
		StopBits: 1,
	}
	port1, err := Open(&config1)
	if err != nil {
		t.Fatal(err)
	} else {
		port1.Write([]byte("HELLO-GO"))
	}
	defer port1.Close()
}
