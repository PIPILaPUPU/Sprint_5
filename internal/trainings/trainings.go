package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: done
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: done
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 3 {
		return errors.New("incorrect data format")
	}

	t.Steps, err = strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}

	if t.Steps <= 0 {
		return errors.New("incorrect amount of steps")
	}

	t.TrainingType = dataSlice[1]

	t.Duration, err = time.ParseDuration(dataSlice[2])
	if err != nil {
		return err
	}

	if t.Duration <= 0 {
		return errors.New("incorrect format about duration")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: done
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		spentedCaloriesWalk, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentedCaloriesWalk), nil
	case "Бег":
		spentedCaloriesRun, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentedCaloriesRun), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
