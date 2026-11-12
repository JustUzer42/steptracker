package actioninfo

import (
	"github.com/Yandex-Practicum/tracker/internal/daysteps"
	"github.com/Yandex-Practicum/tracker/internal/trainings"
)

// DataParser is an interface for parsing activity data and generating info strings.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info processes a dataset of activity strings and prints formatted information.
func Info(dataset []string, dp DataParser) {
	for _, item := range dataset {
		err := dp.Parse(item)
		if err != nil {
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			continue
		}

		println(info)
	}
}

// Ensure compile-time interface implementation checks.
var (
	_ DataParser = (*trainings.Training)(nil)
	_ DataParser = (*daysteps.DaySteps)(nil)
)
