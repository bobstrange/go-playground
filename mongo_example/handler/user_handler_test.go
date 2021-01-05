package handler_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/bobstrange/go-playground/mongo_example/entities"
	"github.com/bobstrange/go-playground/mongo_example/handler"
	"github.com/bobstrange/go-playground/mongo_example/repositories"
)

type UserRepoMock struct {
	repositories.UserRepoIface
	result []*entities.User
	err    error
}

func (u *UserRepoMock) Update(ctx context.Context, data *entities.User) error {
	return nil
}

func (u *UserRepoMock) Find(ctx context.Context, filter interface{}) ([]*entities.User, error) {
	if u.err != nil {
		return nil, u.err
	}
	return u.result, nil
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
		resultMock := []*entities.User{
			{
				ID:        "test_id",
				FirstName: "First Name",
				LastName:  "Last Name",
				Nickname:  String("Nickname"),
				Age:       Int(1),
				CreatedAt: Time(time.Now()),
				UpdatedAt: Time(time.Now()),
			},
		}
		repo := &UserRepoMock{
			result: resultMock,
		}
		log.Printf("result: %v\n", repo.result[0])
		res, err := handler.Handle(repo)

		if err != nil {
			t.Fatal("This shouldn't return err but got", err)
		}

		b, _ := json.Marshal(resultMock)
		if string(b) != string(res) {
			t.Fatalf("Expected %s but got %s\n", string(b), string(res))
		}
	})
}
