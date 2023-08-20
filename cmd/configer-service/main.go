package main

import (
	"configer-service/internal/custom"
	"configer-service/internal/db"
	"configer-service/internal/routing"
	"configer-service/internal/structs"
	"configer-service/pkg/utils"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	cfgFileName := flag.String("cfg", "config-dev", "the name of the config file")
	flag.Parse()

	viper.SetConfigName(*cfgFileName)
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		// TODO: Separate config not found error
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	cfg := structs.CfgData{}
	viper.Unmarshal(&cfg)

	fmt.Println(utils.PrettyPrint(cfg))

	e := echo.New()

	e.Validator = custom.NewCustomValidator()

	// Set number of requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	// e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())
	// e.Use(middleware.Secure())
	// e.Use(middleware.Logger())

	if cfg.Env == "test" {
		os.Remove(cfg.DB.File)
	}

	dbConn := db.GetSQLiteConnection(cfg.DB.File, &gorm.Config{})

	routing.SetRouts(e, dbConn, cfg.Security)

	// autoTLSManager := autocert.Manager{
	// 	Prompt: autocert.AcceptTOS,
	// 	// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
	// 	Cache: autocert.DirCache("/var/www/.cache"),
	// 	//HostPolicy: autocert.HostWhitelist("<DOMAIN>"),
	// }

	// load CA certificate file and add it to list of client CAs
	caCertFile, err := ioutil.ReadFile("./certs/ca.crt")
	if err != nil {
		log.Fatalf("error reading CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCertFile)

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: e,
		TLSConfig: &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
			// MinVersion:               tls.VersionTLS12,
			// CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			// CipherSuites: []uint16{
			// 	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			// 	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			// 	tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// 	tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			// 	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			// 	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			// },
		},
		// &tls.Config{
		// 	//Certificates: nil, // <-- s.ListenAndServeTLS will populate this field
		// 	GetCertificate: autoTLSManager.GetCertificate,
		// 	NextProtos:     []string{acme.ALPNProto},
		// },
		// //ReadTimeout: 30 * time.Second, // use custom timeouts
	}

	fmt.Printf("(HTTPS) Listen on :%d\n", cfg.Port)
	err = s.ListenAndServeTLS(
		"./certs/server-test-20230316.crt",
		"./certs/server-test-20230316.key")
	if err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}

	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
