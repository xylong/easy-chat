package tools

import "reflect"

// UniqueArray 去重
func UniqueArray(slice []string) []string {
	m := make(map[string]bool)
	unique := []string{}
	for _, item := range slice {
		if !m[item] {
			m[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

// M2s map转struct
// 根据struct的字段或者标签转
func M2s(m map[string]interface{}, s interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		panic("s must be a ptr")
	}

	// 从map获取值
	get := func(key, tag string) (value interface{}) {
		for k, v := range m {
			if k == key || k == tag {
				value = v
				break
			}
		}
		return value
	}

	v = v.Elem()
	length := v.NumField()
	for i := 0; i < length; i++ {
		value := get(v.Type().Field(i).Name, v.Type().Field(i).Tag.Get("json"))
		if value != nil && v.Field(i).Kind() == reflect.ValueOf(value).Kind() {
			v.Field(i).Set(reflect.ValueOf(value))
		}
	}
}
