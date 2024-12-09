package util

import (
    "os"
)

func ReadFileToString(fileName string) (string, error) {
    bytes, err := os.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    str := string(bytes)
    return str, nil
}
