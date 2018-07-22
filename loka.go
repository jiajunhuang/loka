package loka

import (
	"log"
	"os"
	"reflect"
	"strconv"
)

func LoadFromEnv(c interface{}) {
	v := reflect.ValueOf(c).Elem()
	t := reflect.TypeOf(c).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		env := os.Getenv(t.Field(i).Name)
		if field.CanSet() && env != "" {
			switch field.Kind() {
			case reflect.Bool:
				b := false
				if env == "true" || env == "True" {
					b = true
				}
				field.SetBool(b)
			case reflect.String:
				field.SetString(env)
			case reflect.Int:
				i, err := strconv.Atoi(env)
				if err != nil {
					log.Printf("failed to set %s to field %s", env, field)
				}
				field.SetInt(int64(i))
			}
		}
	}
}
