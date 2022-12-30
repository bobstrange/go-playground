package handler

import (
	"net/http"

	"github.com/bobstrange/go-playground/go-web-application-dev/handson/entity"
	"github.com/bobstrange/go-playground/go-web-application-dev/handson/store"
)

type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	rsp := []task{}
	for i, t := range tasks {
		rsp[i] = task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		}
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
