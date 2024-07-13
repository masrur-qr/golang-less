package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type user struct {
	Name    string
	Surname string
	Age     int
}
type Login struct {
	Login string
	Password string
} 

func main() {
	var ClientData = user{
		Name:    "Dean",
		Surname: "Winchester",
	}

	isEmpty := CheckField(ClientData)
	fmt.Printf("isEmpty: %v\n", isEmpty)
	// ?============================
	var clData = Login{
		Login: "Hello",
		Password: "asd",
	}
	isEmptyTwo := CheckField(clData)
	fmt.Printf("isEmptyTwo: %v\n", isEmptyTwo)

}

func CheckField(Object any) bool {
	var isEmptyField bool = false // false mean empty
	element := reflect.ValueOf(Object).NumField()
	for i := 0; i < element; i++ {

		FieldType := reflect.TypeOf(Object).Field(i).Type
		if fmt.Sprint(FieldType) == "string" {
			FieldValue := reflect.ValueOf(Object).Field(i)
			if fmt.Sprint(FieldValue) == "" {
				// fmt.Println("error")
				isEmptyField = false
			} else {
				// fmt.Println("Succes")
				isEmptyField = true
			}
		} else {
			FieldValue := reflect.ValueOf(Object).Field(i)
			FieldValueInt, _ := strconv.Atoi(fmt.Sprint(FieldValue))
			if FieldValueInt <= 0 {
				// fmt.Println("error")
				isEmptyField = false
			} else {
				// fmt.Println("Succes")
				isEmptyField = true
			}
		}

	}
	return isEmptyField
}

// ? reflect.ValueOf(Object).NumField() -- gives us numbers of field in struct
// ? reflect.ValueOf(Object).Field(i) -- gives us field value
// ? reflect.TypeOf(Object).Field(i).type -- gives us field type
