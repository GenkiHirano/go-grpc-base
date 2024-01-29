package main

import (
	"context"
	"fmt"

	"github.com/GenkiHirano/go-grpc-base/internal/config"
)

// func initDB(conf config.DBConfig) (*bun.DB, error) {
// 	sqlDB, err := sql.Open(
// 		conf.DBDriver,
// 		fmt.Sprintf(
// 			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
// 			conf.DBUsername, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName,
// 		),
// 	)
// 	if err != nil {
// 		return nil, apperror.WrapWithMessage(
// 			err,
// 			apperror.NewDetail(apperror.CodeInternal, "failed to open db"),
// 		)
// 	}

// 	if err := sqlDB.Ping(); err != nil {
// 		return nil, apperror.WrapWithMessage(
// 			err,
// 			apperror.NewDetail(apperror.CodeInternal, "invalid connection to db"),
// 		)
// 	}

// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetMaxOpenConns(50)
// 	sqlDB.SetConnMaxLifetime(300 * time.Second)
// 	boil.SetDB(sqlDB)

// 	return bun.NewDB(sqlDB, mysqldialect.New()), nil
// }

// func healthCheck(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "ok")
// }

// func getAppNameFromEnvironmentVariable() (string, error) {
// 	appName := os.Getenv("AI_FORTUNE_APP_NAME")
// 	if appName == "" {
// 		return "", apperror.NewError(apperror.NewDetail(
// 			apperror.CodeInternal,
// 			"AI_FORTUNE_APP_NAME not set, this value is required",
// 		))
// 	}

// 	return appName, nil
// }

