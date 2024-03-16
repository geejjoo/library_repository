package dao

import "reflect"

func GetStructInfo(u interface{}) StructInfo {
	val := reflect.ValueOf(u).Elem()
	var structFields []reflect.StructField

	for i := 0; i < val.NumField(); i++ {
		structFields = append(structFields, val.Type().Field(i))
	}
	var out StructInfo
	for _, field := range structFields {
		valueField := val.FieldByName(field.Name)
		out.Fields = append(out.Fields, field.Tag.Get("db"))
		out.FieldsTypes = append(out.FieldsTypes, field.Tag.Get("db_type"))
		out.FieldsDefault = append(out.FieldsDefault, field.Tag.Get("db_default"))
		out.Pointers = append(out.Pointers, valueField.Addr().Interface())
	}
	return out
}

type StructInfo struct {
	Fields        []string
	FieldsTypes   []string
	FieldsDefault []string
	Pointers      []interface{}
}

type Condition struct {
	Equal       map[string]interface{}
	NotEqual    map[string]interface{}
	Order       []*Order
	LimitOffset *LimitOffset
	ForUpdate   bool
	Upsert      bool
}

type Order struct {
	Field string
	Asc   bool
}

type LimitOffset struct {
	Offset int64
	Limit  int64
}
