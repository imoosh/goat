package conf

import (
    "github.com/BurntSushi/toml"
    "go.uber.org/zap/buffer"
    "runtime"
)

type Toml struct {
    x interface{}
}

func NewTomlParser() Parser {
    t := &Toml{}
    runtime.SetFinalizer(t, t.Release)
    return t
}

func (t *Toml) Parse(f string, x interface{}) {
    _, err := toml.DecodeFile(f, x)
    if err != nil {
        panic("toml: DecodeFile error: " + err.Error())
    }
    t.x = x
}

func (t *Toml) Release() {

}

func (t *Toml) String() string {
    bf := &buffer.Buffer{}
    if err := toml.NewEncoder(bf).Encode(t.x); err != nil {
        panic(err)
    }
    return bf.String()
}
