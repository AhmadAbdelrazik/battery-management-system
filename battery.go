package main

// all resistors are in milli Ohms
// all Capacitors are in Kilo Farads
type Battery struct {
	R0 float64
	R1 float64
	R2 float64
	C1 float64
	C2 float64
	Dt float64
	Ni float64 // Coulombic Efficiency, 1 at discharge, n at charge
	Cn float64 // Nominal Capacity
	Zk float64 // State of Charge
}

func (b *Battery) getI1() float64 {
	return 0
}
func (b *Battery) getI2() float64 {
	return 0
}
