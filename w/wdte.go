package w

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// WDTE lexer.
var WDTE = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "WDTE",
		Filenames: []string{"*.wdte"},
	},
	wdteRules,
))

func wdteRules() Rules {
	return Rules{
		"root": {
			{`\n`, Text, nil},
			{`\s+`, Text, nil},
			{`\\\n`, Text, nil},
			{`#(.*?)\n`, CommentSingle, nil},
			{`-?[0-9]+`, LiteralNumberInteger, nil},
			{`-?[0-9]*\.[0-9]+`, LiteralNumberFloat, nil},
			{`"[^"]*"`, LiteralString, nil},
			{`'[^']*'`, LiteralString, nil},
			{Words(``, `\b`, `switch`, `default`, `memo`), KeywordReserved, nil},
			{`{|}|;|->|=>|\(|\)|\[|\]|\.`, Operator, nil},
			{`[^{};()[\].\s]+`, NameVariable, nil},
		},
	}
}
