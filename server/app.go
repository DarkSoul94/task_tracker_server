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

	"github.com/DarkSoul94/task_tracker_server/auth"
	authHttp "github.com/DarkSoul94/task_tracker_server/auth/delivery/http"
	authUsecase "github.com/DarkSoul94/task_tracker_server/auth/usecase"

	"github.com/DarkSoul94/task_tracker_server/tasks"
	tasksHttp "github.com/DarkSoul94/task_tracker_server/tasks/delivery/http"
	tasksRepository "github.com/DarkSoul94/task_tracker_server/tasks/repo/mysql"
	tasksUsecase "github.com/DarkSoul94/task_tracker_server/tasks/usecase"

	"github.com/DarkSoul94/task_tracker_server/user_manager"
	user_managerHTTP "github.com/DarkSoul94/task_tracker_server/user_manager/delivery/http"
	user_managerRepository "github.com/DarkSoul94/task_tracker_server/user_manager/repo/mysql"
	user_managerUsecase "github.com/DarkSoul94/task_tracker_server/user_manager/usecase"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required
	"github.com/spf13/viper"
)

// App ...
type App struct {
	authUC auth.AuthUC

	uManagerUC   user_manager.UserManagerUC
	uManagerRepo user_manager.UserManagerRepo

	tasksUC   tasks.TasksUC
	tasksRepo tasks.TasksRepo

	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	db := initDB()

	uManagerRepo := user_managerRepository.NewRepo(db)
	uManagerUC := user_managerUsecase.NewUsecase(uManagerRepo)

	authUC := authUsecase.NewUsecase(
		uManagerUC,
		viper.GetString("app.auth.secret_key"),
		[]byte(viper.GetString("app.auth.signing_key")),
		viper.GetDuration("app.auth.ttl"))

	tasksRepo := tasksRepository.NewRepo(db)
	tasksUC := tasksUsecase.NewUsecase(tasksRepo, uManagerUC)

	return &App{
		authUC:       authUC,
		uManagerUC:   uManagerUC,
		uManagerRepo: uManagerRepo,
		tasksUC:      tasksUC,
		tasksRepo:    tasksRepo,
	}
}

// Run run task_tracker_serverlication
func (a *App) Run(port string) error {
	defer a.closeDB()

	router := gin.New()
	if viper.GetBool("app.release") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())
	apiRouter := router.Group("/task_tracker")

	authHttp.RegisterHTTPEndpoints(apiRouter, a.authUC)
	authHttpMiddleware := authHttp.NewAuthMiddleware(a.authUC)
	tasksHttp.RegisterHTTPEndpoints(apiRouter, a.tasksUC, authHttpMiddleware)
	user_managerHTTP.RegisterHTTPEndpoints(apiRouter, a.uManagerUC, authHttpMiddleware)

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

func (a *App) closeDB() {
	a.tasksRepo.Close()
	a.uManagerRepo.Close()
}
