#include "kalman.h"

Kalman* InitKalman(Battery* b) {

	return NULL;
}

SoC_Reading KalmanCycle(Kalman* k, float measuredCurrent, float measuredVoltage) {

	return 0;
}


void StepOne(Kalman *k, float measuredCurrent) {
	float **BkxI = MatMulC(k->Bk, 3, 1, measuredCurrent);
	float **FkxXk = MatMul(k->Fk, k->Xk, 3, 3, 3, 1);
	freeMatrix(k->Xk, 3);
	k->Xk = MatAdd(FkxXk, BkxI, 3, 1);

	freeMatrix(BkxI, 3);
	freeMatrix(FkxXk, 3);
}

void StepTwo(Kalman *k) {
	float **FkxPk = MatMul(k->Fk, k->Pk, 3, 3, 3, 3);
	float **Fkt = MatT(k->Fk, 3, 3);
	float **FkxPkxFkt = MatMul(FkxPk, Fkt, 3, 3, 3, 3);
	freeMatrix(k->Pk, 3);
	k->Pk = MatAdd(FkxPkxFkt, k->SigmaWk, 3, 3);
	freeMatrix(FkxPkxFkt, 3);
	freeMatrix(Fkt, 3);
	freeMatrix(FkxPk, 3);
}

void StepThree(Kalman *k, float measuredCurrent) {
	// get the closest reading to the SoC.
	float OCV = Get_Voltage(k->Xk[0][0]);
	k->Yk = OCV - k->b->R0 * measuredCurrent + k->Hk[0][1] * k->Xk[1][0] + k->Hk[0][2] * k->Xk[2][0];
}

void StepFour(Kalman *k) {
	float **HkxPk = MatMul(k->Hk, k->Pk, 1, 3, 3, 3);
	float **Hkt = MatT(k->Hk, 1, 3);
	float **HkxPkxHkt = MatMul(HkxPk, Hkt, 1, 3, 3, 1);
	float SigmaY = HkxPkxHkt[0][0] + k->SigmaVk;

	float **PkxHkt = MatMul(k->Pk, Hkt, 3, 3, 3, 1);
	freeMatrix(k->Kk, 3);
	k->Kk = MatMulC(PkxHkt, 3, 1, SigmaY);

	freeMatrix(HkxPk, 1);
	freeMatrix(Hkt, 3);
	freeMatrix(HkxPkxHkt, 1);
	freeMatrix(PkxHkt, 3);
}

void StepFive(Kalman *k, float measuredVoltage) {
	float **KkxVolt = MatMulC(k->Kk, 3, 1, k->Yk - measuredVoltage);
	float **oldXk = k->Xk;
	k->Xk = MatAdd(k->Xk, KkxVolt, 3, 1);
	freeMatrix(KkxVolt, 3);
	freeMatrix(oldXk, 3);
}

void StepSix(Kalman *k) {
	float **I = MatI(3);
	float **KkxHk = MatMul(k->Kk, k->Hk, 3, 1, 1, 3);
	float **I_KkxHk = MatSub(I, KkxHk, 3, 3);
	float **oldPk = k->Pk;
	k->Pk = MatMul(I_KkxHk, k->Pk, 3, 3, 3, 3);

	freeMatrix(I, 3);
	freeMatrix(KkxHk, 3);
	freeMatrix(I_KkxHk, 3);
	freeMatrix(oldPk, 3);
}
