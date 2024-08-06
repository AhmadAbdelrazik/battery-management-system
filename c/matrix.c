#include "kalman.h"


// MatMul: Multiplies two matrices
double** MatMul(double** A, double** B, int rowA, int colA, int rowB, int colB) {
    if (colA != rowB) {
        return NULL;
    }

    double** result = (double**)malloc(rowA * sizeof(double*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (double*)calloc(colB, sizeof(double));
    }

    for (int row = 0; row < rowA; row++) {
        for (int col = 0; col < colB; col++) {
            for (int i = 0; i < colA; i++) {
                result[row][col] += A[row][i] * B[i][col];
            }
        }
    }

    return result;
}

// MatT: Transpose Matrix
double** MatT(double** A, int rowA, int colA) {
    double** result = (double**)malloc(colA * sizeof(double*));
    for (int i = 0; i < colA; i++) {
        result[i] = (double*)malloc(rowA * sizeof(double));
    }

    for (int row = 0; row < rowA; row++) {
        for (int col = 0; col < colA; col++) {
            result[col][row] = A[row][col];
        }
    }

    return result;
}

// MatDiv: Divides two matrices (Multiplication with transpose)
double** MatDiv(double** A, double** B, int rowA, int colA, int rowB, int colB) {
    double** Bt = MatT(B, rowB, colB);
    double** result = MatMul(A, Bt, rowA, colA, colB, rowB);
    for (int i = 0; i < colB; i++) {
        free(Bt[i]);
    }
    free(Bt);
    return result;
}

// MatAdd: Adds two matrices
double** MatAdd(double** A, double** B, int rowA, int colA) {
    double** result = (double**)malloc(rowA * sizeof(double*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (double*)malloc(colA * sizeof(double));
    }

    for (int i = 0; i < rowA; i++) {
        for (int j = 0; j < colA; j++) {
            result[i][j] = A[i][j] + B[i][j];
        }
    }

    return result;
}

// MatSub: Subtracts two matrices
double** MatSub(double** A, double** B, int rowA, int colA) {
    double** result = (double**)malloc(rowA * sizeof(double*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (double*)malloc(colA * sizeof(double));
    }

    for (int i = 0; i < rowA; i++) {
        for (int j = 0; j < colA; j++) {
            result[i][j] = A[i][j] - B[i][j];
        }
    }

    return result;
}

// MatMulC: Multiplies a matrix with a constant
double** MatMulC(double** A, int rowA, int colA, double C) {
    double** result = (double**)malloc(rowA * sizeof(double*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (double*)malloc(colA * sizeof(double));
    }

    for (int i = 0; i < rowA; i++) {
        for (int j = 0; j < colA; j++) {
            result[i][j] = A[i][j] * C;
        }
    }

    return result;
}

// MatI: Produces Identity Matrix
double** MatI(int Dimension) {
    double** I = (double**)malloc(Dimension * sizeof(double*));
    for (int i = 0; i < Dimension; i++) {
        I[i] = (double*)calloc(Dimension, sizeof(double));
    }

    for (int i = 0; i < Dimension; i++) {
        I[i][i] = 1;
    }

    return I;
}

// MatDiffuse: Converts a 1 * 1 matrix to a float
double MatDiffuse(double** A) {
    return A[0][0];
}


// Helper function to free matrix memory
void freeMatrix(double** matrix, int rows) {
    for (int i = 0; i < rows; i++) {
        free(matrix[i]);
    }
    free(matrix);
}

// Helper function to print matrix
void printMatrix(double** matrix, int rows, int cols) {
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%f ", matrix[i][j]);
        }
        printf("\n");
    }
}


// Example usage
/*int main() {*/
/*    double* Adata[] = { (double[]){1, 2}, (double[]){3, 4} };*/
/*    double** A = (double**)Adata;*/
/**/
/*    double* Bdata[] = { (double[]){5, 6}, (double[]){7, 8} };*/
/*    double** B = (double**)Bdata;*/
/**/
/*    int rowA = 2, colA = 2, rowB = 2, colB = 2;*/
/**/
/*    double** result = MatMul(A, B, rowA, colA, rowB, colB);*/
/*    printMatrix(result, rowA, colB);*/
/*    freeMatrix(result, rowA);*/
/**/
/*    return 0;*/
/*}*/
/**/
