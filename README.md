# linalgo

`linalgo` is a lightweight Go library for linear algebra operations on vectors and matrices. It provides common functionality such as vector operations, matrix multiplication, determinant computation, rank, identity checks, and more.

## Features

Here are some key features, provided by the library:

- Vectors:
  - Create, add, scale, normalize, project, etc.
  - Compute norm, dot product, etc.

- Matrices:
  - Create from 2D or flat data
  - Multiply, invert, compute determinant, rank, etc.
  - Check for identity, zero, or other special matrix types

---

## Preview 

To get a quick preview, please check out:

- `examples/matrix_demo/main.go`
- `examples/vector_demo/main.go`

--- 

## Tests

### Large unit test coverage:

- matrix: 94.1%
- vector: 100%

### Covers edge cases such as:

- Empty or malformed inputs  
- Zero matrices or vectors  
- Non-square matrices  
- Singular and non-invertible matrices

All tests pass with:

```bash
go test -v ./...
```

Demos in examples/ are excluded from testing but useful for interactive runs

---

## Prerequisites

To build and test this project, make sure you have:

- Go 1.18+ installed  
- Make  
- (Optional) CMake if you plan on integrating with other native modules (not required by default)

---

## Installation

```bash
go get github.com/JoLandry/linalgo
```

Then import and use : 

import (
    "github.com/JoLandry/linalgo/vector"
    "github.com/JoLandry/linalgo/matrix"
)

---

## Build and MakeFile commands

You can use the provided Makefile for common development workflows:

```bash
make                  # Runs fmt, vet, and test
make test             # Run unit tests
make test-cover       # Run tests with coverage report
make run-matrix-demo  # Run example matrix demo
make run-vector-demo  # Run example vector demo
make clean            # Clean up coverage files
```

---

## Project Structure

```text
linalgo/
│
├── matrix/         # Matrix types and operations
├── vector/         # Vector types and operations
├── examples/       # Demos for Matrix and Vector types 
├── Makefile        # Build and test automation
└── go.mod          # Module definition

---

## License

This project was created as a personal initiative, outside any official academic coursework, during my Master's in Computer Science at the University of Bordeaux.

It is released under the MIT License.

You are free to:
    use, copy, modify, and distribute this code,
    as long as you retain the copyright notice.

This project is provided "as is", without any warranty of any kind.

© 2025 Landry Jonathan