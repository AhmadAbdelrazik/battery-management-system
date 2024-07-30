package main

import "math"

type Kalman struct {
	Battery *Battery
	Data    SoCOCV
	Pk      [][]float64
	Xk      [][]float64
	Kk      [][]float64
	Fk      [][]float64
	Bk      [][]float64
	SigmaWk [][]float64
	SigmaVk float64
	Yk      float64
}

func Init(b *Battery, d SoCOCV) Kalman {
	var k Kalman
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

	k.Data = d
	return k
}

// Hk
func (k *Kalman) GetHk() [][]float64 {
	Hk := [][]float64{
		{k.Data.derivative(k.Xk[0][0]), -k.Battery.R1, -k.Battery.R2},
	}

	return Hk
}

func (k *Kalman) StepOne() {
	k.Xk = MatAdd(MatMul(k.Fk, k.Xk), MatMulC(k.Bk, k.Battery.I))
}

func (k *Kalman) StepTwo() {
	k.Pk = MatAdd(MatMul(MatMul(k.Fk, k.Pk), MatT(k.Fk)), k.SigmaWk)
}

// Predict System Output
func (k *Kalman) StepThree() {
	k.Yk = k.Data.GetVoltage(k.Battery.Zk) - k.Battery.R0*k.Battery.I - k.Battery.R1*k.Battery.I1 - k.Battery.R2*k.Battery.I2
}

func (k *Kalman) StepFour() {
	hk := k.GetHk()

	HKxPKxHKt := MatMul(MatMul(hk, k.Pk), MatT(hk))
	SigmaY := MatDiffuse(HKxPKxHKt) + k.SigmaVk

	k.Kk = MatMulC(MatMul(k.Pk, hk), 1/SigmaY)
}

func (k *Kalman) StepFive(V float64) {
	k.Xk = MatAdd(k.Xk, MatMulC(k.Kk, V-k.Yk))
}

func (k *Kalman) StepSix() {
	k.Pk = MatMul(MatSub(MatI(3), MatMul(k.Kk, k.GetHk())), k.Pk)
}

func (k *Kalman) Update(I, V float64) {
	k.Battery.Update(I)

	Ni := k.Battery.Ni
	if I > 0 {
		Ni = 1
	}
	k.Bk = [][]float64{
		{-Ni * k.Battery.Dt / (3600 * k.Battery.Cn)},
		{1 - math.Exp(-k.Battery.Dt/(k.Battery.R1*k.Battery.C1))},
		{1 - math.Exp(-k.Battery.Dt/(k.Battery.R2*k.Battery.C2))},
	}
	k.StepOne()
	k.StepTwo()
	k.StepThree()
	k.StepFour()
	k.StepFive(V)
	k.StepSix()
}
