#ifndef KALMAN_H
#define KALMAN_H

#include <stddef.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

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
	float Pk[3][3];
	float Xk[3][1];
	float Kk[3][1];
	float Fk[3][3];
	float Bk[3][1];
	float SigmaWk[1][1];
	float SigmaVk;
	float Yk;
} Kalman;

typedef float SoC_Reading;

// Matrix Functions
double** MatT(double** A, int rowA, int colA);
double** MatMul(double** A, double** B, int rowA, int colA, int rowB, int colB);
double** MatDiv(double** A, double** B, int rowA, int colA, int rowB, int colB);
double** MatAdd(double** A, double** B, int rowA, int colA);
double** MatSub(double** A, double** B, int rowA, int colA);
double** MatMulC(double** A, int rowA, int colA, double C);
double** MatI(int Dimension);
double MatDiffuse(double** A);
void freeMatrix(double** matrix, int rows);
void printMatrix(double** matrix, int rows, int cols);


Kalman* InitKalman(Battery* b);
SoC_Reading KalmanCycle(Kalman* k, float measuredCurrent, float measuredVoltage);


#endif	// KALMAN_H
