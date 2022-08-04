package p

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// Pacmanconf lexer.
var Pacmanconf = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "PacmanConf",
		Aliases:   []string{"pacmanconf"},
		Filenames: []string{"pacman.conf"},
		MimeTypes: []string{},
	},
	pacmanconfRules,
))

func pacmanconfRules() Rules {
	return Rules{
		"root": {
			{`#.*$`, CommentSingle, nil},
			{`^\s*\[.*?\]\s*$`, Keyword, nil},
			{`(\w+)(\s*)(=)`, ByGroups(NameAttribute, Text, Operator), nil},
			{`^(\s*)(\w+)(\s*)$`, ByGroups(Text, NameAttribute, Text), nil},
			{Words(``, `\b`, `$repo`, `$arch`, `%o`, `%u`), NameVariable, nil},
			{`.`, Text, nil},
		},
	}
}
