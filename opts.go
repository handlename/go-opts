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

		switch fv.Kind() {
		case reflect.String:
			flag.StringVar(
				fv.Addr().Interface().(*string),
				ft.Tag.Get("flag"),
				ft.Tag.Get("default"),
				ft.Tag.Get("description"))
		case reflect.Int:
			flag.IntVar(
				fv.Addr().Interface().(*int),
				ft.Tag.Get("flag"),
				atoi(ft.Tag.Get("default")),
				ft.Tag.Get("description"))
		case reflect.Float64:
			flag.Float64Var(
				fv.Addr().Interface().(*float64),
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
