package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email     string    `json:"e-mail" validate:"required,email"`
	Addresses []Address `validate:"required,dive,required"`
}

type Address struct {
	Street *bool
	Phone  int `validate:"required"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	address := Address{
		// Street: true,
		// Phone: 1,
	}

	user := &User{
		Email:     "smith@gmail.com",
		Addresses: []Address{address},
	}
	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(1, err.Namespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(12, err.Field())    // by passing alt name to Report1,Error like below
			fmt.Println(13, err.StructNamespace())
			fmt.Println(15, err.StructField())
			fmt.Println(17, err.Tag())
			fmt.Println(18, err.ActualTag())
			fmt.Println(19, err.Kind())
			fmt.Println(1543, err.Type())
			fmt.Println(145, err.Value())
			fmt.Println(145, err.Param())
		}

		return
	}
	fmt.Println(user)

	fmt.Println(user.Addresses[0].Street)

}
