package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: done
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: done
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 2 {
		return errors.New("неверный формат данных")
	}

	ds.Steps, err = strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return errors.New("неверные данные о шагах")
	}

	ds.Duration, err = time.ParseDuration(dataSlice[1])
	if err != nil {
		return err
	}
	if ds.Duration <= 0 {
		return errors.New("неверные данные о продолжительности")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: done
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	data := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return data, nil
}
