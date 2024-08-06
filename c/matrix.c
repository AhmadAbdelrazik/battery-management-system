#include "kalman.h"


// MatMul: Multiplies two matrices
float** MatMul(float** A, float** B, int rowA, int colA, int rowB, int colB) {
    if (colA != rowB) {
        return NULL;
    }

    float** result = (float**)malloc(rowA * sizeof(float*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (float*)calloc(colB, sizeof(float));
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
float** MatT(float** A, int rowA, int colA) {
    float** result = (float**)malloc(colA * sizeof(float*));
    for (int i = 0; i < colA; i++) {
        result[i] = (float*)malloc(rowA * sizeof(float));
    }

    for (int row = 0; row < rowA; row++) {
        for (int col = 0; col < colA; col++) {
            result[col][row] = A[row][col];
        }
    }

    return result;
}

// MatDiv: Divides two matrices (Multiplication with transpose)
float** MatDiv(float** A, float** B, int rowA, int colA, int rowB, int colB) {
    float** Bt = MatT(B, rowB, colB);
    float** result = MatMul(A, Bt, rowA, colA, colB, rowB);
    for (int i = 0; i < colB; i++) {
        free(Bt[i]);
    }
    free(Bt);
    return result;
}

// MatAdd: Adds two matrices
float** MatAdd(float** A, float** B, int rowA, int colA) {
    float** result = (float**)malloc(rowA * sizeof(float*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (float*)malloc(colA * sizeof(float));
    }

    for (int i = 0; i < rowA; i++) {
        for (int j = 0; j < colA; j++) {
            result[i][j] = A[i][j] + B[i][j];
        }
    }

    return result;
}

// MatSub: Subtracts two matrices
float** MatSub(float** A, float** B, int rowA, int colA) {
    float** result = (float**)malloc(rowA * sizeof(float*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (float*)malloc(colA * sizeof(float));
    }

    for (int i = 0; i < rowA; i++) {
        for (int j = 0; j < colA; j++) {
            result[i][j] = A[i][j] - B[i][j];
        }
    }

    return result;
}

// MatMulC: Multiplies a matrix with a constant
float** MatMulC(float** A, int rowA, int colA, float C) {
    float** result = (float**)malloc(rowA * sizeof(float*));
    for (int i = 0; i < rowA; i++) {
        result[i] = (float*)malloc(colA * sizeof(float));
    }

    for (int i = 0; i < rowA; i++) {
        for (int j = 0; j < colA; j++) {
            result[i][j] = A[i][j] * C;
        }
    }

    return result;
}

// MatI: Produces Identity Matrix
float** MatI(int Dimension) {
    float** I = (float**)malloc(Dimension * sizeof(float*));
    for (int i = 0; i < Dimension; i++) {
        I[i] = (float*)calloc(Dimension, sizeof(float));
    }

    for (int i = 0; i < Dimension; i++) {
        I[i][i] = 1;
    }

    return I;
}

// MatDiffuse: Converts a 1 * 1 matrix to a float
float MatDiffuse(float** A) {
    return A[0][0];
}


// Helper function to free matrix memory
void freeMatrix(float** matrix, int rows) {
    for (int i = 0; i < rows; i++) {
        free(matrix[i]);
    }
    free(matrix);
}

// Helper function to print matrix
void printMatrix(float** matrix, int rows, int cols) {
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%f ", matrix[i][j]);
        }
        printf("\n");
    }
}


// Example usage
int main() {
    float* Adata[] = { (float[]){1, 2}, (float[]){3, 4} };
    float** A = (float**)Adata;

    float* Bdata[] = { (float[]){5, 6}, (float[]){7, 8} };
    float** B = (float**)Bdata;

    int rowA = 2, colA = 2, rowB = 2, colB = 2;

    float** result = MatMul(A, B, rowA, colA, rowB, colB);
    printMatrix(result, rowA, colB);
    freeMatrix(result, rowA);

    return 0;
}

