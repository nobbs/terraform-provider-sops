package utils

import (
	"encoding/json"
	"strings"

	dotenv "github.com/joho/godotenv"
	"github.com/wlevene/ini"
	"gopkg.in/yaml.v3"
)

func IsValidFormat(format string) bool {
	return format == "yaml" || format == "json" || format == "dotenv" || format == "ini" || format == "binary"
}

func IsYAMLFile(path string) bool {
	return strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml")
}

func IsJSONFile(path string) bool {
	return strings.HasSuffix(path, ".json")
}

func IsEnvFile(path string) bool {
	return strings.HasSuffix(path, ".env")
}

func IsIniFile(path string) bool {
	return strings.HasSuffix(path, ".ini")
}

func FileFormatFromPath(path string) string {
	if IsYAMLFile(path) {
		return "yaml"
	} else if IsJSONFile(path) {
		return "json"
	} else if IsEnvFile(path) {
		return "dotenv"
	} else if IsIniFile(path) {
		return "ini"
	}

	return "binary"
}

func ReadYAML(data []byte) ([]byte, error) {
	var v any
	err := yaml.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return json.Marshal(v)
}

func ReadJSON(data []byte) ([]byte, error) {
	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return json.Marshal(v)
}

func ReadINI(data []byte) ([]byte, error) {
	x := ini.New().Load(data)
	if err := x.Err(); err != nil {
		return nil, err
	}

	v := x.Marshal2Json()
	if err := x.Err(); err != nil {
		return nil, err
	}

	return v, nil
}

func ReadENV(data []byte) ([]byte, error) {
	v, err := dotenv.UnmarshalBytes(data)
	if err != nil {
		return nil, err
	}

	return json.Marshal(v)
}
