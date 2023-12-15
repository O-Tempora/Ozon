package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/O-Tempora/Ozon/config"
	"github.com/O-Tempora/Ozon/internal/api/shortener_v1"
	"github.com/O-Tempora/Ozon/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfigPath = "config/default.yaml"
)

var (
	configPath string
	useDB      bool
)

func init() {
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to config file")
	flag.BoolVar(&useDB, "db", false, "Use PostgreSQL database or stick with in-memory solution")
}

func main() {
	flag.Parse()
	config, err := parseConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.CreateServer(useDB, config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()
	reflection.Register(s)
	shortener_v1.RegisterShortenerServiceServer(s, server)

	server.Logger.Info().Msgf("TCP server started on %s:%d", config.Host, config.Port)
	if err = s.Serve(lis); err != nil {
		server.Logger.Fatal().Msgf("Server start error: %s", err.Error())
	}
}

func parseConfig(configPath string) (*config.Config, error) {
	config := &config.Config{}
	if configPath == "" {
		configPath = defaultConfigPath
	}
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %w", err)
	}
	if err = yaml.Unmarshal(bytes, config); err != nil {
		return nil, fmt.Errorf("Failed to parse config: %w", err)
	}
	return config, nil
}
