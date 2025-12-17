package main

import (
	"ascii-art-web/asciigo"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	mux := http.NewServeMux()

	// Обработка статики
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Основные маршруты
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("✅ Server is running at: http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("❌ Server failed:", err)
		os.Exit(1)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page not found")
		return
	}
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error")
	}
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		renderError(w, http.StatusBadRequest, "Bad request: Unable to parse form")
		return
	}

	// ВАЖНО: trim пробелы — чтобы "пустые" строки (только пробелы) считались пустыми
	input := strings.TrimSpace(r.FormValue("inpt"))
	banner := strings.TrimSpace(r.FormValue("bnr"))

	// Если поле пустое — это 400 (Bad Request)
	if input == "" {
		renderError(w, http.StatusBadRequest, "Bad request: input is empty")
		return
	}
	if banner == "" {
		renderError(w, http.StatusBadRequest, "Bad request: banner is required")
		return
	}

	asciiOutput, err := asciigo.GenerateAsciiArt(input, banner)
	if err != nil {
		// normalize error string
		e := strings.ToLower(err.Error())

		// Ошибки, относящиеся к некорректному запросу (400)
		if strings.Contains(e, "bad request") ||
			strings.Contains(e, "invalid banner") ||
			strings.Contains(e, "not supported") ||
			strings.Contains(e, "non-ascii") ||
			strings.Contains(e, "input contains invalid") ||
			strings.Contains(e, "400") {
			renderError(w, http.StatusBadRequest, "Bad request: "+err.Error())
			return
		}

		// Ошибки, относящиеся к проблемам с файлами/сервером (500)
		if strings.Contains(e, "internal server error") ||
			strings.Contains(e, "banner file") ||
			strings.Contains(e, "file is invalid") ||
			strings.Contains(e, "failed to open") ||
			strings.Contains(e, "500") {
			renderError(w, http.StatusInternalServerError, "Internal server error: "+err.Error())
			return
		}

		// По умолчанию — 500
		renderError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Успех — 200 и показываем результат
	data := struct {
		Output string
	}{
		Output: asciiOutput,
	}

	if err := tpl.ExecuteTemplate(w, "output.html", data); err != nil {
		renderError(w, http.StatusInternalServerError, "Error rendering output")
	}
}

func renderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	data := struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	}
	if err := tpl.ExecuteTemplate(w, "error.html", data); err != nil {
		http.Error(w, message, code)
	}
}
