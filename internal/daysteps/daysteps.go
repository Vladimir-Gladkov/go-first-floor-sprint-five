package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	//Вызов функции и проверка на ошибку
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//проверяем кол-во шагов
	if steps <= 0 {
		return ""
	}
	//Вычисляем дистанцию
	distance := float64(steps) * stepLength
	distanceKm := distance / float64(mInKm)
	//Расчет калорий
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKm, calories)

}
