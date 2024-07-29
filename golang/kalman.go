package main

import "math"

type Kalman struct {
	Battery *Battery
	Pk      [][]float64
	Xk      [][]float64
	Kk      [][]float64
	Fk      [][]float64
	Bk      [][]float64
	SigmaWk [][]float64
	SigmaVk float64
	Yk      float64
}

func (k *Kalman) Init(b *Battery) {
	// Initial State
	k.Xk = MatT([][]float64{{1, 0, 0}})
	// Initial Error Covariance
	k.Pk = [][]float64{
		{0.005, 0, 0},
		{0, 0.0001, 0},
		{0, 0, 0.0001},
	}
	// Process Jacobian
	k.Fk = [][]float64{
		{1, 0, 0},
		{0, math.Exp(-k.Battery.Dt / (k.Battery.R1 * k.Battery.C1)), 0},
		{0, 0, math.Exp(-k.Battery.Dt / (k.Battery.R2 * k.Battery.C2))},
	}

	k.Bk = [][]float64{
		{-k.Battery.Ni * k.Battery.Dt / (3600 * k.Battery.Cn)},
		{1 - math.Exp(-k.Battery.Dt/(k.Battery.R1*k.Battery.C1))},
		{1 - math.Exp(-k.Battery.Dt/(k.Battery.R2*k.Battery.C2))},
	}

	k.Battery = b
	k.SigmaWk = MatMul(k.Bk, MatT(k.Bk))
	k.SigmaVk = 1
}

// Hk
func (k *Kalman) GetHk() [][]float64 {
	var delta float64
	zk := k.Xk[0][0] // State of charge
	if zk != 0 {
		delta = (OCV(zk) - OCV(zk-0.5)) / 0.5
	} else {
		delta = (OCV(zk+0.5) - OCV(zk)) / 0.5
	}
	Hk := [][]float64{
		{delta, -k.Battery.R1, -k.Battery.R2},
	}

	return Hk
}

func (k *Kalman) StepOne(ik float64) {
	k.Xk = MatAdd(MatMul(k.Fk, k.Xk), MatMulC(k.Bk, ik))
}

func (k *Kalman) StepTwo() {
	k.Pk = MatAdd(MatMul(MatMul(k.Fk, k.Pk), MatT(k.Fk)), k.SigmaWk)
}

// Predict System Output
func (k *Kalman) StepThree(ik float64) {
	k.Yk = OCV(k.Battery.Zk) - k.Battery.R0*ik - k.Battery.R1*k.Battery.getI1() - k.Battery.R2*k.Battery.getI2()
}

func (k *Kalman) StepFour() {
	hk := k.GetHk()

	HKxPKxHKt := MatMul(MatMul(hk, k.Pk), MatT(hk))
	SigmaY := MatDiffuse(HKxPKxHKt) + k.SigmaVk

	k.Kk = MatMulC(MatMul(k.Pk, hk), 1/SigmaY)
}

func (k *Kalman) StepFive(outputVoltage float64) {
	k.Xk = MatAdd(k.Xk, MatMulC(k.Kk, outputVoltage-k.Yk))
}

func (k *Kalman) StepSix() {
	k.Pk = MatMul(MatSub(MatI(3), MatMul(k.Kk, k.GetHk())), k.Pk)
}
