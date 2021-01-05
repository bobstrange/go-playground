package repositories

import (
	"context"

	"github.com/bobstrange/go-playground/mongo_example/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = "my_db"
var collName = "users"

func newClient() (*mongo.Client, error) {
	cli, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:37017"))
	if err != nil {
		return nil, err
	}
	err = cli.Connect(context.TODO())
	return cli, err
}

func newDB() (*mongo.Database, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

func newCollection() (*mongo.Collection, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	return db.Collection(collName), nil
}

type userRepo struct{}

func NewUserRepo() UserRepoIface {
	return &userRepo{}
}

type UserRepoIface interface {
	Update(ctx context.Context, data *entities.User) error
	Find(ctx context.Context, filter interface{}) ([]*entities.User, error)
}

func (u *userRepo) Update(ctx context.Context, replacement *entities.User) error {
	filter := bson.M{"id": replacement.ID}
	coll, _ := newCollection()
	// upsert の設定
	opts := options.Replace().SetUpsert(true)
	_, err := coll.ReplaceOne(ctx, filter, replacement, opts)
	return err
}

func (u *userRepo) Find(ctx context.Context, filter interface{}) ([]*entities.User, error) {
	coll, _ := newCollection()
	cur, err := coll.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	var res []*entities.User
	if err := cur.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
