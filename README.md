# go-i18n

Minimum i18n package for Go.

## Usage

```Go
import (
	"fmt"

	"github.com/kohkimakimoto/go-i18n"
)

func main() {
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

	// You can override default fallback locale
	// I18n.FallbackLocale = "ja"

	// Create translate function for 'ja' locale
	t := I18n.TFunc("ja")

	fmt.Println(t("hello world!"))
	// こんにちは世界!
	fmt.Println(t("hoge"))
	// Hoge
	fmt.Println(t("foo"))
	// foo
}
```

## Author

Kohki Makimoto <kohki.makimoto@gmail.com>

## License

The MIT License (MIT)

