package jsonhelper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// map[resourceType]map[model]PropertyMap
type PropertyMap struct {
	Addr       string `json:"addr"`
	LinkGithub string `json:"link_github,omitempty"`
	Ref        string `json:"ref"`
	LinkLocal  string `json:"link_local,omitempty"`
}

// return map[name]map[appAddr]PropertyMap
func ParseBridgeFile(path string) (map[string]map[string]PropertyMap, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("open file: %v", err)
	}

	defer f.Close()

	jsonByte, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}

	var bridgeMap map[string]map[string]PropertyMap
	if err := json.Unmarshal(jsonByte, &bridgeMap); err != nil {
		return nil, fmt.Errorf("unmarshal json: %v", err)
	}

	return bridgeMap, nil
}