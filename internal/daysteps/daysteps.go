package daysteps

import (
	"errors"
	"fmt"
	"main/internal/personaldata"
	"main/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

const (
	StepLength = 0.65
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parsed := strings.Split(datastring, ",")

	if len(parsed) != 2 {
		return errors.New("data string must contains steps and duration")
	}

	ds.Steps, err = strconv.Atoi(parsed[0])
	if err != nil {
		return errors.New("steps count are in an incorrect format")
	}

	ds.Duration, err = time.ParseDuration(parsed[1])
	if err != nil {
		return errors.New("training duration in an incorrect format")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("steps count must be greater than a zero")
	}

	if ds.Duration <= 0 {
		return "", errors.New("duration must be greater than a zero")
	}

	distance := spentenergy.Distance(ds.Steps)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	answr := fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.\n", ds.Steps, distance, calories)

	return answr, nil
}
