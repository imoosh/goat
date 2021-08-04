package tool

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "github.com/BurntSushi/toml"
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

func FromJson(in []byte) (x interface{}, err error) {
    err = json.Unmarshal(in, &x)
    return
}

func FromYaml(in []byte) (x interface{}, err error) {
    err = yaml.Unmarshal(in, &x)
    return
}

func FromToml(in []byte) (x interface{}, err error) {
    err = toml.Unmarshal(in, &x)
    return
}

func FromXml(in []byte) (x interface{}, err error) {
    if err = xml.Unmarshal(in, &x); err != nil {
        fmt.Println(in, err)
    }
    return
}

func FromJsonFile(file string) (interface{}, error) {
    if in, err := ioutil.ReadFile(file); err != nil {
        return nil, err
    } else {
        return FromJson(in)
    }
}

func FromYamlFile(file string) (interface{}, error) {
    if in, err := ioutil.ReadFile(file); err != nil {
        return nil, err
    } else {
        return FromYaml(in)
    }
}

func FromTomlFile(file string) (interface{}, error) {
    if in, err := ioutil.ReadFile(file); err != nil {
        return nil, err
    } else {
        return FromToml(in)
    }
}

func FromXmlFile(file string) (interface{}, error) {
    if in, err := ioutil.ReadFile(file); err != nil {
        fmt.Println(err)
        return nil, err
    } else {
        return FromXml(in)
    }
}
