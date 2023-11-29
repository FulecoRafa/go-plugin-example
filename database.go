package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

var DATABASE_FILE = "./database.json"

var Database []Todo

func init() {
    data, err := LoadDatabase()
    if err != nil {
        panic(err)
    }
    slog.Debug("Loaded database", "data", data)
    Database = data
}

func LoadDatabase() (data []Todo, err error) {
    // Load from file
    slog.Debug("Loading database", "file", DATABASE_FILE)
    fileData, err := os.ReadFile(DATABASE_FILE)
    if err != nil {
        return
    }
    slog.Debug("Loading database, read file", "file", DATABASE_FILE, "fileData", fileData)

    // Parse JSON
    data, err = UnmarshalJSONToTodoSlice(fileData)
    slog.Debug("Loaded database", "file", DATABASE_FILE, "data", data)
    return
}

func SaveDatabase(data []Todo) {
    content, err := json.Marshal(data)
    if err != nil {
        slog.Error("Could not save database: %w", err)
        return
    }

    err = os.WriteFile(DATABASE_FILE, content, 0644)
    if err != nil {
        slog.Error("Could not save database: %w", err)
        return
    }
}
