package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <data-file>\n", os.Args[0])
		os.Exit(1)
	}
	filePath := os.Args[1]

	values, err := readValues(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading values: %v\n", err)
		os.Exit(1)
	}

	mean := computeMean(values)
	median := computeMedian(values)
	variance := computeVariance(values, mean)
	stddev := math.Sqrt(variance)

	// Print rounded integer results
	fmt.Printf("Average: %d\n", int64(math.Round(mean)))
	fmt.Printf("Median: %d\n", int64(math.Round(median)))
	fmt.Printf("Variance: %d\n", int64(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int64(math.Round(stddev)))
}

func readValues(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []float64
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			// skip non-numeric lines with a warning
			fmt.Fprintf(os.Stderr, "Warning: skipping non-numeric line: '%s'\n", line)
			continue
		}
		values = append(values, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, fmt.Errorf("no numeric data found in the file")
	}
	return values, nil
}

func computeMean(vals []float64) float64 {
	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	return sum / float64(len(vals))
}

func computeMedian(vals []float64) float64 {
	sorted := make([]float64, len(vals))
	copy(sorted, vals)
	sort.Float64s(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return sorted[n/2]
	}
	mid1 := sorted[(n/2)-1]
	mid2 := sorted[n/2]
	return (mid1 + mid2) / 2
}

func computeVariance(vals []float64, mean float64) float64 {
	sumSq := 0.0
	for _, v := range vals {
		d := v - mean
		sumSq += d * d
	}
	return sumSq / float64(len(vals))
}
