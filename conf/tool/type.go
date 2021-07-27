package tool

import (
    "fmt"
    "reflect"
)

type MetaData map[string]interface{}

type item struct {
    level int
    rtype reflect.Type
    name  string
}

type structInfo struct {
    name    string
    typeof  reflect.Type
    members []structInfo
}

func (info structInfo) String() string {
    return fmt.Sprint(info.name, "-", info.typeof, "-", info.members)
}

func (info structInfo) print(level int) {
    var desc string
    for i := 0; i < level; i++ {
        desc += "    "
    }
    switch info.typeof {
    case reflect.TypeOf(""):
        fmt.Printf("%s%s string\n", desc, info.name)
        return
    case reflect.TypeOf(float64(0)):
        fmt.Printf("%s%s int\n", desc, info.name)
        return
    case reflect.TypeOf(map[string]interface{}{}):
        fmt.Printf("%stype %s struct {\n", desc, info.name)
        for _, m := range info.members {
            m.print(level + 1)
            continue
        }
    case reflect.TypeOf([]interface{}{}):
        fmt.Printf("%s%s []interface\n", desc, info.name)
        //if len()
    default:
        panic(fmt.Errorf("unsupported type: %v", info.typeof))
    }
    desc += "}"
    fmt.Println(desc)
}

func (info structInfo) Print() {
    info.print(0)
}

// name: 字段名，typeof：字段类型
func (m MetaData) travel(name string, typeof reflect.Type) (info structInfo) {
    info.name = name
    info.typeof = typeof
    for k, v := range m {
        switch reflect.TypeOf(v) {
        case reflect.TypeOf(map[string]interface{}{}):
            tmp := MetaData(v.(map[string]interface{})).travel(k, reflect.TypeOf(map[string]interface{}{}))
            info.members = append(info.members, tmp)
            continue
        case reflect.TypeOf([]interface{}{}):
            if len(v.([]interface{})) != 0 {
                tmp := MetaData(v.([]interface{})[0].(map[string]interface{})).travel(k, reflect.TypeOf(v.([]interface{})[0]))
                info.members = append(info.members, tmp)
            }
            continue
        }
        info.members = append(info.members, structInfo{
            name:   k,
            typeof: reflect.TypeOf(v),
        })
    }
    return
}

func (m MetaData) Travel() {
    info := m.travel("Struct", reflect.TypeOf(map[string]interface{}{}))
    fmt.Println(info)

    info.Print()
}
