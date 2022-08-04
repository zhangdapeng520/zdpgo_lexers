package b

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// Bibtex lexer.
var Bibtex = internal.Register(MustNewLazyLexer(
	&Config{
		Name:            "BibTeX",
		Aliases:         []string{"bib", "bibtex"},
		Filenames:       []string{"*.bib"},
		MimeTypes:       []string{"text/x-bibtex"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	bibtexRules,
))

func bibtexRules() Rules {
	return Rules{
		"root": {
			Include("whitespace"),
			{`@comment`, Comment, nil},
			{`@preamble`, NameClass, Push("closing-brace", "value", "opening-brace")},
			{`@string`, NameClass, Push("closing-brace", "field", "opening-brace")},
			{"@[a-z_@!$&*+\\-./:;<>?\\[\\\\\\]^`|~][\\w@!$&*+\\-./:;<>?\\[\\\\\\]^`|~]*", NameClass, Push("closing-brace", "command-body", "opening-brace")},
			{`.+`, Comment, nil},
		},
		"opening-brace": {
			Include("whitespace"),
			{`[{(]`, Punctuation, Pop(1)},
		},
		"closing-brace": {
			Include("whitespace"),
			{`[})]`, Punctuation, Pop(1)},
		},
		"command-body": {
			Include("whitespace"),
			{`[^\s\,\}]+`, NameLabel, Push("#pop", "fields")},
		},
		"fields": {
			Include("whitespace"),
			{`,`, Punctuation, Push("field")},
			Default(Pop(1)),
		},
		"field": {
			Include("whitespace"),
			{"[a-z_@!$&*+\\-./:;<>?\\[\\\\\\]^`|~][\\w@!$&*+\\-./:;<>?\\[\\\\\\]^`|~]*", NameAttribute, Push("value", "=")},
			Default(Pop(1)),
		},
		"=": {
			Include("whitespace"),
			{`=`, Punctuation, Pop(1)},
		},
		"value": {
			Include("whitespace"),
			{"[a-z_@!$&*+\\-./:;<>?\\[\\\\\\]^`|~][\\w@!$&*+\\-./:;<>?\\[\\\\\\]^`|~]*", NameVariable, nil},
			{`"`, LiteralString, Push("quoted-string")},
			{`\{`, LiteralString, Push("braced-string")},
			{`[\d]+`, LiteralNumber, nil},
			{`#`, Punctuation, nil},
			Default(Pop(1)),
		},
		"quoted-string": {
			{`\{`, LiteralString, Push("braced-string")},
			{`"`, LiteralString, Pop(1)},
			{`[^\{\"]+`, LiteralString, nil},
		},
		"braced-string": {
			{`\{`, LiteralString, Push()},
			{`\}`, LiteralString, Pop(1)},
			{`[^\{\}]+`, LiteralString, nil},
		},
		"whitespace": {
			{`\s+`, Text, nil},
		},
	}
}
