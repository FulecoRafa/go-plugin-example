package main

import (
	"fmt"
	"log/slog"
	"os"
)

func init() {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelInfo)
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))
	slog.SetDefault(logger)

    slog.Debug("Force loading database")
    var err error
    Database, err = LoadDatabase()
    if err != nil {
        panic(err)
    }
}

func main() {
    slog.Debug("Loaded database", "Database", Database)
    fmt.Println("### My Todos ###")
	for _, todo := range Database {
		fmt.Println(ExecutePlugins(todo))
	}
}
