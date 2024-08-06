#include "kalman.h"

float SOC_OCV_Curve_Readings[201] = {
3.0,
3.06,
3.12,
3.18,
3.24,
3.3,
3.3600000000000003,
3.42,
3.48,
3.54,
3.6,
3.62,
3.64,
3.66,
3.68,
3.7,
3.7199999999999998,
3.7399999999999998,
3.76,
3.78,
3.8,
3.8024999999999998,
3.8049999999999997,
3.8074999999999997,
3.81,
3.8125,
3.815,
3.8175,
3.82,
3.8225,
3.825,
3.8275,
3.83,
3.8325,
3.835,
3.8375,
3.84,
3.8425000000000002,
3.845,
3.8475,
3.85,
3.8505000000000003,
3.851,
3.8515,
3.8520000000000003,
3.8525,
3.853,
3.8535,
3.854,
3.8545000000000003,
3.855,
3.8555,
3.856,
3.8565,
3.857,
3.8575,
3.858,
3.8585000000000003,
3.859,
3.8595,
3.8600000000000003,
3.8605,
3.861,
3.8615,
3.862,
3.8625000000000003,
3.863,
3.8635,
3.864,
3.8645,
3.865,
3.8655,
3.866,
3.8665000000000003,
3.867,
3.8675,
3.8680000000000003,
3.8685,
3.869,
3.8695,
3.87,
3.87075,
3.8715,
3.87225,
3.873,
3.8737500000000002,
3.8745000000000003,
3.87525,
3.876,
3.87675,
3.8775,
3.87825,
3.879,
3.87975,
3.8805,
3.88125,
3.882,
3.88275,
3.8835,
3.88425,
3.885,
3.88575,
3.8865,
3.88725,
3.888,
3.88875,
3.8895,
3.89025,
3.891,
3.89175,
3.8925,
3.89325,
3.894,
3.89475,
3.8954999999999997,
3.8962499999999998,
3.897,
3.89775,
3.8985,
3.89925,
3.9,
3.9025,
3.905,
3.9074999999999998,
3.91,
3.9125,
3.915,
3.9175,
3.92,
3.9225,
3.925,
3.9274999999999998,
3.9299999999999997,
3.9325,
3.935,
3.9375,
3.94,
3.9425,
3.945,
3.9475,
3.95,
3.9525,
3.955,
3.9575,
3.96,
3.9625,
3.965,
3.9675,
3.9699999999999998,
3.9725,
3.975,
3.9775,
3.98,
3.9825,
3.985,
3.9875,
3.99,
3.9925,
3.995,
3.9975,
4.0,
4.005,
4.01,
4.015,
4.02,
4.025,
4.03,
4.035,
4.04,
4.045,
4.05,
4.055,
4.0600000000000005,
4.065,
4.07,
4.075,
4.08,
4.085,
4.09,
4.095,
4.1,
4.105,
4.11,
4.115,
4.12,
4.125,
4.13,
4.135,
4.140000000000001,
4.1450000000000005,
4.15,
4.155,
4.16,
4.165,
4.17,
4.175,
4.18,
4.1850000000000005,
4.19,
4.195,
4.2
};

float Get_Voltage(float SOC) {
	float low = floorf(SOC), mid = low + 0.5, high = ceilf(SOC);

	if (high - SOC < SOC - low) {
		return (SOC_OCV_Curve_Readings[(int)(high * 2 + 1)] + SOC_OCV_Curve_Readings[(int)(mid * 2 + 1)]) / 2.0;
	} else {
		return (SOC_OCV_Curve_Readings[(int)(low * 2 + 1)] + SOC_OCV_Curve_Readings[(int)(mid * 2 + 1)]) / 2.0;
	}
}
