package tool

import (
	"encoding/json"
	"io/ioutil"
)

func FromJson(in []byte) (x interface{}, err error) {
	if err = json.Unmarshal(in, &x); err != nil {
		return x, err
	}
	return
	//switch m.(type) {
	//case map[string]interface{}:
	//	return m.(map[string]interface{}), nil
	//}
	//return nil, fmt.Errorf("unsupported type parsing: %v", reflect.TypeOf(m))
}

//func FromYaml(in []byte) (MetaData, error) {
//	var m interface{}
//	if err := yaml.Unmarshal(in, &m); err != nil {
//		return nil, err
//	}
//	switch m.(type) {
//	case map[string]interface{}:
//		return m.(map[string]interface{}), nil
//	}
//	return nil, fmt.Errorf("unsupported type parsing: %v", reflect.TypeOf(m))
//}
//
//func FromToml(in []byte) (MetaData, error) {
//	var m interface{}
//	if err := toml.Unmarshal(in, &m); err != nil {
//		return nil, err
//	}
//	switch m.(type) {
//	case map[string]interface{}:
//		return m.(map[string]interface{}), nil
//	}
//	return nil, fmt.Errorf("unsupported type parsing: %v", reflect.TypeOf(m))
//}
//
func FromJsonFile(file string) (interface{}, error) {
	if in, err := ioutil.ReadFile(file); err != nil {
		return nil, err
	} else {
		return FromJson(in)
	}
}

//
//func FromYamlFile(file string) (MetaData, error) {
//	if in, err := ioutil.ReadFile(file); err != nil {
//		return nil, err
//	} else {
//		return FromYaml(in)
//	}
//}
//
//func FromTomlFile(file string) (MetaData, error) {
//	if in, err := ioutil.ReadFile(file); err != nil {
//		return nil, err
//	} else {
//		return FromToml(in)
//	}
//}
