# go serial port
## Example
```go
package main

import (
	"log"
	"github.com/goburrow/serial"
)

func main() {
	port, err := serial.Open(&serial.Config{Address: "/dev/ttyUSB0"})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	_, err = port.Write([]byte("ok"))
	if err != nil {
		log.Fatal(err)
	}
}
```
## Testing

### Linux
socat
```sh
sudo apt install socat
socat -d -d pty,raw,echo=0 pty,raw,echo=0
```
### Windows
- [Null-modem emulator](http://com0com.sourceforge.net/)
- [Terminal](https://sites.google.com/site/terminalbpp/)
