package scripting

import (
	"fmt"
	"log"

    "go.starlark.net/starlark"
    "go.starlark.net/syntax"

    "dcbrwn.io/gogame/data"
)

func Load(
    filename string,
    threadname string,
    api starlark.StringDict,
) (starlark.StringDict, error) {
    thread := &starlark.Thread{
        Name: threadname,
        Print: func(thread *starlark.Thread, msg string) {
            frame := thread.CallFrame(1)
            log.Printf("%s:%d %s", frame.Pos.Filename(), frame.Pos.Line, msg)
        },
        Load: func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
            source, err := data.ReadFile(module)
            if err != nil {
                return nil, fmt.Errorf("failed to read %s: %w", module, err)
            }

            return starlark.ExecFileOptions(
                &syntax.FileOptions{
                    Set: true,
                    While: true,
                    TopLevelControl: true,
                    GlobalReassign: true,

                    // compiler
                    Recursion: true,
                }, thread, module, source, api,
            )
        },
    }

    res, err := thread.Load(thread, filename)
    if err != nil {
        return res, fmt.Errorf("failed to run %s: %w", filename, err)
    }

    return res, nil
}

