package mquery

import (
	"testing"
)

var m = map[string]interface{}{
	"foo": "bar",
	"hoge": map[string]interface{}{
		"name": "otiai10",
	},
	"fuga": map[int]map[string]interface{}{
		0: {"greet": "Hello"},
		1: {"greet": "こんにちは"},
	},
	"langs":    []string{"Go", "JavaScript", "English"},
	"baz":      nil,
	"required": false,
}

func TestQuery(t *testing.T) {
	// Expect(t, Query(m, "foo")).ToBe("bar")
	// Expect(t, Query(m, "hoge.name")).ToBe("otiai10")
	// Expect(t, Query(m, "fuga.1.greet")).ToBe("こんにちは")
	// Expect(t, Query(m, "langs.0")).ToBe("Go")
}
