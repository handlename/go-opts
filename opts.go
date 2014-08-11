package opts

import (
	"flag"
	"reflect"
	"strconv"
)

// Parse command line options according to definition.
func Parse(d interface{}) {
	rv := reflect.ValueOf(d).Elem()

	for i := 0; i < rv.NumField(); i++ {
		ft := rv.Type().Field(i)
		fv := rv.Field(i)

		switch x := fv.Addr().Interface().(type) {
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
		}
	}

	flag.Parse()
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

func atof(s string) (f float64) {
	f, _ = strconv.ParseFloat(s, 64)
	return f
}
