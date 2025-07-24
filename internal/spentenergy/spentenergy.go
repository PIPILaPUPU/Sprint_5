package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: done
	if steps <= 0 {
		return 0, errors.New("incorrect data about steps")
	}

	if weight <= 0 || height <= 0 {
		return 0, errors.New("incorrect data about human (height/weight)")
	}

	if duration <= 0 {
		return 0, errors.New("incorret data about duration")
	}

	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: done
	if steps <= 0 {
		return 0, errors.New("ncorrect data about steps")
	}

	if weight <= 0 || height <= 0 {
		return 0, errors.New("incorrect data about human (height/weight)")
	}

	if duration <= 0 {
		return 0, errors.New("incorret data about duration")
	}

	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: done
	if steps < 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: done
	stepLenght := height * stepLengthCoefficient
	return float64(steps) * stepLenght / mInKm
}
