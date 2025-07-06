package timerservice

import (
	"fmt"
	"strings"
)

type TimerService interface {
	GenerateProgressBar(barLen, current, total int) string
}

type TimerServiceDefault struct {
	filledEmoji string
	emptyEmoji  string
}

func New(FilledEmoji, EmptyEmoji string) *TimerServiceDefault {
	return &TimerServiceDefault{
		filledEmoji: FilledEmoji,
		emptyEmoji:  EmptyEmoji,
	}
}

func (t *TimerServiceDefault) GenerateProgressBar(barLen, current, total int) string {
	progress := int(float64(current) / float64(total) * float64(barLen))
	filled := strings.Repeat(t.filledEmoji, barLen-progress)
	empty := strings.Repeat(t.emptyEmoji, progress)
	return fmt.Sprintf("|%s%s|", filled, empty)
}
