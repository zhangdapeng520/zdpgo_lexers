package i

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// Ini lexer.
var Ini = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "INI",
		Aliases:   []string{"ini", "cfg", "dosini"},
		Filenames: []string{"*.ini", "*.cfg", "*.inf", ".gitconfig", ".editorconfig"},
		MimeTypes: []string{"text/x-ini", "text/inf"},
	},
	iniRules,
))

func iniRules() Rules {
	return Rules{
		"root": {
			{`\s+`, Text, nil},
			{`[;#].*`, CommentSingle, nil},
			{`\[.*?\]$`, Keyword, nil},
			{`(.*?)([ \t]*)(=)([ \t]*)(.*(?:\n[ \t].+)*)`, ByGroups(NameAttribute, Text, Operator, Text, LiteralString), nil},
			{`(.+?)$`, NameAttribute, nil},
		},
	}
}
