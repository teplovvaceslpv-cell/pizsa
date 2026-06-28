package main

import (
	"encoding/json"
	"os"
	"strings"
)

// Data - читает JSON файл и возвращает массив объектов
func Data(filename string) []interface{} {
	// 1. Читаем файл
	data, _ := os.ReadFile(filename)
	
	// 2. Создаем пустой массив
	var result []interface{}
	
	// 3. Превращаем JSON в массив объектов
	json.Unmarshal(data, &result)
	
	// 4. Возвращаем массив
	return result
}

// Search - ищет в массиве объекты, содержащие строку q
func Search(data []interface{}, q string) []interface{} {
	// Создаем массив для результатов
	results := []interface{}{}
	
	// Перебираем каждый объект в массиве
	for _, item := range data {
		// Превращаем объект в map (словарь)
		obj := item.(map[string]interface{})
		
		// Достаем поля
		url := obj["url"].(string)
		name := obj["name"].(string)
		desc := obj["description"].(string)
		
		// Проверяем, содержит ли какое-то поле нашу строку
		if strings.Contains(url, q) || 
		   strings.Contains(name, q) || 
		   strings.Contains(desc, q) {
			// Если да - добавляем в результаты
			results = append(results, item)
		}
	}
	
	return results
}