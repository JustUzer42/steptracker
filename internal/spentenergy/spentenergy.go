package spentenergy

import (
	"errors"
	"time"
)

const (
	stepLengthCoefficient        = 0.415
	mInKm                        = 1000.0
	minInH                       = 60.0
	walkingCaloriesCoefficient   = 0.9
)

var (
	ErrInvalidSteps   = errors.New("invalid steps: must be non-negative")
	ErrInvalidWeight  = errors.New("invalid weight: must be positive")
	ErrInvalidHeight  = errors.New("invalid height: must be positive")
	ErrInvalidDuration = errors.New("invalid duration: must be positive")
)

// Distance calculates the distance in kilometers based on steps and user height.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}

// MeanSpeed calculates the average speed in km/h.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// RunningSpentCalories calculates calories burned during running.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 {
		return 0, ErrInvalidSteps
	}
	if weight <= 0 {
		return 0, ErrInvalidWeight
	}
	if height <= 0 {
		return 0, ErrInvalidHeight
	}
	if duration <= 0 {
		return 0, ErrInvalidDuration
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH

	return calories, nil
}

// WalkingSpentCalories calculates calories burned during walking.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 {
		return 0, ErrInvalidSteps
	}
	if weight <= 0 {
		return 0, ErrInvalidWeight
	}
	if height <= 0 {
		return 0, ErrInvalidHeight
	}
	if duration <= 0 {
		return 0, ErrInvalidDuration
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH * walkingCaloriesCoefficient

	return calories, nil
}
