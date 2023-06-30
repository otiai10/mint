package mquery

var a = map[string]interface{}{
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

func ExampleQuery() {
	// Query(m, "foo") => "bar"
	// Query(m, "hoge.name") => "otiai10"
	// Query(m, "fuga.1.greet") => "こんにちは"
	// Query(m, "langs.0") => "Go"
	// Query(m, "baz") => nil
	// Query(m, "required") => false
}
