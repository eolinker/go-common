package auto

import "reflect"

const (
	TagAutoI18n = "aoi18n"
)

func I18nConvert(rv reflect.Value, i18nMap map[string]string) {
	i18nSet(rv, i18nMap)
	return
}

func i18nSet(rv reflect.Value, i18nMap map[string]string) {
	if rv.Kind() == reflect.Interface || rv.Kind() == reflect.Ptr {
		if !rv.IsValid() {
			return
		}
		if rv.IsNil() {
			return
		}
		i18nSet(rv.Elem(), i18nMap)
		return
	}

	switch rv.Kind() {
	case reflect.Struct:

		num := rv.NumField()
		rt := rv.Type()
		for i := 0; i < num; i++ {
			fieldValue := rv.Field(i)
			fieldType := rt.Field(i)
			if fieldType.Anonymous {
				i18nSet(fieldValue, i18nMap)
				continue
			}

			_, has := fieldType.Tag.Lookup(TagAutoI18n)
			if has {
				if fieldValue.Kind() == reflect.String {
					val := fieldValue.String()
					v, ok := i18nMap[val]
					if ok {
						fieldValue.Set(reflect.ValueOf(v))
					}
					continue
				}
			}
			i18nSet(fieldValue, i18nMap)
		}
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			i18nSet(rv.Index(i), i18nMap)
		}
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			i18nSet(rv.MapIndex(key), i18nMap)
		}
	default:
		return
	}

}
