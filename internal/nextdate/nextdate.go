package nextdate

import (
	"errors"
	"example/config"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// NextDate возвращает дату и ошибку, исходя из правил указанных в repeat.
func NextDate(now time.Time, date string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", errors.New("правило повторения не указано")
	}

	dayMatched, _ := regexp.MatchString(`d \d{1,3}`, repeat)
	yearMatched, _ := regexp.MatchString(`y`, repeat)

	if dayMatched {
		days, err := strconv.Atoi(strings.TrimPrefix(repeat, "d "))
		if err != nil {
			return "", err
		}

		if days > 400 {
			return "", errors.New("максимальное количество дней должно быть 400")
		}

		parsedDate, err := time.Parse(config.DateFormat, date)
		if err != nil {
			return "", err
		}

		newDate := parsedDate.AddDate(0, 0, days)

		for newDate.Before(now) {
			newDate = newDate.AddDate(0, 0, days)
		}

		return newDate.Format(config.DateFormat), nil
	} else if yearMatched {
		parsedDate, err := time.Parse(config.DateFormat, date)
		if err != nil {
			return "", err
		}

		newDate := parsedDate.AddDate(1, 0, 0)

		for newDate.Before(now) {
			newDate = newDate.AddDate(1, 0, 0)
		}

		return newDate.Format(config.DateFormat), nil
	}
	return "", errors.New("неверный формат повторения")
}
