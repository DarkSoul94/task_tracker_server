package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DarkSoul94/task_tracker_server/task_tracker_server"
	task_tracker_serverhttp "github.com/DarkSoul94/task_tracker_server/task_tracker_server/delivery/http"
	task_tracker_serverrepo "github.com/DarkSoul94/task_tracker_server/task_tracker_server/repo/mysql"
	task_tracker_serverusecase "github.com/DarkSoul94/task_tracker_server/task_tracker_server/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required
	"github.com/spf13/viper"
)

// App ...
type App struct {
	task_tracker_serverUC   task_tracker_server.Usecase
	task_tracker_serverRepo task_tracker_server.Repository
	httpServer              *http.Server
}

// NewApp ...
func NewApp() *App {
	db := initDB()
	repo := task_tracker_serverrepo.NewRepo(db)
	uc := task_tracker_serverusecase.NewUsecase(repo)
	return &App{
		task_tracker_serverUC:   uc,
		task_tracker_serverRepo: repo,
	}
}

// Run run task_tracker_serverlication
func (a *App) Run(port string) error {
	defer a.task_tracker_serverRepo.Close()
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	task_tracker_serverhttp.RegisterHTTPEndpoints(router, a.task_tracker_serverUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *sql.DB {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("app.db.login"),
		viper.GetString("app.db.pass"),
		viper.GetString("app.db.host"),
		viper.GetString("app.db.port"),
		viper.GetString("app.db.name"),
		viper.GetString("app.db.args"),
	)
	db, err := sql.Open(
		"mysql",
		dbString,
	)
	if err != nil {
		panic(err)
	}
	runMigrations(db)
	return db
}

func runMigrations(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		viper.GetString("app.db.name"),
		driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion {
		fmt.Println(err)
	}
}
