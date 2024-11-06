package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Comic представляет комикс
type Comic struct {
	ID          int     `json:"id"`
	ImageURL    string  `json:"imageUrl"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
	IsFavorite  bool    `json:"isFavorite"`
}

// Пример списка комиксов
var comics = []Comic{
	{ID: 1, ImageURL: "https://i.pinimg.com/736x/1d/a4/3d/1da43d8742b591be0dd6557410e701b7.jpg", Title: "Spider-Man: Blue", Author: "Jeph Loeb", Description: "Это трогательная история о Питере Паркере и Гвен Стейси.", Price: 14.99, Category: "Spider-Man", Quantity: 5, IsFavorite: false},
	{ID: 2, ImageURL: "https://cdn1.ozone.ru/s3/multimedia-1-d/6933150625.jpg", Title: "Spider-Man: Kraven's Last Hunt", Author: "J.M. DeMatteis", Description: "Последняя охота Кравена на Человека-паука.", Price: 17.99, Category: "Spider-Man", Quantity: 3, IsFavorite: false},
	{ID: 3, ImageURL: "https://s3.amazonaws.com/www.covernk.com/Covers/L/B/Batman+Year+One/batmanyearonetradepaperback1.jpg", Title: "Batman: Year One", Author: "Frank Miller", Description: "Происхождение Бэтмена.", Price: 19.99, Category: "Batman", Quantity: 7, IsFavorite: false},
	{ID: 4, ImageURL: "https://static.tvtropes.org/pmwiki/pub/images/batman_long_halloween_cover.jpg", Title: "Batman: The Long Halloween", Author: "Jeph Loeb", Description: "Годовая тайна для Бэтмена.", Price: 22.99, Category: "Batman", Quantity: 4, IsFavorite: false},
	{ID: 5, ImageURL: "https://vignette.wikia.nocookie.net/tmnt/images/b/b0/Idw91.jpg/revision/latest?cb=20190214074102", Title: "Teenage Mutant Ninja Turtles: City at War", Author: "Kevin Eastman", Description: "Черепашки сталкиваются с городской войной.", Price: 15.99, Category: "Ninja Turtles", Quantity: 6, IsFavorite: false},
	{ID: 6, ImageURL: "https://i.dailymail.co.uk/1s/2024/04/11/21/83526179-13298943-image-a-142_1712867277901.jpg", Title: "Teenage Mutant Ninja Turtles: The Last Ronin", Author: "Kevin Eastman", Description: "Последний выживший Черепашка ищет мести.", Price: 24.99, Category: "Ninja Turtles", Quantity: 2, IsFavorite: false},
	{ID: 7, ImageURL: "https://static.wikia.nocookie.net/marveldatabase/images/5/5b/True_Believers_X-Men_-_Pyro_Vol_1_1.jpg/revision/latest?cb=20191004182739", Title: "X-Men: Days of Future Past", Author: "Chris Claremont", Description: "Путешествие во времени, чтобы изменить будущее мутантов.", Price: 18.99, Category: "X-Men", Quantity: 4, IsFavorite: false},
	{ID: 8, ImageURL: "https://cdn.marvel.com/u/prod/marvel/i/mg/9/20/58b5d00e39b3c/clean.jpg", Title: "X-Men: The Dark Phoenix Saga", Author: "Chris Claremont", Description: "История о разрушительной силе Темной Феникс.", Price: 21.99, Category: "X-Men", Quantity: 3, IsFavorite: false},
	{ID: 9, ImageURL: "https://static.wikia.nocookie.net/marveldatabase/images/0/05/Infinity_War_Vol_1_2.jpg/revision/latest/scale-to-width-down/650?cb=20190417003856", Title: "Avengers: Infinity War", Author: "Jonathan Hickman", Description: "Битва за судьбу вселенной против Таноса.", Price: 23.99, Category: "Avengers", Quantity: 5, IsFavorite: false},
	{ID: 10, ImageURL: "https://i.pinimg.com/736x/7c/f5/55/7cf55562edf630e4cf5d734ab34da648.jpg", Title: "Avengers: Endgame", Author: "Jason Aaron", Description: "Финальная битва за восстановление вселенной.", Price: 25.99, Category: "Avengers", Quantity: 6, IsFavorite: false},
}

// обработчик для GET-запроса, возвращает список комиксов
func getComicsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список комиксов в JSON
	json.NewEncoder(w).Encode(comics)
}

// обработчик для POST-запроса, добавляет комикс
func createComicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newComic Comic
	err := json.NewDecoder(r.Body).Decode(&newComic)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Comic: %+v\n", newComic)
	var lastID int = len(comics)

	for _, comicItem := range comics {
		if comicItem.ID > lastID {
			lastID = comicItem.ID
		}
	}
	newComic.ID = lastID + 1
	comics = append(comics, newComic)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newComic)
}

// обработчик для получения одного комикса по ID
func getComicByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/comics/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Comic ID", http.StatusBadRequest)
		return
	}

	// Ищем комикс с данным ID
	for _, comic := range comics {
		if comic.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comic)
			return
		}
	}

	// Если комикс не найден
	http.Error(w, "Comic not found", http.StatusNotFound)
}

// удаление комикса по id
func deleteComicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/comics/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Comic ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем комикс с данным ID
	for i, comic := range comics {
		if comic.ID == id {
			// Удаляем комикс из среза
			comics = append(comics[:i], comics[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	// Если комикс не найден
	http.Error(w, "Comic not found", http.StatusNotFound)
}

// Обновление комикса по id
func updateComicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/comics/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Comic ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные комикса
	var updatedComic Comic
	err = json.NewDecoder(r.Body).Decode(&updatedComic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем комикс для обновления
	for i, comic := range comics {
		if comic.ID == id {

			comics[i].ImageURL = updatedComic.ImageURL
			comics[i].Title = updatedComic.Title
			comics[i].Author = updatedComic.Author
			comics[i].Description = updatedComic.Description
			comics[i].Price = updatedComic.Price
			comics[i].Category = updatedComic.Category
			comics[i].Quantity = updatedComic.Quantity
			comics[i].IsFavorite = updatedComic.IsFavorite

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comics[i])
			return
		}
	}

	// Если комикс не найден
	http.Error(w, "Comic not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/comics", getComicsHandler)           // Получить все комиксы
	http.HandleFunc("/comics/create", createComicHandler)  // Создать комикс
	http.HandleFunc("/comics/", getComicByIDHandler)       // Получить комикс по ID
	http.HandleFunc("/comics/update/", updateComicHandler) // Обновить комикс
	http.HandleFunc("/comics/delete/", deleteComicHandler) // Удалить комикс

	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
