package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// Разделяем строку на шаги активность продолжительность и ошибку
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("invalid input format")
	}
	// Преобразуем количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, err
	}
	// Преобразуем продолжительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, err
	}
	// Возвращаем шаги активность продолжительность и отсутсвие ошибки
	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// Вычисляем дистанцию
	distance := (float64(steps) * height * stepLengthCoefficient) / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	speed := distance(steps, height) / duration.Hours()
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// Вызов функции и проверка на ошибку
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	switch trainingType {
	case "Ходьба":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч.\nДистанция %.2f км.\nСкорость:%.2f км/ч\nСожгли калории: %.2fккал\n", "Ходьба", duration.Hours(), dist, speed, calories)
		return result, nil

	case "Бег":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч.\nДистанция %.2f км.\nСкорость:%.2f км/ч\nСожгли калории: %.2fккал\n", "Бег", duration.Hours(), dist, speed, calories)
		return result, nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}
	if weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("data must be positive")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("data must be positive")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes() / minInH) * walkingCaloriesCoefficient
	return calories, nil
}
