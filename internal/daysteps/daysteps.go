package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Разделяем входную строку на слайс строк
	// Проверяем, что слайс содержит ровно два элемента
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid input format")
	}
	//преобразование слайса (кол-во шагов)
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	//проверка кол-ва шагов на 0
	if steps <= 0 {
		return 0, 0, errors.New("steps must be positive")
	}
	//продолжительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}
	//проверка на 0
	if duration <= 0 {
		return 0, 0, errors.New("duration must be positive")
	}
	//возвращаем переменные с данными
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
