package p

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// Pkgconfig lexer.
var Pkgconfig = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "PkgConfig",
		Aliases:   []string{"pkgconfig"},
		Filenames: []string{"*.pc"},
		MimeTypes: []string{},
	},
	pkgconfigRules,
))

func pkgconfigRules() Rules {
	return Rules{
		"root": {
			{`#.*$`, CommentSingle, nil},
			{`^(\w+)(=)`, ByGroups(NameAttribute, Operator), nil},
			{`^([\w.]+)(:)`, ByGroups(NameTag, Punctuation), Push("spvalue")},
			Include("interp"),
			{`[^${}#=:\n.]+`, Text, nil},
			{`.`, Text, nil},
		},
		"interp": {
			{`\$\$`, Text, nil},
			{`\$\{`, LiteralStringInterpol, Push("curly")},
		},
		"curly": {
			{`\}`, LiteralStringInterpol, Pop(1)},
			{`\w+`, NameAttribute, nil},
		},
		"spvalue": {
			Include("interp"),
			{`#.*$`, CommentSingle, Pop(1)},
			{`\n`, Text, Pop(1)},
			{`[^${}#\n]+`, Text, nil},
			{`.`, Text, nil},
		},
	}
}
