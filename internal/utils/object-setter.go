package utils

import (
	"log"
	"reflect"
)

func SetField(obj interface{}, fieldName string, value interface{}) {
	// 1. Get the reflect.Value of the object and then use Elem()
	//    to get the value it points to (the actual struct).
	structValue := reflect.ValueOf(obj).Elem()

	// 2. Find the field by name.
	fieldValue := structValue.FieldByName(fieldName)

	// 3. Check if the field is valid.
	if !fieldValue.IsValid() {
		log.Printf("Field %s not found\n", fieldName)
		return
	}

	// 4. Check if the field is settable. It must be exported and the
	//    original variable must have been addressable (passed as a pointer).
	if !fieldValue.CanSet() {
		log.Printf("Field %s cannot be set (must be exported)\n", fieldName)
		return
	}

	// 5. Set the value using the appropriate type-specific setter or general Set.
	//    The provided value must be of a type assignable to the field's type.
	val := reflect.ValueOf(value)
	if fieldValue.Type() != val.Type() {
		log.Printf("Provided value type (%s) is not assignable to field type (%s)\n", val.Type(), fieldValue.Type())
		return
	}

	fieldValue.Set(val) // Use the general Set() method
}
