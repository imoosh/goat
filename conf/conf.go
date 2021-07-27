package conf

import "path/filepath"

type Parser interface {
    Parse(string, interface{})
    String() string
    Release()
}

const (
    yamlFileParser = iota
    tomlFileParser
    jsonFileParser
)

var extMapping = map[string]int{
    ".yml":  yamlFileParser,
    ".yaml": yamlFileParser,
    ".toml": tomlFileParser,
    ".json": jsonFileParser,
}

type Decoder struct {
    v interface{}
    p Parser
}

// Parse 解析配置文件
func (d *Decoder) Parse(f string, v interface{}) {
    ext := filepath.Ext(f)
    t, ok := extMapping[ext]
    if !ok {
        panic("unsupported format file parsing: " + filepath.Base(f))
    }

    switch t {
    case yamlFileParser:
        d.p = NewYamlParser()
    case tomlFileParser:
        d.p = NewTomlParser()
    case jsonFileParser:
        d.p = NewJsonParser()
    default:
        panic("unsupported format file parsing: " + filepath.Base(f))
    }
    d.p.Parse(f, v)
    d.v = v
}

func (d *Decoder) String() string {
    return d.p.String()
}
