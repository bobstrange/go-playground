package store

import (
	"context"

	"github.com/bobstrange/go-playground/go-web-application-dev/handson/entity"
)

func (r *Repository) ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
				id, title,
				status, created, updated
			FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) AddTask(ctx context.Context, db Execer, t *entity.Task) error {
	t.Created = r.Clocker.Now()
	t.Updated = r.Clocker.Now()
	sql := `INSERT INTO task
				(title, status, created, updated)
			VALUES
				(?, ?, ?, ?);`
	result, err := db.ExecContext(ctx, sql, t.Title, t.Status, t.Created, t.Updated)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}
