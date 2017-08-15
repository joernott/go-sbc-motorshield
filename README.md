# Library for PiMotor Shield V2

- Developed by: Ott-Consult UG (haftungsbeschränkt)
- Author: Jörn Ott
- Project: RPi Motor Shield
- Copyright: (c) 2017 Joern Ott
- License: Apache License, Version 2.0 (see LICENSE)

## Usage
```
    import (
		"fmt"
		"time"
		"github.com/joernott/go-sbc-motorshield/motor"
	)
	
func main() {
	fmt.Println("Initializing")
	Motor, err := motor.NewMotor()
	if err != nil {
		panic(err)
	}
	defer Motor.CloseMotor()
	Motor.ArrowOn("MOTOR1")
	Motor.Forward("MOTOR1")
	time.Sleep(1 * time.Second)
	Motor.Reverse("MOTOR1")
	time.Sleep(4 * time.Second)
	Motor.ArrowOff("MOTOR1")
	Motor.Stop("MOTOR1")
}
```

