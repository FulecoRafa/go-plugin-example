package main

import (
	"log/slog"
	"os"
	"path"
	"plugin"
)

var plugins =  []string {
    "priority",
}

var functions [](func(otherFields map[string]interface{}, previousMsg string) string)

func buildPath(pluginName string) string {
    exPath, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    return path.Join(exPath, "plugins", pluginName, "plugin")
}

func init() {
    functions = LoadPlugins()
}

func LoadPlugins() (loadedFunctions [](func(otherFields map[string]interface{}, previousMsg string) string)) {
    loadedFunctions = make([](func(otherFields map[string]interface{}, previousMsg string) string), 0, len(plugins))
    for _, pluginName := range plugins {
        slog.Debug("Loading plugin", "pluginName", pluginName)
        pluginPath := buildPath(pluginName)

        // Load plugin. Fail to load == skip
        pExe, err := plugin.Open(pluginPath)
        if err != nil {
            slog.Warn("Failed to load plugin, skipping...",
                "pluginName", pluginName,
                "pluginPath", pluginPath,
            )
            continue
        }

        // Plugins are expected to implement
        // RewriteDisplay(otherFields map[string]interface{}, previousMsg string) string
        f, err := pExe.Lookup("RewriteDisplay")
        if err != nil {
            slog.Warn("Plugin does not implement expected function RewriteDisplay",
                "pluginName", pluginName,
                "pluginPath", pluginPath,
            )
            continue
        }
        actualF := f.(func(otherFields map[string]interface{}, previousMsg string) string)

        slog.Debug("Plugin loaded successfully", "pluginName", pluginName)
        loadedFunctions = append(loadedFunctions, actualF)
    }
    return
}

func ExecutePlugins(todo Todo) string {
    msg := todo.String()
    for _, f := range functions {
        msg = f(todo.OtherProperties, msg)
    }
    slog.Debug("Plugin parsed to new message", "newMsg", msg)
    return msg
}
