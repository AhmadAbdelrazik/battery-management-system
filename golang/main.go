package main

import (
	// "math"
	"fmt"
	"math"
	"os"
	"time"
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
	Zk := 1.0    // State of Charge

	battery := NewBattery(R0, R1, R2, C1, C2, Dt, Ni, Cn, Zk)
	kalman := NewKalman(battery, data)

	i := 0
	fmt.Printf("\n\nCalculating SoC using Kalman Filter with constant current of 3A and maximum Voltage = 4.2V\n")
	fmt.Printf("----------------------------------------------------------------------------------------\n\n")
	time.Sleep(3 * time.Second)
	for {
		soc, v := kalman.UpdateTest(3)

		if i%2000 == 0 {
			fmt.Printf("Iteration number: %6v\t|   ", i)
			fmt.Printf("State of Charge (SoC): %.2f%%\t|   ", soc*100)
			fmt.Printf("Terminal Voltage: %.3fV\n", v)
			time.Sleep(1 * time.Second)
		}
		if soc < 0.02 {
			break
		}
		i++
	}
}
