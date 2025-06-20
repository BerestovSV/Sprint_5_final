package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// duration time.Duration — длительность тренировки.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 && height <= 0 {
		return 0, errors.New("weight and height must be greater than zero\n")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero\n")
	}

	meanSpeed := MeanSpeed(steps, duration)

	calories := ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH

	return calories, nil
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}

	meanSpeed := MeanSpeed(steps, duration)

	calories := ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight

	return calories, nil
}

// МeanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func MeanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	avg := Distance(steps) / duration.Hours()
	return avg
}

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func Distance(steps int) float64 {
	dist := (float64(steps) * lenStep) / mInKm
	return dist
}
