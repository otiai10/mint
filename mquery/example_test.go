package mquery

var a = map[string]any{
	"foo": "bar",
	"hoge": map[string]any{
		"name": "otiai10",
	},
	"fuga": map[int]map[string]any{
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
