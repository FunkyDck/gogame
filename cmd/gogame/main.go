package main

import (
    "dcbrwn.io/gogame/engine"
)

func main() {
    eng, err := engine.NewEngine("config.toml")
    if err != nil {
        panic(err)
    }

    err = eng.Run()
    if err != nil {
        panic(err)
    }
}

