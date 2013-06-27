/*
  Take a filepath as input. Read all files in that directory,
  if the file is a valid Golang file, print out its package name
*/
package main

import (
    "fmt"
    "os"
    "flag"
    "io/ioutil"
    "strings"
)

func main() {
    flag.Parse()
    path := flag.Arg(0)

    dir, _ := os.Open(path)
    defer dir.Close()

    fileInfos, _ := dir.Readdir(-1)
    for _, fi := range fileInfos {
        if (fi.Mode().IsRegular() &&
            !strings.HasPrefix(fi.Name(), ".") &&
            strings.HasSuffix(fi.Name(), ".go")) {
            f,_ := os.Open(fi.Name())
            defer f.Close()
            contents, _ := ioutil.ReadFile(fi.Name())
            first_line := strings.Split(string(contents), "\n")[0]
            if (strings.HasPrefix(first_line, "package")) {
                pkg_name := strings.Split(first_line, " ")[1]
                fmt.Printf("%s: %s\n" , fi.Name(), pkg_name)
            }
        }
    }
}
