package c

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// Cmake lexer.
var Cmake = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "CMake",
		Aliases:   []string{"cmake"},
		Filenames: []string{"*.cmake", "CMakeLists.txt"},
		MimeTypes: []string{"text/x-cmake"},
	},
	cmakeRules,
))

func cmakeRules() Rules {
	return Rules{
		"root": {
			{`\b(\w+)([ \t]*)(\()`, ByGroups(NameBuiltin, Text, Punctuation), Push("args")},
			Include("keywords"),
			Include("ws"),
		},
		"args": {
			{`\(`, Punctuation, Push()},
			{`\)`, Punctuation, Pop(1)},
			{`(\$\{)(.+?)(\})`, ByGroups(Operator, NameVariable, Operator), nil},
			{`(\$ENV\{)(.+?)(\})`, ByGroups(Operator, NameVariable, Operator), nil},
			{`(\$<)(.+?)(>)`, ByGroups(Operator, NameVariable, Operator), nil},
			{`(?s)".*?"`, LiteralStringDouble, nil},
			{`\\\S+`, LiteralString, nil},
			{`[^)$"# \t\n]+`, LiteralString, nil},
			{`\n`, Text, nil},
			Include("keywords"),
			Include("ws"),
		},
		"string": {},
		"keywords": {
			{`\b(WIN32|UNIX|APPLE|CYGWIN|BORLAND|MINGW|MSVC|MSVC_IDE|MSVC60|MSVC70|MSVC71|MSVC80|MSVC90)\b`, Keyword, nil},
		},
		"ws": {
			{`[ \t]+`, Text, nil},
			{`#.*\n`, Comment, nil},
		},
	}
}
