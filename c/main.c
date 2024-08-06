#include "kalman.h"

int main() {
	
	float A[2][3] = {{1, 2, 3}, {4, 5, 6}}, B[3][2] = {{10, 11},{20, 21},{30, 31}};
	
	float **C = MatMul(A, 2, 3, B, 3, 2);

	return 0;
}
