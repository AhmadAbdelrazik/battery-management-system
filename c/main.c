#include "kalman.h"

int main() {


	double* Adata[] = { (double[]){1, 2}, (double[]){3, 4} };
	double** A = (double**)Adata;

	double* Bdata[] = { (double[]){5, 6}, (double[]){7, 8} };
	double** B = (double**)Bdata;

	int rowA = 2, colA = 2, rowB = 2, colB = 2;

	double** result = MatMul(A, B, rowA, colA, rowB, colB);
	printMatrix(result, rowA, colB);
	freeMatrix(result, rowA);

	return 0;
}
