package s

import (
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
)

var SYSTEMD = internal.Register(MustNewLazyLexer(
	&Config{
		Name:    "SYSTEMD",
		Aliases: []string{"systemd"},
		// Suspects: man systemd.index | grep -E 'systemd\..*configuration'
		Filenames: []string{"*.automount", "*.device", "*.dnssd", "*.link", "*.mount", "*.netdev", "*.network", "*.path", "*.scope", "*.service", "*.slice", "*.socket", "*.swap", "*.target", "*.timer"},
		MimeTypes: []string{"text/plain"},
	},
	systemdRules,
))

func systemdRules() Rules {
	return Rules{
		"root": {
			{`\s+`, Text, nil},
			{`[;#].*`, Comment, nil},
			{`\[.*?\]$`, Keyword, nil},
			{`(.*?)(=)(.*)(\\\n)`, ByGroups(NameAttribute, Operator, LiteralString, Text), Push("continuation")},
			{`(.*?)(=)(.*)`, ByGroups(NameAttribute, Operator, LiteralString), nil},
		},
		"continuation": {
			{`(.*?)(\\\n)`, ByGroups(LiteralString, Text), nil},
			{`(.*)`, LiteralString, Pop(1)},
		},
	}
}
