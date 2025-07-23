package actioninfo

import (
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, str := range dataset {
		err := dp.Parse(str)

		if err != nil {
			log.Println(err)
		}

		dp.ActionInfo()
	}
}
