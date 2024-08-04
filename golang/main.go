package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	inputRaw, err := os.ReadFile("ocv_vs_soc.csv")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	data := NewData(string(inputRaw))

	R0 := 2.1552 * math.Pow10(-3)
	R1 := 6.7587 * math.Pow10(-3)
	R2 := 2.5683 * math.Pow10(-3)
	C1 := 123.1912 * math.Pow10(3)
	C2 := 4.5240 * math.Pow10(3)
	Dt := 0.1    // Time step in seconds
	Ni := 0.9894 // Coulombic Efficiency
	Cn := 6.8080 // Nominal Capacity
	Zk := 1.0    // State of Charge, 1.0 == 100%

	battery := NewBattery(R0, R1, R2, C1, C2, Dt, Ni, Cn, Zk)

	kalman := NewKalman(battery, data)

	measuredCurrent := 5.0 // in Amperes
	measuredVoltage := 4.0 // in Voltage

	// Usage: For each measurement iteration (which is equal to Dt)
	// you should call the Cycle function and pass the measured Current
	// and the measured voltage

	kalman.Cycle(measuredCurrent, measuredVoltage)
}
