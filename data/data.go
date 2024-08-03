package data

import (
    "log"
    "os"
    "path"
    "io/fs"
    "errors"
    "embed"
)

//go:embed shaders/*
//go:embed scripts/*
//go:embed config.toml
var data embed.FS

func ReadFile(filepath string) ([]byte, error) {
    log.Printf("Loading %s", filepath)
    result, err := os.ReadFile(path.Join("data", filepath))

    if err == nil {
        return result, nil
    } else if !errors.Is(err, fs.ErrNotExist) {
        return nil, err
    }

    return data.ReadFile(filepath)
}

