package main

import (
	"fmt"
	"github.com/Venukishore-R/microservice1_auth/endpoints"
	"github.com/Venukishore-R/microservice1_auth/models"
	"github.com/Venukishore-R/microservice1_auth/protos"
	"github.com/Venukishore-R/microservice1_auth/services"
	"github.com/Venukishore-R/microservice1_auth/transports"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	var logger log.Logger

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Log("during", "db", "err", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Log("during", "db", "err", err)
		os.Exit(1)
	}
}

func main() {

	//myFunc := func(token *stdjwt.Token) (interface{}, error) {
	//	return []byte(models.JwtUserKey), nil
	//}
	//
	//myClaims := func() stdjwt.Claims {
	//	return &models.UserClaims{}
	//}
	//
	//parseMiddleware := jwt.NewParser(myFunc, stdjwt.SigningMethodHS256, myClaims)

	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := services.NewLoggerService(logger)
	makeEndpoints := endpoints.MakeEndpoints(service)
	//makeEndpoints.Authenticate = parseMiddleware(makeEndpoints.Authenticate)
	server := transports.NewMyServer(makeEndpoints, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":5005")
	if err != nil {
		logger.Log("during", "listen", "err", err)
		os.Exit(1)
	}

	go func() {
		serverRegistrar := grpc.NewServer()
		protos.RegisterUserServiceServer(serverRegistrar, &server)
		level.Info(logger).Log("msg", "server created successfully")
		serverRegistrar.Serve(grpcListener)
	}()

	logger.Log("exit", <-errs)

}
