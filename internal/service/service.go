package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	probel := strings.TrimSpace(input) //убрали пробелы
	if probel == "" {
		return "", errors.New("пустая строка") // проверка на пустую строку, если пустая ошика
	}
	deliteMorse := strings.Trim(probel, ".- ") // убираем из текста символы морзы
	if deliteMorse != "" {
		return morse.ToMorse(probel), nil // если остались буквы , переводим в морзу
	}
	return morse.ToText(probel), nil // если нет , то это морза
}
