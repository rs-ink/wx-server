package util

import (
	"reflect"
)

type StructInfo struct {
	Name  string
	Tag   reflect.StructTag
	Filed reflect.StructField
}

func Info(ty reflect.Type) (result []StructInfo) {
	for i := 0; i < ty.NumField(); i++ {
		f := ty.Field(i)
		if f.Anonymous {
			result = append(result, Info(f.Type)...)
		} else {
			result = append(result, StructInfo{
				Name: f.Name,
				Tag:  f.Tag,
			})
		}
	}
	return
}
