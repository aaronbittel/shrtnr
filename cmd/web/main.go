package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	sessionManager *scs.SessionManager

	memStore map[string]string

	logger *slog.Logger
	debug  bool
}

func main() {
	var logCfg loggerConfig

	port := flag.Int("port", 8080, "specify server port")
	flag.Func("log-level", "set the log-level (default: info)", logCfg.parseLogLevel)
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logCfg.logLevel,
	}))

	app := &application{
		sessionManager: scs.New(),
		memStore:       make(map[string]string),

		logger: logger,
		debug:  logCfg.debug,
	}

	slog.Info("server starting", "port", *port, "logLevel", logCfg.logLevel)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), app.routes())
}

type loggerConfig struct {
	logLevel slog.Level
	debug    bool
}

func (l *loggerConfig) parseLogLevel(level string) error {
	switch strings.TrimSpace(strings.ToLower(level)) {
	case "debug":
		l.debug = true
		l.logLevel = slog.LevelDebug
	case "info":
		l.logLevel = slog.LevelInfo
	case "warn":
		l.logLevel = slog.LevelWarn
	case "error":
		l.logLevel = slog.LevelError
	default:
		return fmt.Errorf("unknown log level %q", level)
	}
	return nil
}
