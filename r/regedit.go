package r

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// Reg lexer.
var Reg = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "reg",
		Aliases:   []string{"registry"},
		Filenames: []string{"*.reg"},
		MimeTypes: []string{"text/x-windows-registry"},
	},
	regRules,
))

func regRules() Rules {
	return Rules{
		"root": {
			{`Windows Registry Editor.*`, Text, nil},
			{`\s+`, Text, nil},
			{`[;#].*`, CommentSingle, nil},
			{`(\[)(-?)(HKEY_[A-Z_]+)(.*?\])$`, ByGroups(Keyword, Operator, NameBuiltin, Keyword), nil},
			{`("(?:\\"|\\\\|[^"])+")([ \t]*)(=)([ \t]*)`, ByGroups(NameAttribute, Text, Operator, Text), Push("value")},
			{`(.*?)([ \t]*)(=)([ \t]*)`, ByGroups(NameAttribute, Text, Operator, Text), Push("value")},
		},
		"value": {
			{`-`, Operator, Pop(1)},
			{`(dword|hex(?:\([0-9a-fA-F]\))?)(:)([0-9a-fA-F,]+)`, ByGroups(NameVariable, Punctuation, LiteralNumber), Pop(1)},
			{`.+`, LiteralString, Pop(1)},
			Default(Pop(1)),
		},
	}
}
