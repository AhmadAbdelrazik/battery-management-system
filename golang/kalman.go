package main

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

type Kalman struct {
	Battery *Battery
	SoC_OCV SoCOCV
	Pk      [][]float64 // Error Covariance
	Xk      [][]float64 // State Estimate
	Kk      [][]float64 // Kalman Gain
	Fk      [][]float64 // Process Jacobian
	Bk      [][]float64 // Control Input Matrix
	SigmaWk [][]float64 // Process Noise Covariance Matrix
	SigmaVk float64     // Variance of the voltage sensor
	Yk      float64     // Estimated Terminal Voltage
}

// Initialize a new Kalman Filter for a specific battery and
// SOC-OCV Curve data.
func NewKalman(b *Battery, d SoCOCV) Kalman {
	var k Kalman

	k.Battery = b
	k.SoC_OCV = d

	// Initial State
	k.Xk = MatT([][]float64{{b.Zk, 0, 0}})
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

	// Control Input Matrix
	k.Bk = [][]float64{
		{-k.Battery.Ni * k.Battery.Dt / (3600 * k.Battery.Cn)},
		{1 - math.Exp(-k.Battery.Dt/(k.Battery.R1*k.Battery.C1))},
		{1 - math.Exp(-k.Battery.Dt/(k.Battery.R2*k.Battery.C2))},
	}

	k.SigmaWk = MatMul(k.Bk, MatT(k.Bk))
	k.SigmaVk = math.Pow(0.035, 2)

	return k
}

// Hk is the Measurement Jacobian
func (k *Kalman) GetHk() [][]float64 {
	// the rate of change of the Open Circuit Voltage to the State of Charge
	derivative := k.SoC_OCV.Derivative(k.Xk[0][0] * 100)

	Hk := [][]float64{
		{derivative, -k.Battery.R1, -k.Battery.R2},
	}

	return Hk
}

// StepOne: State Estimate Propagation
// In this step we propagate the previous state to the current state
// using the input current to predict the current state.
// X_{k+1} = F_k X_k + B_k i_k + W_k
func (k *Kalman) StepOne() {
	val := MatMulC(k.Bk, k.Battery.I)
	k.Xk = MatAdd(MatMul(k.Fk, k.Xk), val)
}

// StepTwo: Error Covariance Propagation
// In this step we propagate the error covariance of the system.
// Error covariance is used to count for errors in the measurement
// and in prediction. The Error Covariance uses a gaussian distribution
// for error margains rather than the completely random distribution
func (k *Kalman) StepTwo() {
	k.Pk = MatAdd(MatMul(MatMul(k.Fk, k.Pk), MatT(k.Fk)), k.SigmaWk)
}

// StepThree: Predict System Output
// The system uses 2nd-order Thevenin equivalent circuit to simulate the
// Car Battery, we calculate the terminal voltage using the Battery parameters
// and get the OCV using the SOC-OCV graph using the current SoC.
func (k *Kalman) StepThree() {
	OCV := k.SoC_OCV.GetVoltage(k.Xk[0][0] * 100)
	k.Yk = OCV - k.Battery.R0*k.Battery.I - k.Battery.R1*k.Battery.I1 - k.Battery.R2*k.Battery.I2
}

// StepFour: Calculate the kalman gain
// Kalman gain is used as an indication of the stability of the system.
// In the first iterations, the measured values of voltage are considered
// to be more accurate and though has more weight than the predicted output.
// as the system matures, the measured voltage weight becomes lesser as the
// system becomes more stable and based on many measurements.
func (k *Kalman) StepFour() {
	hk := k.GetHk()

	HKxPKxHKt := MatMul(MatMul(hk, k.Pk), MatT(hk))
	SigmaY := MatDiffuse(HKxPKxHKt) + k.SigmaVk

	k.Kk = MatMulC(MatMul(k.Pk, MatT(hk)), 1/SigmaY)
}

// StepFive: State Esitmate Update
// We update the system state using the measured and predicted terminal
// voltages. The severity of the update is based on the kalman gain
// calculated in the previous step.
func (k *Kalman) StepFive(measuredVoltage float64) {
	k.Xk = MatAdd(k.Xk, MatMulC(k.Kk, k.Yk-measuredVoltage))
}

// StepSix: Error Covariance Update
// We update the error covariance to reflect the maturity of the system
// as more measured values are added to the system.
func (k *Kalman) StepSix() {
	k.Pk = MatMul(MatSub(MatI(3), MatMul(k.Kk, k.GetHk())), k.Pk)
}

// Cycle: Represents a full Kalman Filter Cycle
// measuredVoltage: the terminal voltage of the battery.
// measuredCurrent: The current drawn from the battery to feed the system.
func (k *Kalman) Cycle(measuredCurrent, measuredVoltage float64) float64 {
	k.Battery.UpdateCurrents(measuredCurrent)

	// Coulombic Efficiency
	// 1  at discharge
	// Ni at charge
	Ni := k.Battery.Ni
	if measuredCurrent > 0 {
		Ni = 1
	}

	k.Bk[0][0] = -Ni * k.Battery.Dt / (3600 * k.Battery.Cn)

	// prediction
	k.StepOne()
	k.StepTwo()
	k.StepThree()

	// update
	k.StepFour()
	k.StepFive(measuredVoltage)
	k.StepSix()

	return k.Xk[0][0]
}

// MockCycle: full Kalman Filter Cycle, this is used for demonstration
// purposes. In this MockCycle the measuredVoltage equals to the
// predicted voltage + gaussian noise. it aims to simulate the working
// of a voltage sensor
func (k *Kalman) MockCycle(I float64) (float64, float64) {
	k.Battery.UpdateCurrents(I)

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
	mockV := addGaussianError(k.Yk, k.SigmaVk)
	k.StepFour()
	k.StepFive(mockV)
	k.StepSix()

	return k.Xk[0][0], k.Yk
}

func addGaussianError(mean, sigma float64) float64 {
	normalDist := distuv.Normal{
		Mu:    mean,
		Sigma: sigma,
	}

	return normalDist.Rand()
}
