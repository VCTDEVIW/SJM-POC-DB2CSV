package project

import (
    "encoding/json"
    . "fmt"
    _ "log"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    person := Person{Name: "Alice", Age: 30}

    // Convert struct to JSON
    jsonData, err := json.Marshal(person)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonData)) // Output: {"name":"Alice","age":30}
}

