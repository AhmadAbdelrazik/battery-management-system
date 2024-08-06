package main

type Battery struct {
	R0 float64
	R1 float64
	R2 float64
	C1 float64
	C2 float64
	I1 float64
	I2 float64
	I  float64
	Dt float64 // Time step
	Ni float64 // Coulombic Efficiency, 1 at discharge, Ni at charge
	Cn float64 // Nominal Capacity
	Zk float64 // State of Charge
}

// Initializes a new Battery.
func NewBattery(R0, R1, R2, C1, C2, Dt, Ni, Cn, Zk float64) *Battery {
	return &Battery{
		R0: R0,
		R1: R1,
		R2: R2,
		C1: C1,
		C2: C2,
		Dt: Dt,
		Ni: Ni,
		Cn: Cn,
		Zk: Zk,
	}
}
