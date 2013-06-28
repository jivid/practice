package logging

import (
    "log"
    "os"
    "strings"
)

/*
    Check if a file with the specified path exists
*/
func CheckFileExists (path string) (bool, error) {
    _, err := os.OpenFile(path, os.O_RDWR, 0660)

    if (err != nil) {
        if (os.IsExist(err)) {
            return true, err
        } else {
            return false, nil
        }
    }
    // File could be opened successfully.
    return true, nil
}

/*
    Create a logger with the name provided as an argument
    The file gets created in a log/ directory under the
    current working directory.

    TODO: Allow specifying of a custom logging directory, not
    just a file name
*/
func CreateLogger (name string, level int) (*log.Logger) {
    cwd,_ := os.Getwd()
    log_path := strings.Join([]string{cwd, "log"}, "/")
    log_file := strings.Join([]string{log_path, name}, "/")

    // MkdirAll will not do anything if the directory already exists
    os.MkdirAll(log_path, 0755)
    exists, err := CheckFileExists(log_file)

    if (err != nil) {
        return nil
    }

    // Create file so we can open it
    if (!exists) {
        os.Create(log_file)
    }

    f, _ := os.OpenFile(log_file, os.O_RDWR|os.O_APPEND, 0660)
    return log.New(f, "", level)
}
