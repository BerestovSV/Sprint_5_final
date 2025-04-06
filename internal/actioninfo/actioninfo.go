package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}
