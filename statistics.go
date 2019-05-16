package data_science

import (
	"math"
	"sort"
)

// Mean vector
func Mean(vec Vector) float64 {
	return Sum(vec) / float64(len(vec))
}

// Median
func Median(vec Vector) float64 {
	n := len(vec)
	sortedVec := make(Vector, len(vec))
	copy(sortedVec, vec)
	sort.Sort(sortedVec)
	midPoint := n / 2
	if n%2 == 1 {
		return sortedVec[midPoint]
	}
	low := midPoint - 1
	high := midPoint
	return sortedVec[low] + sortedVec[high]/2
}

// Quantile
func Quantile(vec Vector, p float64) float64 {
	pIdx := int(p * float64(len(vec)))
	sortedVec := make(Vector, len(vec))
	copy(sortedVec, vec)
	sort.Sort(sortedVec)
	return sortedVec[pIdx]
}

// Range
func Range(vec Vector) float64 {
	return Max(vec) - Min(vec)
}

// Min of vector
func Min(vec Vector) (res float64) {
	for i, e := range vec {
		if i == 0 || e < res {
			res = e
		}
	}
	return
}

// Max of vector
func Max(vec Vector) (res float64) {
	for i, e := range vec {
		if i == 0 || e > res {
			res = e
		}
	}
	return
}

// De Mean.
func DeMean(vec Vector) (res Vector) {
	res = make(Vector, len(vec))
	mean := Mean(vec)
	for i, val := range vec {
		res[i] = val - mean
	}
	return
}

// Variance method.
func Variance(vec Vector) float64 {
	n := len(vec)
	deviations := DeMean(vec)
	return SumOfSquares(deviations) / float64(n-1)
}

// Standard Deviation.
func StandarDeviatiom(vec Vector) float64 {
	return math.Sqrt(Variance(vec))
}

// Inter quartile range.
func InterQuartileRange(vec Vector) float64 {
	return Quantile(vec, 0.75) - Quantile(vec, 0.25)
}

// Covariance
func Covariance(vec1 Vector, vec2 Vector) float64 {
	n := len(vec1)
	return Dot(DeMean(vec1), DeMean(vec2)) / float64(n-1)
}

// Correlation
func Correlation(x, y Vector) float64 {
	stdX := StandarDeviatiom(x)
	stdY := StandarDeviatiom(y)
	if stdX > 0 && stdY > 0 {
		return Covariance(x, y) / (stdX * stdY)
	}
	return 0
}
