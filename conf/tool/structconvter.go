package tool

import (
	"fmt"
	"reflect"
)

type item struct {
	level int
	rtype reflect.Kind
	name  string
}

type structInfo struct {
	name    string
	tag     string
	typ     reflect.Type
	members []structInfo
}

func (info *structInfo) tagString() string {
	return " `json:\"" + info.tag + "\"`"
}

func (info *structInfo) String() string {
	return fmt.Sprint(info.name, " -- ", info.typ, " -- ", info.members, "\n")
}

func (info *structInfo) print(level int) {
	var indent string
	for i := 0; i < level; i++ {
		indent += "    "
	}
	switch info.typ.Kind() {
	case reflect.Bool:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uintptr:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		fallthrough
	case reflect.String:
		fmt.Printf("%s%s %s%s\n", indent, info.name, info.typ.Name(), info.tagString())
		return
	case reflect.Struct:
		fmt.Printf("%s%s struct {\n", indent, info.name)
		for _, m := range info.members {
			m.print(level + 1)
			continue
		}
		//indent += "}" + info.tagString()
		fmt.Println(indent + "}" + info.tagString())
	case reflect.Slice:
		if info.typ.Elem().Kind() == reflect.Struct {
			fmt.Printf("%s%s []struct {\n", indent, info.name)
			for _, m := range info.members {
				m.print(level + 1)
				continue
			}
			fmt.Println(indent + "}" + info.tagString())
		} else {
			fmt.Printf("%s%s %s%s\n", indent, info.name, info.typ.String(), info.tagString())
		}
		return
	default:
		panic(fmt.Errorf("unsupported type: %v", info.typ))
	}
}

func (info *structInfo) Print() {
	info.print(0)
}

func (info *structInfo) Fix() {

}

// name: 字段名，typ：字段类型
func travel(name string, m interface{}) (info structInfo) {
	info.tag = name
	info.name = name
	info.typ = reflect.TypeOf(m)

	switch info.typ.Kind() {
	case reflect.String:
	case reflect.Float64:
	case reflect.Map:
		info.typ = reflect.TypeOf(struct{}{})
		for k, v := range m.(map[string]interface{}) {
			info.members = append(info.members, travel(k, v))
		}
	case reflect.Slice:
		if len(m.([]interface{})) == 0 {
			info.typ = reflect.TypeOf([]struct{}{})
		} else {
			elemInfo := travel("", m.([]interface{})[0])
			info.members = elemInfo.members
			info.typ = reflect.SliceOf(elemInfo.typ)
		}
	}

	return info
}

func Travel(x interface{}) {
	info := travel("type Struct", x)
	info.Print()
}
