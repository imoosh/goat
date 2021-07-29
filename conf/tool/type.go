package tool

import (
	"fmt"
	"github.com/fatih/structs"
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
	return fmt.Sprint(info.name, " -- ", info.typeof, " -- ", info.members, "\n")
}

func (info structInfo) print(level int) {
	var desc string
	for i := 0; i < level; i++ {
		desc += "    "
	}
	switch info.typeof.Kind() {
	case reflect.String:
		fmt.Printf("%s%s string\n", desc, info.name)
		return
	case reflect.Float64:
		fmt.Printf("%s%s int\n", desc, info.name)
		return
	case reflect.Map:
		fmt.Printf("%stype %s struct {\n", desc, info.name)
		for _, m := range info.members {
			m.print(level + 1)
			continue
		}
	case reflect.Struct:
		fmt.Printf("%stype %s struct {\n", desc, info.name)
		for _, m := range info.members {
			m.print(level + 1)
			continue
		}
	case reflect.Slice:
		fmt.Printf("%s%s []interface{}\n", desc, info.name)
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
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.Map:
			md := MetaData(v.(map[string]interface{}))
			//tmp := md.travel(k, reflect.TypeOf(map[string]interface{}{}))
			tmp := md.travel(k, reflect.TypeOf(struct{}{}))
			info.members = append(info.members, tmp)
		case reflect.Slice:
			//info.members = append(info.members, structInfo{name: k, typeof: })
			if len(v.([]interface{})) != 0 {
				md := MetaData(v.([]interface{})[0].(map[string]interface{}))
				tmp := md.travel(k, reflect.SliceOf(reflect.TypeOf(v.([]interface{})[0])))
				info.members = append(info.members, tmp)
			} else {
				// the basic type of slice is unknown
				info.members = append(info.members, structInfo{name: k, typeof: reflect.TypeOf([]interface{}{})})
			}
		default:
			info.members = append(info.members, structInfo{name: k, typeof: reflect.TypeOf(v)})
		}
	}
	structs.New()
	return
}

func (m MetaData) Travel() {
	//info := m.travel("", reflect.TypeOf(map[string]interface{}{}))
	info := m.travel("Struct", reflect.TypeOf(struct{}{}))
	fmt.Println(info)

	info.Print()
}
