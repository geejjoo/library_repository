package adapter_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/geejjoo/library_repository/internal/config"
	"github.com/geejjoo/library_repository/internal/db"
	"github.com/geejjoo/library_repository/internal/infrastructure/db/migrate"
	"github.com/geejjoo/library_repository/internal/infrastructure/logs"
	"github.com/geejjoo/library_repository/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func GenerateFakeUser() models.UserDTO {
	return models.UserDTO{
		ID:        gofakeit.Number(0, 1000),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Age:       gofakeit.Number(10, 50),
	}
}

func TestSQLAdapter_Create(t *testing.T) {
	err := godotenv.Load("/home/razzbitb/go/src/gitlab.com/hazhbulat/golibrary/.env")
	conf := config.NewAppConf()
	logger := logs.NewLogger(conf, os.Stdout)
	conf.Init(logger)

	dbx, adapter, err := db.NewSqlDB(conf.DB, logger)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	migrator := migrate.NewMigrator(dbx, conf.DB)
	err = migrator.Migrate(&models.UserDTO{})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	//user := models.UserDTO{
	//	ID:      832,
	//	Name:    "Serenity",
	//	Surname: "Kuhic",
	//	Age:     31,
	//	DeletedAt: null.Time{
	//		Time:  time.Now(),
	//		Valid: true,
	//	},
	//}

	// DELETE
	//err = adapter.Delete(context.Background(), &user, dao.Condition{Equal: map[string]interface{}{
	//	"id": user.GetID(),
	//}})
	//if err != nil {
	//	t.Errorf("Error: %v", err)
	//}

	// LIST
	//var Users []models.UserDTO
	//err = adapter.List(context.Background(), &Users, &models.UserDTO{}, dao.Condition{LimitOffset: &dao.LimitOffset{
	//	Offset: 0,
	//	Limit:  50,
	//}})
	//if err != nil {
	//	return
	//}
	//fmt.Println(Users)

	// UPDATE
	//err = adapter.Update(context.Background(), &user, dao.Condition{Equal: map[string]interface{}{
	//	"id": user.GetID(),
	//}})
	//if err != nil {
	//	t.Errorf("Error: %v", err)
	//}

	//CREATE
	for i := 0; i < 50; i++ {
		user := GenerateFakeUser()
		err = adapter.Create(context.Background(), &user)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
