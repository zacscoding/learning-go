package taskman

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

const (
	apiPathPrefix  = "/api/v1/task/"
	htmlPathPrefix = "/task/"
	idPattern      = "/{id:[0-9]+}"
)

var m = NewInMemoryAccessor()
var tmpl = template.Must(template.ParseGlob("html/*.html"))

func main() {
	r := mux.NewRouter()
	r.PathPrefix(htmlPathPrefix).
		Path(idPattern).
		Methods("GET").
		HandlerFunc(htmlHandler)

	s := r.PathPrefix(apiPathPrefix).Subrouter()
	s.HandleFunc(idPattern, apiGetHandler).Methods("GET")
	s.HandleFunc(idPattern, apiPutHandler).Methods("PUT")
	s.HandleFunc("/", apiPostHandler).Methods("POST")
	s.HandleFunc(idPattern, apiDeleteHandler).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8884", nil))
}

func getTasks(r *http.Request) ([]Task, error) {
	var result []Task
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	encodedTasks, ok := r.PostForm["task"]
	if !ok {
		return nil, errors.New("task parameter expected")
	}
	for _, encodedTask := range encodedTasks {
		var t Task
		if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := ID(r.URL.Path[len(htmlPathPrefix):])
		if id == "" {
			return id, errors.New("ID is empty")
		}
		return id, nil
	}
	id, err := getID()
	if err != nil {
		log.Println(err)
		return
	}
	t, err := m.Get(id)
	err = tmpl.ExecuteTemplate(w, "task.html", &Response{
		ID:    id,
		Task:  t,
		Error: ResponseError{err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func apiGetHandler(w http.ResponseWriter, r *http.Request) {
	id := ID(mux.Vars(r)["id"])
	t, err := m.Get(id)
	err = json.NewEncoder(w).Encode(Response{
		ID:    id,
		Task:  t,
		Error: ResponseError{err},
	})
	if err != nil {
		log.Println(err)
	}
}

func apiPostHandler(w http.ResponseWriter, r *http.Request) {

}

func apiPutHandler(w http.ResponseWriter, r *http.Request) {

}

func apiDeleteHandler(w http.ResponseWriter, r *http.Request) {

}
