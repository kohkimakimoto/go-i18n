package i18n

import "fmt"

type I18n struct {
	Translations   Translations
	FallbackLocale string
}

type Translations map[string]map[string]string

type TranslateFunc func(t string, a ...interface{}) string

var DefaultFallbacksLocale = "en"

func (i18n *I18n) T(locale string, t string, a ...interface{}) string {
	var fallbackLocal = i18n.FallbackLocale
	if fallbackLocal == "" {
		fallbackLocal = DefaultFallbacksLocale
	}

	fallback, ok := i18n.Translations[fallbackLocal]
	if !ok {
		fallback = map[string]string{}
	}

	l, ok := i18n.Translations[locale]
	if !ok {
		l = fallback
	}

	txt, ok := l[t]
	if !ok {
		txt, ok := fallback[t]
		if !ok {
			return t
		}
		return txt
	}

	if len(a) > 0 {
		txt = fmt.Sprintf(txt, a...)
	}

	return txt
}

func (i18n *I18n) TFunc(locale string) TranslateFunc {
	return func(t string, a ...interface{}) string {
		return i18n.T(locale, t, a...)
	}
}
