package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps holds data about daily activity.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal // embedded
}

// Parse parses a string in format "678,0h50m" and stores data in the DaySteps struct.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format: expected 2 fields, got %d", len(parts))
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("invalid steps value: %w", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return fmt.Errorf("invalid duration value: %w", err)
	}
	ds.Duration = duration

	return nil
}

// ActionInfo returns a formatted string with daily activity information.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps,
		distance,
		calories,
	)

	return info, nil
}

// Print outputs the personal data (inherited from embedded Personal).
func (ds DaySteps) Print() {
	ds.Personal.Print()
}
