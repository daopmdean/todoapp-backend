package main

import (
	"context"
	"log"
	"os"
	"time"
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/adapter/mongodb"
	"todoapp/internal/adapter/mssqldb"
	"todoapp/internal/common/config"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func main() {
	var userRepo repo.UserRepo
	var todoRepo repo.TodoRepo

	switch os.Getenv("DB_MODE") {
	case "mssql":
		db, err := mssqldb.Connect()
		if err != nil {
			panic("can't connect to database")
		}
		migrateDB(db)

		userRepo = mssqldb.NewUserRepo(db)
		todoRepo = mssqldb.NewTodoRepo(db)
	case "mongodb":
		atlas_uri := config.GetAtlasUri()
		client, err := mongo.NewClient(options.Client().ApplyURI(atlas_uri))
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		database := client.Database("todoapp")

		userCollection := database.Collection("users")
		_, err = userCollection.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: "username", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		todoCollection := database.Collection("todos")
		_, err = todoCollection.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: "id", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)

		userRepo = mongodb.NewUserRepo(userCollection)
		todoRepo = mongodb.NewTodoRepo(todoCollection)
	default:
		userRepo = memdb.NewUserRepo()
		todoRepo = memdb.NewTodoRepo()
	}

	s := rest.NewServer(userRepo, todoRepo)
	p := config.GetPort()
	s.Run(p)
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Todo{})
}
