package conf

import (
    "github.com/go-yaml/yaml"
    "go.uber.org/zap/buffer"
    "io/ioutil"
)

type Yaml struct {
    x interface{}
}

func NewYamlParser() Parser {
    return &Yaml{}
}

func (y *Yaml) Parse(f string, x interface{}) {
    file, err := ioutil.ReadFile(f)
    if err != nil {
        panic(err)
    }
    if err := yaml.Unmarshal(file, x); err != nil {
        panic(err)
    }
}

func (y *Yaml) Release() {

}

func (y *Yaml) String() string {
    bf := &buffer.Buffer{}
    if err := yaml.NewEncoder(bf).Encode(y.x); err != nil {
        panic(err)
    }
    return bf.String()
}
