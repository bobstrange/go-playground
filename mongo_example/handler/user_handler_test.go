package handler_test

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/bobstrange/go-playground/mongo_example/entities"
	"github.com/bobstrange/go-playground/mongo_example/handler"
	"github.com/bobstrange/go-playground/mongo_example/repositories"
)

type RepoMock struct {
	repositories.RepoIface
	result []map[string]interface{}
	err    error
}

func (r *RepoMock) Update(ctx context.Context, data entities.EntityIface) error {
	return nil
}

func (r *RepoMock) Find(ctx context.Context, filter interface{}) ([]map[string]interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.result, nil
}

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}

func TestHandler(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		resultMock := []map[string]interface{}{
			{
				"id":         "test_id",
				"first_name": "First Name",
				"last_name":  "Last Name",
				"nickname":   String("Nickname"),
				"age":        Int(1),
				"created_at": Time(time.Now()),
				"updated_at": Time(time.Now()),
			},
		}
		repo := &RepoMock{
			result: resultMock,
		}
		log.Printf("result: %v\n", repo.result[0])
		res, err := handler.Handle(repo)

		if err != nil {
			t.Fatal("This shouldn't return err but got", err)
		}

		var result []map[string]interface{}
		json.Unmarshal(res, &result)
		if reflect.DeepEqual(result, resultMock) {
			t.Fatalf("Expected %v \nGot %v\n", resultMock, result)
		}
	})
}
