package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/GenkiHirano/go-grpc-base/database"
	"github.com/GenkiHirano/go-grpc-base/internal/config"
	"github.com/GenkiHirano/go-grpc-base/internal/gen/sample/v1/samplev1connect"
	"github.com/GenkiHirano/go-grpc-base/internal/interface-adapter/controller"
	"github.com/GenkiHirano/go-grpc-base/internal/usecase"
)

func run(ctx context.Context) error {
	fmt.Println("🔥")

	// if err := logging.NewZap(appName, logging.LogTypeApp); err != nil {
	// 	return err
	// }

	cfg, err := config.Init(ctx)
	if err != nil {
		return err
	}

	db, err := database.Init(cfg.DB)
	if err != nil {
		return err
	}

	defer db.Close()

	fmt.Println("db: ", db)

	sampleHandlerOption := []connect.HandlerOption{
		connect.WithInterceptors(
		// interceptor.NewErrorLog().WrapErrorAndSetLogContext(),
		// 現在時刻
		),
	}

	// gatewaySampleUser := gateway.NewSampleUser()

	usecaseSampleUser := usecase.NewSampleUser(db)

	sampleAPI := controller.NewSample(
		usecaseSampleUser,
	)

	mux := http.NewServeMux()
	mux.Handle(samplev1connect.NewSampleServiceHandler(sampleAPI, sampleHandlerOption...))

	services := []string{
		samplev1connect.SampleServiceName,
	}

	checker := grpchealth.NewStaticChecker(services...)
	mux.Handle(grpchealth.NewHandler(checker))

	reflector := grpcreflect.NewStaticReflector(services...)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle("/healthcheck", http.HandlerFunc(healthCheck))

	// corsHandler := cors.New(cors.Options{
	// 	Debug:            false,
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowedOrigins:   []string{"http://localhost:3000"},
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

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// TODO: ログ出力
	fmt.Fprintf(w, "ok")
}
