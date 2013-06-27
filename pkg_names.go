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

func print_pkg(name string) {
    var pkg_line string
    contents,_ := ioutil.ReadFile(name)
    lines := strings.Split(string(contents), "\n")
    for _,l := range lines {
        if(strings.HasPrefix(l, "package")) {
            pkg_line = l
            break
        }
    }
    if (len(pkg_line) == 0) {
        return
    } else {
        pkg_name := strings.Split(pkg_line, " ")[1]
        fmt.Printf("%s: %s\n", name, pkg_name)
    }
    return
}

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
            print_pkg(fi.Name())
        }
    }
}
