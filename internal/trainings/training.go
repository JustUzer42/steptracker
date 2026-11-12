package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training holds data about a training session.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal // embedded
}

// Parse parses a string in format "3456,Ходьба,3h00m" and stores data in the Training struct.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format: expected 3 fields, got %d", len(parts))
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("invalid steps value: %w", err)
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])

	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return fmt.Errorf("invalid duration value: %w", err)
	}
	t.Duration = duration

	return nil
}

// ActionInfo returns a formatted string with training information.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return info, nil
}

// Print outputs the personal data (inherited from embedded Personal).
func (t Training) Print() {
	t.Personal.Print()
}
