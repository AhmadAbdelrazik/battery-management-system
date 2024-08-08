#include "kalman.h"

int main() {
	float R0, R1, R2, C1, C2, Dt, Ni, Cn, Zk;

	Battery B = {
	R0 = 2.1552 * pow(10, -3),
	R1 = 6.7587 * pow(10, -3),
	R2 = 2.5683 * pow(10, -3),
	C1 = 123.1912 * pow(10,3),
	C2 = 4.5240 * pow(10, 3),
	Dt = 0.1,    // Time step in seconds
	Ni = 0.9894, // Coulombic Efficiency
	Cn = 6.8080, // Nominal Capacity
	Zk = 1.0    // State of Charge, 1.0 == 100%
};

	Kalman *k = (Kalman *)InitKalman(&B);

	int i;
	float SoC, voltage;
	for (i = 0; i < 10; i++) {
		SoC = KalmanMockCycle(k, 0.4, &voltage);
		printf("SOC = %.3f\tVoltage = %.3f\n", SoC, voltage);
	}

	return 0;
}
