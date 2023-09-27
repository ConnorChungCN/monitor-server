package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func NewTempFilePath(fileName string) string {
	d := path.Join(os.TempDir(), "hanglok")
	p := path.Join(d, fileName)
	d = path.Dir(p)
	os.MkdirAll(d, 0755)
	return p
}

func ReadJsonFile(fileName string) (map[string]any, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("open output file[%s] failed, %w", fileName, err)
	}
	defer f.Close()
	// 读取结果
	decoder := json.NewDecoder(f)
	var output map[string]any
	if err := decoder.Decode(&output); err != nil {
		return nil, fmt.Errorf("decode from file[%s] failed, %w", fileName, err)
	}
	return output, nil
}
