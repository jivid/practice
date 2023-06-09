package main

import (
	"fmt"
	"strings"
	"bufio"
	"github.com/jivid/bittorrent/bencoding"
	"encoding/json"
)

func main() {
	input := "l5:diviji3el2:meed3:cow3:moo5:namesl5:divij7:kanchane4:datad2:idi1234eeee"
	decoded, _ := bencoding.Decode(bufio.NewReader(strings.NewReader(input)))
	b, err := json.MarshalIndent(decoded, "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
