package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "smplrstapp/docs"

	"smplrstapp/internal/config"
	httpcontrollers "smplrstapp/internal/controller/http"
	"smplrstapp/internal/infrasctructure/repository"
	"smplrstapp/internal/service"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := startServer(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}

func startServer(ctx context.Context) error {
	mux := mux.NewRouter()

	initComposites(mux)

	srv := &http.Server{
		Addr:    ":8020",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	err := srv.Shutdown(ctxShutDown)
	if err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}

func initComposites(router *mux.Router) {
	mydir, _ := os.Getwd()
	envFile := path.Join(mydir, ".env")
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		envFile = path.Join(mydir, "../../.env")
	}

	fmt.Println(envFile)
	cfg, err := config.Init(envFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("create psql repository")
	dbUrl := cfg.Postgres.GetConnectionString()
	gormDB, err := gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	migrator := repository.NewMigrator(gormDB)
	userRepository := repository.NewUserRepository(gormDB)
	activityRepository := repository.NewActivityRepository(gormDB)
	fmt.Println("migrate database")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("migration user entity")
	fmt.Println("migration user activity")
	migrator.Migrate(ctx)
	fmt.Println("migrate successfull")

	fmt.Println("initialize composite for user")
	userService := service.NewUserService(&userRepository)
	userController := httpcontrollers.NewUserHandler(&userService)
	userController.Register(router)

	fmt.Println("initialize composite for activity")
	activityService := service.NewActivityService(&activityRepository, &userRepository)
	activityController := httpcontrollers.NewActivityHandler(&activityService)
	activityController.Register(router)

	// Swagger
	fmt.Println("initialize swagger")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
