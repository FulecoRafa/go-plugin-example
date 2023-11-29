package main

import "fmt"

func RewriteDisplay(otherFields map[string]interface{}, previousMsg string) string {
    if priority, found := otherFields["priority"]; found {
        return fmt.Sprintf("%f - %s", priority.(float64), previousMsg)
    }
    return previousMsg
}
