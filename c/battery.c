#include <stdlib.h>

typedef struct {
	double R0;
	double R1;
	double R2;
	double C1;
	double C2;
	double I1;
	double I2;
	double I;
	double Dt; 
	double Ni; 
	double Cn;
	double Zk;
} Battery;

Battery* NewBattery(
	double R0,
	double R1,
	double R2,
	double C1,
	double C2,
	double I1,
	double I2,
	double I,
	double Dt, 
	double Ni, 
	double Cn,
	double Zk
) {
	Battery *b = NULL;

	b = malloc(sizeof(Battery));


	b->R0 = R0;
	b->R1 = R1;
	b->R2 = R2;
	b->C1 = C1;
	b->C2 = C2;
	b->I1 = I1;
	b->I2 = I2;
	b->I = I;
	b->Dt = Dt;
	b->Ni = Ni;
	b->Cn = Cn;
	b->Zk = Zk;

	return b;
};
