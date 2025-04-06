package trainings

import (
	"errors"
	"fmt"
	"main/internal/personaldata"
	"main/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	parsed := strings.Split(datastring, ",")

	if len(parsed) != 3 {
		return errors.New("data string must contains steps, type and duration")
	}

	t.Steps, err = strconv.Atoi(parsed[0])

	if err != nil {
		return errors.New("steps count are in an incorrect format")
	}

	switch parsed[1] {
	case "Бег":
		t.TrainingType = parsed[1]
	case "Ходьба":
		t.TrainingType = parsed[1]
	default:
		return errors.New("unknown training type")
	}

	t.Duration, err = time.ParseDuration(parsed[2])
	if err != nil {
		return errors.New("training duration in an incorrect format")
	}
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)
	if distance <= 0 {
		return "", errors.New("distance is zero")
	}

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "":
		return "", errors.New("unknown training type")
	}
	if err != nil {
		return "", err
	}

	answer := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч.\nСожгли калорий: %.2f ккал.\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)
	return answer, nil
}
