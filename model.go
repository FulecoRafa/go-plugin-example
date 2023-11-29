package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type Todo struct {
	Content         string `json:"content"`
	Done            bool   `json:"done"`
	OtherProperties map[string]interface{}
}

func (t Todo) String() string {
	var checkmark = " "
	if t.Done {
		checkmark = "X"
	}
	return fmt.Sprintf("[%s] %s", checkmark, t.Content)
}

func UnmarshalJSONToTodoSlice(data []byte) (todos []Todo, err error) {
	raw := []map[string]interface{}{}
	json.Unmarshal(data, &raw)
    slog.Debug("Unmarshalled Json", "raw", raw)

	todos = make([]Todo, 0, len(raw))
	for _, obj := range raw {
		content, found := obj["content"]
		if !found {
			return nil, fmt.Errorf("Missing field `content`")
		}

		done, found := obj["done"]
		if !found {
			return nil, fmt.Errorf("Missing field `done`")
		}

		delete(obj, "content")
		delete(obj, "done")

		todos = append(todos, Todo{
			Content:         content.(string),
			Done:            done.(bool),
			OtherProperties: obj,
		})
	}
	return
}
