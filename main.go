package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stderr, "INFO ", log.LstdFlags)
	warningLogger = log.New(os.Stderr, "WARNING ", log.LstdFlags)
	errorLogger = log.New(os.Stderr, "ERROR ", log.LstdFlags)
}

func main() {
	configFile := flag.String("config", "", "config file")
	dataDir := flag.String("data_dir", path.Join(xdg.DataHome, "sshesame"), "data directory")
	flag.Parse()

	configString := ""
	if *configFile != "" {
		configBytes, err := ioutil.ReadFile(*configFile)
		if err != nil {
			errorLogger.Fatalf("Failed to read config file: %v", err)
		}
		configString = string(configBytes)
	}

	cfg, err := getConfig(configString, *dataDir)
	if err != nil {
		errorLogger.Fatalf("Failed to get config: %v", err)
	}

	listener, err := net.Listen("tcp", cfg.Server.ListenAddress)
	if err != nil {
		errorLogger.Fatalf("Failed to listen for connections: %v", err)
	}
	defer listener.Close()

	infoLogger.Printf("Listening on %v", listener.Addr())

	if cfg.Server.MetricsAddress != "" {
		http.Handle("/metrics", promhttp.Handler())
		infoLogger.Printf("Serving metrics on %v", cfg.Server.MetricsAddress)
		go func() {
			if err := http.ListenAndServe(cfg.Server.MetricsAddress, nil); err != nil {
				errorLogger.Fatalf("Failed to serve metrics: %v", err)
			}
		}()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			warningLogger.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn, cfg)
	}
}
