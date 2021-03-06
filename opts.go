package opts

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

// Parse command line options according to definition.
func Parse(d interface{}) (err error) {
	rv := reflect.ValueOf(d).Elem()

	for i := 0; i < rv.NumField(); i++ {
		ft := rv.Type().Field(i)
		fv := rv.Field(i)

		switch x := fv.Addr().Interface().(type) {
		default:
			return fmt.Errorf("invalid type: %s", ft.Name)
		case *string:
			flag.StringVar(
				x,
				ft.Tag.Get("flag"),
				ft.Tag.Get("default"),
				ft.Tag.Get("description"))
		case *int:
			flag.IntVar(
				x,
				ft.Tag.Get("flag"),
				atoi(ft.Tag.Get("default")),
				ft.Tag.Get("description"))
		case *float64:
			flag.Float64Var(
				x,
				ft.Tag.Get("flag"),
				atof(ft.Tag.Get("default")),
				ft.Tag.Get("description"))
		case *bool:
			flag.BoolVar(
				x,
				ft.Tag.Get("flag"),
				atob(ft.Tag.Get("default")),
				ft.Tag.Get("description"))
		}
	}

	flag.Parse()

	return nil
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

func atof(s string) (f float64) {
	f, _ = strconv.ParseFloat(s, 64)
	return f
}

func atob(s string) (b bool) {
	switch s {
	case "true":
		b = true
	case "false":
		b = false
	}

	return b
}
