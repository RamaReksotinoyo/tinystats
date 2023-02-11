# Tiny Statistics Library in Go

Im on way to learning go, and i want to make a simple statistics library in go. I hope this library can help you to learn go too.

This library provides a set of functions for performing basic statistical computations, such as calculating the mean, median, mode. (soon) standard deviation, etc. The library is written in Go and is designed to be simple and easy to use. 

## Features

- Calculate the mean, median, and mode of a set of numbers
- Calculate the data range of a set of numbers
- Calculate the quantiles of a set of numbers
- Calculate the de mean

## Installation

You can install the library by using `go get`:

go get github.com/yourusername/tinystatistics


## Usage

Here's an example of how to use the library to calculate the mean of a set of numbers:

```go
package main

import (
	"fmt"
	"github.com/RamaReksotinoyo/tinystatistics"
)

func main() {
	numbers := []float64{1, 2, 3, 4, 5}
	mean := tinystatistics.Mean(numbers)
	fmt.Println("Mean:", mean)
}
```
This will output:

Mean: 3

go get github.com/RamaReksotinoyo/tinystatistics


# API Reference
The library provides the following functions:

- Mean(numbers []float64) float64: calculates the mean of a set of numbers
- Median(numbers []float64) float64: calculates the median of a set of numbers
- Mode(numbers []float64) float64: calculates the mode of a set of numbers
- De Mean(numbers []float64) float64: calculates the de mean of a set of numbers

# Soon
I will build more advanced of statistical computation, such as standard deviation, variance, correlation, covariance, etc.


# Contributing
Contributions are welcome! If you'd like to contribute to the development of this library, please fork the repository and submit a pull request.

# License
The library is released under the MIT license. See the LICENSE file for more information.