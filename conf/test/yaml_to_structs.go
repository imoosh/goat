package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

func main() {
    file, err := ioutil.ReadFile("./config.yaml")
    if err != nil {
        panic(err)
    }

    var x interface{}
    if err := yaml.Unmarshal(file, &x); err != nil {
        panic(err)
    }
    fmt.Println(x)
}
