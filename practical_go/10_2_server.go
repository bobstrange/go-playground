package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/go-playground/validator"
)

type Comment struct {
	Message  string `validate:"required,min=1,max=140" json:"message"`
	UserName string `validate:"required,min=1,max=30" json:"user_name"`
}

func main() {
	mutex := &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			mutex.RLock()

			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()
		case http.MethodPost:
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			v := validator.New()
			if err := v.Struct(c); err != nil {
				var out []string
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					for _, fe := range ve {
						switch fe.Field() {
						case "Message":
							out = append(out, fmt.Sprintf("message should be 1 to 140"))
						case "UserName":
							out = append(out, fmt.Sprintf("user_name should be 1 to 30"))
						}
					}
				}
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, strings.Join(out, ",")), http.StatusBadRequest)
				return
			}

			mutex.Lock()
			comments = append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status":"created"}`))
		default:
			http.Error(w, `{"status":"method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusBadRequest)
				return
			}

			word := r.FormValue("searchword")
			log.Printf("searchword = %s\n", word)

			words, ok := r.Form["searchword"]
			log.Printf("search words = %v has values %v\n", words, ok)

			log.Print("all params")
			for key, values := range r.Form {
				log.Printf("%s = %v\n", key, values)
			}
		}
	})

	// curl -X POST -F file=@README.md -F data='{"key": "value"}' http://localhost:8080/upload
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			err := r.ParseMultipartForm(32 * 1024 * 1024)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusBadRequest)
				return
			}
			f, h, err := r.FormFile("file")
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusBadRequest)
				return
			}
			log.Println(h.Filename)
			o, err := os.Create(h.Filename)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			defer o.Close()

			_, err = io.Copy(o, f)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}

			value := r.PostFormValue("data")
			log.Printf("data = %s\n", value)
		}
	})

	http.ListenAndServe(":8080", nil)
}
