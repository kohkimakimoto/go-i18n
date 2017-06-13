package example

import (
	"fmt"

	"github.com/kohkimakimoto/go-i18n"
)

func Example1() {
	I18n := &i18n.I18n{
		Translations: i18n.Translations{
			// 'en' is a default fallback locale.
			"en": {
				"hello world!": "hello world!",
				"hoge":         "Hoge",
			},
			"ja": {
				"hello world!": "こんにちは世界!",
			},
		},
	}

	// you can override default fallback locale
	// I18n.FallbackLocale = "ja"

	// create translate function for 'ja' locale
	t := I18n.TFunc("ja")

	fmt.Println(t("hello world!"))
	fmt.Println(t("hoge"))
	fmt.Println(t("foo"))

	// Output: こんにちは世界!
	// Hoge
	// foo
}

func Example2() {
	I18n := &i18n.I18n{
		Translations: i18n.Translations{
			"en": {
				"My name is %s": "My name is %s",
			},
			"ja": {
				"My name is %s": "私の名前は%sです",
			},
		},
	}

	t := I18n.TFunc("ja")

	fmt.Print(t("My name is %s", "kohkimakimoto"))
	// Output: 私の名前はkohkimakimotoです
}
