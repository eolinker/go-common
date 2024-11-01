package pm3

import "github.com/eolinker/eosc"

var (
	i18nData    = eosc.BuildUntyped[string, map[string]string]()
	defaultI18n = "zh-CN"
)

func I18nRegister(language string, data map[string]string) {
	i18nData.Set(language, data)
}

func I18nGet(language string) map[string]string {
	result, has := i18nData.Get(language)
	if has {
		return result
	}

	result, has = i18nData.Get(defaultI18n)
	if has {
		return result
	}
	return make(map[string]string)
}
