# Battery Modeling and Kalman Filter-based State-of-Charge Estimation for a Race Car Application

## Overview

This project involves implementing a method to estimate the State-of-Charge (SOC) of a lithium-ion polymer battery used in a Formula Student electric race car. The method utilizes a battery equivalent circuit model and an Extended Kalman Filter (EKF) to provide accurate SOC estimations. The project will be implemented in two programming languages: Go and C.

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Running the Go Implementation](#running-the-go-implementation)
  - [Running the C Implementation](#running-the-c-implementation)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Revolve NTNU, a student organization, designs and builds electric race cars for the Formula Student competition. The battery pack, consisting of lithium-polymer cells, is subjected to harsh conditions, requiring precise monitoring to avoid over-charging, over-discharging, and overheating. Estimating the SOC is crucial for efficient battery management and optimal performance during races.

The Extended Kalman Filter (EKF) is used to dynamically estimate the SOC based on a battery model. This project aims to implement the battery modeling and SOC estimation using Go and C languages.

## Project Structure

```
battery-soc-estimation/
├── go/
│   ├── main.go
│   ├── battery.go
│   ├── kalman.go
│   ├── matrix.go
│   ├── matrix_test.go
│   └── README.md
└── c/
    ├── main.c
    ├── battery_model.c
    ├── ekf.c
    ├── battery_model.h
    ├── ekf.h
    └── README.md
```

## Getting Started

### Prerequisites

- Go 1.22 or higher
- GCC compiler for C

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/battery-soc-estimation.git
   cd battery-soc-estimation
   ```

## Usage

### Running the Go Implementation

1. Navigate to the Go directory:
   ```bash
   cd go
   ```

2. Run the Go program:
   ```bash
   go run main.go
   ```

### Running the C Implementation

1. Navigate to the C directory:
   ```bash
   cd c
   ```

2. Compile the C program:
   ```bash
   gcc main.c battery_model.c ekf.c -o battery_soc_estimation
   ```

3. Run the C program:
   ```bash
   ./battery_soc_estimation
   ```

## Testing

For both implementations, ensure you have test data and expected results to validate the SOC estimation accuracy. Unit tests should be written to cover various aspects of the battery model and EKF functionality.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
