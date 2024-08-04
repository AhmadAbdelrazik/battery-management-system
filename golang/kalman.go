package main

import (
	"math"
)

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

func NewKalman(b *Battery, d SoCOCV) Kalman {
	var k Kalman

	k.Battery = b

	k.Data = d
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

	k.SigmaWk = MatMul(k.Bk, MatT(k.Bk))
	k.SigmaVk = math.Pow(0.035, 2)

	return k
}

// Hk
func (k *Kalman) GetHk() [][]float64 {
	derivative := k.Data.derivative(k.Xk[0][0] * 100)
	Hk := [][]float64{
		{derivative, -k.Battery.R1, -k.Battery.R2},
	}

	return Hk
}

func (k *Kalman) StepOne() {
	val := MatMulC(k.Bk, k.Battery.I)
	k.Xk = MatAdd(MatMul(k.Fk, k.Xk), val)
}

func (k *Kalman) StepTwo() {
	k.Pk = MatAdd(MatMul(MatMul(k.Fk, k.Pk), MatT(k.Fk)), k.SigmaWk)
}

// Predict System Output
func (k *Kalman) StepThree() {
	OCV := k.Data.GetVoltage(k.Xk[0][0] * 100)
	k.Yk = OCV - k.Battery.R0*k.Battery.I - k.Battery.R1*k.Battery.I1 - k.Battery.R2*k.Battery.I2
}

func (k *Kalman) StepFour() {
	hk := k.GetHk()

	HKxPKxHKt := MatMul(MatMul(hk, k.Pk), MatT(hk))
	SigmaY := MatDiffuse(HKxPKxHKt) + k.SigmaVk

	k.Kk = MatMulC(MatMul(k.Pk, MatT(hk)), 1/SigmaY)
}

func (k *Kalman) StepFive(V float64) {
	k.Xk = MatAdd(k.Xk, MatMulC(k.Kk, k.Yk-V))
}

func (k *Kalman) StepSix() {
	k.Pk = MatMul(MatSub(MatI(3), MatMul(k.Kk, k.GetHk())), k.Pk)
}

func (k *Kalman) Update(I, V float64) float64 {
	k.Battery.Update(I)

	// Coulombic Efficiency
	Ni := k.Battery.Ni
	if I > 0 {
		Ni = 1
	}

	k.Bk[0][0] = -Ni * k.Battery.Dt / (3600 * k.Battery.Cn)

	// prediction
	k.StepOne()
	k.StepTwo()
	k.StepThree()

	// update
	k.StepFour()
	k.StepFive(V)
	k.StepSix()

	return k.Xk[0][0]
}

func (k *Kalman) UpdateTest(I float64) (float64, float64) {
	k.Battery.Update(I)

	// Coulombic Efficiency
	Ni := k.Battery.Ni
	if I > 0 {
		Ni = 1
	}

	k.Bk[0][0] = -Ni * k.Battery.Dt / (3600 * k.Battery.Cn)

	// prediction
	k.StepOne()
	k.StepTwo()
	k.StepThree()

	// update
	mockV := min(generateTestValues(k.Yk, 0.035))
	k.StepFour()
	k.StepFive(mockV)
	k.StepSix()

	return k.Xk[0][0], k.Yk
}
