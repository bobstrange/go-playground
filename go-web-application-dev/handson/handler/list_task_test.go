package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobstrange/go-playground/go-web-application-dev/handson/entity"
	"github.com/bobstrange/go-playground/go-web-application-dev/handson/handler"
	"github.com/bobstrange/go-playground/go-web-application-dev/handson/store"
	"github.com/bobstrange/go-playground/go-web-application-dev/handson/testutil"
)

func TestListTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}

	t.Run("ok", func(tt *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

		lt := handler.ListTask{
			Store: &store.TaskStore{
				Tasks: map[entity.TaskID]*entity.Task{
					1: {
						ID:     1,
						Title:  "Example Task 1",
						Status: entity.TaskStatusDoing,
					},
					2: {
						ID:     2,
						Title:  "Example Task 2",
						Status: entity.TaskStatusTodo,
					},
				},
			},
		}
		lt.ServeHTTP(w, r)

		resp := w.Result()
		testutil.AssertResponse(tt, resp, http.StatusOK, testutil.LoadFile(tt, "testdata/list_task/ok_rsp.json.golden"))
	})

}
