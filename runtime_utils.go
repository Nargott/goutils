package goutils

import (
	"fmt"
	"runtime"
	"strings"
)

func GetFuncName(skip int) (string, error) {
	fpcs := make([]uintptr, 1)
	_ = runtime.Callers(skip, fpcs)
	if fpcs[0] < 1 {
		return "", fmt.Errorf("bad func pc ptr")
	}
	caller := runtime.FuncForPC(fpcs[0] - 1)
	path := strings.Split(caller.Name(), ".")
	if len(path) < 1 {
		return "", fmt.Errorf("bad path returned")
	}

	return path[len(path)-1], nil
}
