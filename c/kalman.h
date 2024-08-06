#ifndef KALMAN_H
#define KALMAN_H

#include <stddef.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <math.h>

typedef struct {
	float R0;
	float R1;
	float R2;
	float C1;
	float C2;
	float Dt; 
	float Ni; 
	float Cn;
	float Zk;
} Battery;

typedef struct {
	Battery* b;
	float **Pk;
	float **Xk;
	float **Kk;
	float **Fk;
	float **Bk;
	float **Hk;
	float **SigmaWk;
	float SigmaVk;
	float Yk;
} Kalman;

extern float SOC_OCV_Curve_Readings[201];

typedef float SoC_Reading;

// Kalman Functions
Kalman* InitKalman(Battery* b);
SoC_Reading KalmanCycle(Kalman* k, float measuredCurrent, float measuredVoltage);

// Matrix Functions
float** MatT(float** A, int rowA, int colA);
float** MatMul(float** A, float** B, int rowA, int colA, int rowB, int colB);
float** MatDiv(float** A, float** B, int rowA, int colA, int rowB, int colB);
float** MatAdd(float** A, float** B, int rowA, int colA);
float** MatSub(float** A, float** B, int rowA, int colA);
float** MatMulC(float** A, int rowA, int colA, float C);
float** MatI(int Dimension);
float MatDiffuse(float** A);
void freeMatrix(float** matrix, int rows);
void printMatrix(float** matrix, int rows, int cols);

// Curve Related Functions
float Get_Voltage(float SOC);
float Get_Derivative(float SOC);

#endif	// KALMAN_H
