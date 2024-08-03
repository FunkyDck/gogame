package engine

import (
    "github.com/BurntSushi/toml"

    "dcbrwn.io/gogame/data"
)

type Config struct {
    Window struct {
        Width int
        Height int
    }
}

func LoadConfig(filepath string) (*Config, error) {
    content, err := data.ReadFile(filepath)
    if err != nil {
        return nil, err
    }

    config := Config{}

	err = toml.Unmarshal(content, &config)

    return &config, nil
}

