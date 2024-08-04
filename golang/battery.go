package main

// all resistors are in milli Ohms
// all Capacitors are in Kilo Farads
type Battery struct {
	R0 float64
	R1 float64
	R2 float64
	C1 float64
	C2 float64
	I1 float64
	I2 float64
	I  float64
	Dt float64
	Ni float64 // Coulombic Efficiency, 1 at discharge, n at charge
	Cn float64 // Nominal Capacity
	Zk float64 // State of Charge
}

// UpdateCurrents: Updates I, I1, I2 based on thevenin's 2nd-order circuit
// I: measured current
func (b *Battery) UpdateCurrents(I float64) {
	b.I = I
	dI1 := (1 / (b.R1 * b.C1)) * (I - b.I1) * b.Dt
	dI2 := (1 / (b.R2 * b.C2)) * (I - b.I2) * b.Dt

	b.I1 += dI1
	b.I1 += dI2
}

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
