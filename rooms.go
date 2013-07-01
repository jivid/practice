package main

import (
    "fmt"
    "logging"
)

type Furniture struct{
    name string
}

type Room struct {
    name string
    capacity int
    items []Furniture
}

func (r Room) ShowContents() {
    fmt.Printf("Name: %s\n", r.name)
    fmt.Printf("\tCapacity: %d\n", r.capacity)
    fmt.Printf("\tItems: [")
    for _,v := range r.items {
        fmt.Printf(" %s ", v.name)
    }
    fmt.Printf("]\n");
}

var (
    r_name string
    r_cap int
)

func main(){
    logfile := logging.CreateLogger("room.log", 3)

    fmt.Printf(" Name your room: ")
    fmt.Scanf("%s", &r_name)
    fmt.Printf(" Enter capacity: ")
    fmt.Scan(&r_cap)

    message := fmt.Sprintf("Creating room with name: %s and capacity: %d", r_name, r_cap)
    logfile.Output(2, message)
    room := Room{ name: r_name,
                  capacity: r_cap}

    fmt.Println("----------------\n")
    c := Furniture{"Chair"}
    d := Furniture{"Desk"}
    room.items = []Furniture{d, c}

    room.ShowContents()
}
