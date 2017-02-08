package db

import (
	"errors"
	"reflect"
)

// Document represents an element that can be put or get from a collection
type Document map[string]interface{}

// Marshal takes an interface and returns an Document compatible with tidedot
func Marshal(o interface{}) map[string]interface{} {
	oType := reflect.TypeOf(o)
	// Get the number of fields in the interface
	numField := oType.NumField()
	result := make(map[string]interface{}, numField)
	for i := 0; i < numField; i++ {
		f := oType.Field(i)
		// No need to save a private field as we won't be able to restore it
		if f.PkgPath == "" {
			result[f.Name] = reflect.ValueOf(o).Field(i).Interface()
		}
	}
	return result
}

// Unmarshal parses i and sets it if possible into o
// returns an error if not possible or if interface is not a struct
func Unmarshal(i map[string]interface{}, o interface{}) error {
	oType := reflect.TypeOf(o)
	if oType.Kind() != reflect.Ptr {
		return errors.New("Expected pointer")
	}
	//temp := reflect.New(oType)
	oVal := reflect.ValueOf(o).Elem()
	if oVal.Kind() != reflect.Struct {
		return errors.New("Interface is " + oType.Kind().String() + ", not struct")
	}
	for k, v := range i {
		f := oVal.FieldByName(k)
		if !f.IsValid() {
			return errors.New("Field \"" + k + "\" is not valid")
		}
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if !f.CanSet() {
			return errors.New("Cannot set field " + k)
		}
		if f.Kind() == reflect.ValueOf(v).Kind() {
			f.Set(reflect.ValueOf(v))
		} else {
			return errors.New("Cannot set " + reflect.ValueOf(v).Kind().String() + " into " + f.Kind().String())
		}
	}
	return nil
}
