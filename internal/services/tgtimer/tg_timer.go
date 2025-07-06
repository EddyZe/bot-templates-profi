package tgtimer

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	"strings"
)

type TgTimer interface {
	GenerateProgressBar(pattern string, barLen, current, total, tick int) (string, []models.MessageEntity)
}

type TgTimerDefault struct {
	filledEmoji   string
	emptyEmoji    string
	filledEmojiId string
	emptyEmojiId  string
}

func New(filledEmoji, emptyEmoji, filledEmojiId, emptyEmojiId string) *TgTimerDefault {
	return &TgTimerDefault{
		filledEmoji:   filledEmoji,
		emptyEmoji:    emptyEmoji,
		filledEmojiId: filledEmojiId,
		emptyEmojiId:  emptyEmojiId,
	}
}

func (t *TgTimerDefault) GenerateProgressBar(pattern string, barLen, current, total, tick int) (string, []models.MessageEntity) {
	progress := int(float64(current) / float64(total) * float64(barLen))
	filledCount := barLen - progress
	emptyCount := progress

	// Чередование кастомных эмодзи для анимации
	emojis := []string{t.filledEmoji, "🪙", "🗝️"} // Заглушки для визуализации
	filled := strings.Repeat(emojis[tick%len(emojis)], filledCount)
	empty := strings.Repeat(t.emptyEmoji, emptyCount)
	bar := fmt.Sprintf("%s%s", filled, empty)

	// Создаём сущности для кастомных эмодзи
	entities := make([]models.MessageEntity, 0)
	filledEmojiLen := len([]rune(emojis[tick%len(emojis)])) // Длина одного эмодзи в UTF-16
	emptyEmojiLen := len([]rune(t.emptyEmoji))

	// Добавляем сущности для заполненной части
	for i := 0; i < filledCount; i++ {
		entities = append(entities, models.MessageEntity{
			Type:          "custom_emoji",
			Offset:        len([]rune(pattern)) + i*filledEmojiLen,
			Length:        filledEmojiLen,
			CustomEmojiID: t.filledEmojiId,
		})
	}

	// Добавляем сущности для пустой части
	for i := 0; i < emptyCount; i++ {
		entities = append(entities, models.MessageEntity{
			Type:          "custom_emoji",
			Offset:        len([]rune(pattern)) + filledCount*filledEmojiLen + i*emptyEmojiLen,
			Length:        emptyEmojiLen,
			CustomEmojiID: t.emptyEmojiId,
		})
	}

	return bar, entities
}
