package i18n

import (
	"testing"
)

func TestI18n_T(t *testing.T) {
	i18n := &I18n{Translations: Translations{
		"en": {
			"hello": "hello!",
			"world": "world!",
		},
		"ja": {
			"hello": "こんにちは!",
		},
	}}

	if ret := i18n.T("en", "hello"); ret != "hello!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("ja", "hello"); ret != "こんにちは!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("en", "world"); ret != "world!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("ja", "world"); ret != "world!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("ja", "foo"); ret != "foo" {
		t.Errorf("unexpected result %s", ret)
	}

	i18n = &I18n{
		Translations: Translations{
			"en": {
				"hello": "hello!",
			},
			"ja": {
				"hello": "こんにちは!",
				"world": "世界!",
			},
		},
		FallbackLocale: "ja",
	}

	if ret := i18n.T("en", "hello"); ret != "hello!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("ja", "hello"); ret != "こんにちは!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("en", "world"); ret != "世界!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("ja", "world"); ret != "世界!" {
		t.Errorf("unexpected result %s", ret)
	}

	if ret := i18n.T("ja", "foo"); ret != "foo" {
		t.Errorf("unexpected result %s", ret)
	}
}

func TestI18n_TranslateFunc(t *testing.T) {
	i18n := &I18n{
		Translations: Translations{
			"en": {
				"hello": "hello!",
				"world": "world!",
			},
			"ja": {
				"hello": "こんにちは!",
			},
		},
	}

	T := i18n.TFunc("ja")

	if ret := T("hello"); ret != "こんにちは!" {
		t.Errorf("unexpected result %s", ret)
	}
}
