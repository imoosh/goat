package conf

import (
    "encoding/json"
    "github.com/BurntSushi/toml"
    "go.uber.org/zap/buffer"
    "io/ioutil"
    "runtime"
)

type Json struct {
    x interface{}
}

func NewJsonParser() Parser {
    t := &Toml{}
    runtime.SetFinalizer(t, t.Release)
    return t
}

func (t *Json) Parse(f string, x interface{}) {
    file, err := ioutil.ReadFile(f)
    if err != nil {
        panic(err)
    }
    if err := json.Unmarshal(file, x); err != nil {
        panic(err)
    }
}

func (t *Json) Release() {

}

func (t *Json) String() string {
    bf := &buffer.Buffer{}
    if err := toml.NewEncoder(bf).Encode(t.x); err != nil {
        panic(err)
    }
    return bf.String()
}

