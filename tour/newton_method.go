package main

import (
	"errors"
	"math"
)

func NewtonGuess1(k int) float64 {
	return NewtonWithGuess(1)(k)
}

func NewtonGuessHalf(k int) float64 {
	return NewtonWithGuess(k / 2)(k)
}

func NewtonGuessConvergence01(k int) float64 {
	return NewtonWithConvergence(k/2, 0.1)(k)
}

func NewtonGuessConvergence001(k int) float64 {
	return NewtonWithConvergence(k/2, 0.001)(k)
}

func NewtonWithError(k int) (res float64, err error) {
	if k < 0 {
		return -1, errors.New("can't work with negative number")
	}
	return NewtonWithConvergence(k/2, 0.001)(k), nil
}

func NewtonWithGuess(guess int) func(k int) float64 {
	return func(k int) float64 {
		var res = float64(guess)
		for i := 0; i < 10; i++ {
			res -= (res*res - (float64(k))) / (2 * res)
		}
		return res
	}
}

func NewtonWithConvergence(guess int, delta float64) func(k int) float64 {
	return func(k int) float64 {
		var res = float64(guess)
		for d := math.MaxFloat64; d >= delta; d = math.Abs(float64(k) - res*res) {
			res -= (res*res - (float64(k))) / (2 * res)
		}
		return res
	}
}