func run(ctx context.Context) error {
	fmt.Println("🔥")
	// appName, err := getAppNameFromEnvironmentVariable()
	// if err != nil {
	// 	return err
	// }

	// if err := logging.NewZap(appName, logging.LogTypeApp); err != nil {
	// 	return err
	// }

	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	fmt.Println("cfg: ", cfg)

	// db, err := initDB(cfg.DB)
	// if err != nil {
	// 	return err
	// }

	// defer db.Close()

	// awsCfg, err := awsconfig.LoadDefaultConfig(ctx)
	// if err != nil {
	// 	return apperror.WrapWithMessage(err, apperror.NewDetail(apperror.CodeInternal, "failed to load aws config"))
	// }

	// dynamodbClient := dynamodb.NewFromConfig(awsCfg)
	// s3Client := s3.NewPresignClient(s3.NewFromConfig(awsCfg))

	// linebotClient, err := linebot.New(cfg.LINE.Messaging.ChannelSecret, cfg.LINE.Messaging.AccessToken)
	// if err != nil {
	// 	return apperror.WrapWithMessage(err, apperror.NewDetail(apperror.CodeInternal, "failed to line bot client"))
	// }

	// llmURL := &url.URL{
	// 	Scheme: "http",
	// 	Host:   net.JoinHostPort(cfg.LLM.AppName, cfg.LLM.AppPort),
	// }
	// connectHTTPClient := http.Client{
	// 	Transport: &http2.Transport{
	// 		AllowHTTP: true,
	// 		DialTLSContext: func(ctx context.Context, network, addr string, t *tls.Config) (net.Conn, error) {
	// 			return net.Dial(network, addr)
	// 		},
	// 	},
	// }
	// llmClient := llmv1connect.NewLLMServiceClient(&connectHTTPClient, llmURL.String(), connect.WithGRPC())

	// consumerHandlerOption := []connect.HandlerOption{
	// 	connect.WithInterceptors(
	// 		interceptor.NewErrorLog().WrapErrorAndSetLogContext(),
	// 		interceptor.NewConsumerAuth(cookie.New(), db, gateway.NewConsumer(), gateway.NewConsumerAuth(cfg.LINE.Login)).LINEAuth(),
	// 	),
	// }

	// consumerToBHandlerOption := []connect.HandlerOption{
	// 	connect.WithInterceptors(
	// 		interceptor.NewErrorLog().WrapErrorAndSetLogContext(),
	// 	),
	// }

	// gatewayPlan := gateway.NewPlan()
	// univapayConf := cfg.Subscription.Univapay
	// gatewaySubscription := gateway.NewSubscription(univapayConf.StoreID, univapayConf.AppToken, univapayConf.TokenSecret)
	// gatewayCompany := gateway.NewCompany(cfg.AWS.S3.BucketName, s3Client)
	// gatewayConsumer := gateway.NewConsumer()
	// gatewayConsumerAuth := gateway.NewConsumerAuth(cfg.LINE.Login)
	// gatewayFortune := gateway.NewFortune(cfg.Fortune, cfg.AWS.DynamoDB.FriendTarotHistoryTableName, cfg.AWS.S3.BucketName, llmClient, dynamodbClient, s3Client)
	// gatewayChat := gateway.NewChat(dynamodbClient, cfg.AWS.DynamoDB.ChatHistoryTableName, linebotClient, llmClient)
	// gatewayNotification := gateway.NewNotification(cfg.Notification.SlackConfig.WebhookURL)

	// usecaseConsumer := usecase.NewConsumer(db, gatewayConsumer, gatewayConsumerAuth, gatewaySubscription, gatewaySubscription, gatewayPlan, gatewayNotification)
	// usecaseConsumerAuth := usecase.NewConsumerAuth(db, gatewayConsumer, gatewayConsumerAuth, gatewaySubscription, gatewayPlan)
	// usecaseFortune := usecase.NewFortune(db, gatewayConsumer, gatewayFortune, gatewayChat, gatewayPlan)
	// usecaseSubscription := usecase.NewSubscription(db, gatewayConsumer, gatewayPlan, gatewaySubscription, gatewayNotification)

	// consumerAPI := controller.NewConsumer(
	// 	cookie.New(),
	// 	db,
	// 	cfg.LLM,
	// 	cfg.Subscription.Univapay,
	// 	gatewayConsumerAuth,
	// 	gatewayChat,
	// 	gatewayConsumer,
	// 	gatewayFortune,
	// 	gatewayPlan,
	// 	usecaseConsumerAuth,
	// 	usecaseFortune,
	// 	usecaseConsumer,
	// 	usecaseSubscription,
	// )

	// consumerToBAPI := controller.NewConsumerToB(cfg.LLM, usecaseFortune, gatewayFortune, gatewayCompany, db)

	// mux := http.NewServeMux()
	// mux.Handle(aifortunev1connect.NewConsumerServiceHandler(consumerAPI, consumerHandlerOption...))
	// mux.Handle(aifortunev1connect.NewConsumerToBServiceHandler(consumerToBAPI, consumerToBHandlerOption...))

	// services := []string{
	// 	aifortunev1connect.ConsumerServiceName,
	// 	aifortunev1connect.ConsumerToBServiceName,
	// }

	// checker := grpchealth.NewStaticChecker(services...)
	// mux.Handle(grpchealth.NewHandler(checker))

	// reflector := grpcreflect.NewStaticReflector(services...)
	// mux.Handle(grpcreflect.NewHandlerV1(reflector))
	// mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// mux.Handle("/healthcheck", http.HandlerFunc(healthCheck))

	// corsHandler := cors.New(cors.Options{
	// 	Debug:            false,
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowedOrigins:   []string{"https://localhost:3001", "http://localhost:3000", "http://localhost:3001", "http://localhost:3002", "https://dev-ai-fortune.vercel.app"},
	// 	AllowCredentials: true,
	// 	AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	// 	MaxAge:           1000,
	// }).Handler(h2c.NewHandler(mux, &http2.Server{}))

	// srv := &http.Server{
	// 	Addr:              ":" + cfg.App.Port,
	// 	Handler:           corsHandler,
	// 	ReadHeaderTimeout: time.Second,
	// 	ReadTimeout:       5 * time.Minute,
	// 	WriteTimeout:      5 * time.Minute,
	// 	MaxHeaderBytes:    8 * 1024,
	// }

	// eg, gctx := errgroup.WithContext(ctx)
	// eg.Go(func() error {
	// 	<-gctx.Done()
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 	defer cancel()
	// 	return srv.Shutdown(ctx)
	// })

	// eg.Go(func() error {
	// 	if cfg.Environment.Environment == config.Local {
	// 		if err := srv.ListenAndServeTLS("localhost.pem", "localhost-key.pem"); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 			return err
	// 		}
	// 	} else {
	// 		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 			return err
	// 		}
	// 	}
	// 	return nil
	// })

	// return eg.Wait()
	return nil
}
