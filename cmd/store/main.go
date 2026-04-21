package main

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		*configPath = os.Getenv("CONFIG_PATH")
	}
	if *configPath == "" {
		slog.Error("config path required: use --config or $CONFIG_PATH")
		os.Exit(1)
	}

	cfg := Config{configPath: *configPath}

	stopCh := make(chan struct{})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sig
		signal.Stop(sig)
		close(stopCh)
	}()

	os.Exit(RunStoreWithExitCode(cfg, &slogAdapter{}, stopCh))
}

type slogAdapter struct{}

func (s *slogAdapter) add(level, msg string, attrs map[string]string) {
	slogAttrs := make([]any, 0, len(attrs)*2)
	for k, v := range attrs {
		slogAttrs = append(slogAttrs, k, v)
	}

	var lvl slog.Level
	switch level {
	case "DEBUG":
		lvl = slog.LevelDebug
	case "WARN":
		lvl = slog.LevelWarn
	case "ERROR":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	slog.Log(nil, lvl, msg, slogAttrs...)
}
