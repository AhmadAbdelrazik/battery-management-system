#ifndef KALMAN_H
#define KALMAN_H

typedef struct {
	float R0;
	float R1;
	float R2;
	float C1;
	float C2;
	float I1;
	float I2;
	float I;
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


Kalman* InitKalman(Battery* b);



#endif	// KALMAN_H
