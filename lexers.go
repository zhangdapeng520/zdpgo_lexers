package zdpgo_lexers

// nolint
import (
	_ "github.com/zhangdapeng520/zdpgo_lexers/a"
	_ "github.com/zhangdapeng520/zdpgo_lexers/b"
	_ "github.com/zhangdapeng520/zdpgo_lexers/c"
	_ "github.com/zhangdapeng520/zdpgo_lexers/circular"
	_ "github.com/zhangdapeng520/zdpgo_lexers/d"
	_ "github.com/zhangdapeng520/zdpgo_lexers/e"
	_ "github.com/zhangdapeng520/zdpgo_lexers/f"
	_ "github.com/zhangdapeng520/zdpgo_lexers/g"
	_ "github.com/zhangdapeng520/zdpgo_lexers/h"
	_ "github.com/zhangdapeng520/zdpgo_lexers/i"
	"github.com/zhangdapeng520/zdpgo_lexers/internal"
	_ "github.com/zhangdapeng520/zdpgo_lexers/j"
	_ "github.com/zhangdapeng520/zdpgo_lexers/k"
	_ "github.com/zhangdapeng520/zdpgo_lexers/l"
	_ "github.com/zhangdapeng520/zdpgo_lexers/m"
	_ "github.com/zhangdapeng520/zdpgo_lexers/n"
	_ "github.com/zhangdapeng520/zdpgo_lexers/o"
	_ "github.com/zhangdapeng520/zdpgo_lexers/p"
	_ "github.com/zhangdapeng520/zdpgo_lexers/q"
	_ "github.com/zhangdapeng520/zdpgo_lexers/r"
	_ "github.com/zhangdapeng520/zdpgo_lexers/s"
	_ "github.com/zhangdapeng520/zdpgo_lexers/t"
	_ "github.com/zhangdapeng520/zdpgo_lexers/v"
	_ "github.com/zhangdapeng520/zdpgo_lexers/w"
	_ "github.com/zhangdapeng520/zdpgo_lexers/x"
	_ "github.com/zhangdapeng520/zdpgo_lexers/y"
	_ "github.com/zhangdapeng520/zdpgo_lexers/z"
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Registry of Lexers.
var Registry = internal.Registry

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string { return internal.Names(withAliases) }

// Get a Lexer by name, alias or file extension.
func Get(name string) zdpgo_pygments.Lexer { return internal.Get(name) }

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) zdpgo_pygments.Lexer { return internal.MatchMimeType(mimeType) }

// Match returns the first lexer matching filename.
func Match(filename string) zdpgo_pygments.Lexer { return internal.Match(filename) }

// Analyse text content and return the "best" lexer..
func Analyse(text string) zdpgo_pygments.Lexer { return internal.Analyse(text) }

// Register a Lexer with the global registry.
func Register(lexer zdpgo_pygments.Lexer) zdpgo_pygments.Lexer { return internal.Register(lexer) }

// Fallback lexer if no other is found.
var Fallback = internal.Fallback
