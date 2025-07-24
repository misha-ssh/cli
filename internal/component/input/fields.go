package input

import "reflect"

const (
	fieldNumberAlias = iota
	fieldNumberLogin
	fieldNumberPassword
	fieldNumberPort
	fieldNumberPrivateKey
)

type Fields struct {
	Alias      string
	Login      string
	Password   string
	Port       int
	PrivateKey string
}

func (f *Fields) count() int {
	fieldsType := reflect.TypeOf(*f)
	return fieldsType.NumField()
}

func (f *Fields) getNameByNumber(number int) string {
	val := reflect.ValueOf(f)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	if number > typ.NumField() {
		return ""
	}

	return typ.Field(number).Name
}
