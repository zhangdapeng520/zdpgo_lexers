package d

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

// Diff lexer.
var Diff = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Diff",
		Aliases:   []string{"diff", "udiff"},
		EnsureNL:  true,
		Filenames: []string{"*.diff", "*.patch"},
		MimeTypes: []string{"text/x-diff", "text/x-patch"},
	},
	diffRules,
))

func diffRules() Rules {
	return Rules{
		"root": {
			{` .*\n`, Text, nil},
			{`\+.*\n`, GenericInserted, nil},
			{`-.*\n`, GenericDeleted, nil},
			{`!.*\n`, GenericStrong, nil},
			{`@.*\n`, GenericSubheading, nil},
			{`([Ii]ndex|diff).*\n`, GenericHeading, nil},
			{`=.*\n`, GenericHeading, nil},
			{`.*\n`, Text, nil},
		},
	}
}
