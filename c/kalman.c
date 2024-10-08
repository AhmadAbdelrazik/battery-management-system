#include "kalman.h"
#include <math.h>


// Helper Functions
float* allocate_1(float element);
float* allocate_3(float e1, float e2, float e3);
float **allocate_3_ptrs(float *ptr1, float *ptr2, float *ptr3);
float **allocate_1_ptr(float *ptr);

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
	float OCV = Get_Voltage(k->Xk[0][0] * 100);
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


Kalman* InitKalman(Battery* b) {
	Kalman *k = NULL;
	k = (Kalman *)malloc(sizeof(Kalman));

	k->b = b;

	// Initial State
	k->Xk = allocate_3_ptrs(
		allocate_1(b->Zk),
		allocate_1(0),
		allocate_1(0)
	);

	// Initial Error Covariance
	k->Pk = allocate_3_ptrs(
		allocate_3(0.005, 0, 0),
		allocate_3(0, 0.0001, 0),
		allocate_3(0, 0, 0.001)
	);

	// Process Jacobian
	k->Fk = allocate_3_ptrs(
		allocate_3(1, 0, 0),
		allocate_3(0,expf(-k->b->Dt / (k->b->R1 * k->b->C1)) ,0),
		allocate_3(0,0,expf(-k->b->Dt / (k->b->R2 * k->b->C2)))
	);


	k->Bk = allocate_3_ptrs(
		allocate_1(-k->b->Ni * k->b->Dt / (3600 * k->b->Cn)),
		allocate_1(1 - expf(-k->b->Dt / (k->b->R1 * k->b->C1))),
		allocate_1(1 - expf(-k->b->Dt / (k->b->R2 * k->b->C2)))
	);


	float derivative = Get_Derivative(k->Xk[0][0] * 100);

	k->Hk = allocate_1_ptr(allocate_3(derivative, -k->b->R1, -k->b->R2));

	float **Bkt = MatT(k->Bk, 3, 1);
	k->SigmaWk = MatMul(k->Bk, Bkt, 3, 1, 1, 3);
	k->SigmaVk = pow(0.035, 2);

	freeMatrix(Bkt, 1);
	return k;
}



SoC_Reading KalmanCycle(Kalman* k, float measuredCurrent, float measuredVoltage) {
	float Ni = k->b->Ni;
	if (measuredCurrent > 0) {
		Ni = 1;
	}
	
	k->Bk[0][0] = - Ni * k->b->Dt / (3600 * k->b->Cn);

	StepOne(k, measuredCurrent);
	StepTwo(k);
	StepThree(k, measuredCurrent);

	k->Hk[0][0] = Get_Derivative(k->Xk[0][0] * 100);

	StepFour(k);
	StepFive(k, measuredVoltage);
	StepSix(k);

	return k->Xk[0][0];
}


SoC_Reading KalmanMockCycle(Kalman* k, float measuredCurrent, float *voltage) {
	float Ni = k->b->Ni;
	if (measuredCurrent > 0) {
		Ni = 1;
	}
	
	k->Bk[0][0] = - Ni * k->b->Dt / (3600 * k->b->Cn);

	StepOne(k, measuredCurrent);
	StepTwo(k);
	StepThree(k, measuredCurrent);

	k->Hk[0][0] = Get_Derivative(k->Xk[0][0] * 100);

	StepFour(k);
	StepFive(k, k->Yk);
	StepSix(k);

	*voltage = k->Yk;
	return k->Xk[0][0];
}


float* allocate_3(float e1, float e2, float e3) {
	float *arr = NULL;

	arr = (float *)malloc(sizeof(float) * 3);
	if (arr == NULL) {
		return arr;
	}

	arr[0] = e1;
	arr[1] = e2;
	arr[2] = e3;

	return arr;
}

float* allocate_1(float element) {
	float *arr = NULL;

	arr = (float *)malloc(sizeof(float));
	if (arr == NULL) {
		return arr;
	}

	arr[0] = element;

	return arr;
}

float **allocate_3_ptrs(float *ptr1, float *ptr2, float *ptr3) {
	float **arr = NULL;

	arr = (float **)malloc(sizeof(float *) * 3);
	if (arr == NULL) {
		return arr;
	}
	
	arr[0] = ptr1;
	arr[1] = ptr2;
	arr[2] = ptr3;

	return arr;
}

float **allocate_1_ptr(float *ptr) {
	float **arr = NULL;

	arr = (float **)malloc(sizeof(float *));
	if (arr == NULL) {
		return arr;
	}
	
	arr[0] = ptr;

	return arr;
}
