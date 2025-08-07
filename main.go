package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Конфигурация сайтов прямо в коде
	sites := map[string]string{
		"mysite1.local": "./site1/static",
		"mysite2.local": "./site2/static",
	}
	
	// Проверяем существование директорий
	for domain, dir := range sites {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			log.Fatalf("Директория %s для сайта %s не существует", dir, domain)
		}
	}

	// Обработчик запросов
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// dir, exists := sites[r.Host]
		// if !exists {
		// 	http.Error(w, "Сайт не найден", http.StatusNotFound)
		// 	return
		// }

		path := filepath.Join(dir, r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(dir, "index.html"))
			return
		}

		http.ServeFile(w, r, path)
	})

	// Запуск сервера
	log.Println("Сервер запущен на порту 80")
	log.Println("Доступные сайты:")
	for domain := range sites {
		log.Printf(" - http://%s", domain)
	}

	log.Fatal(http.ListenAndServe(":80", nil))
}
