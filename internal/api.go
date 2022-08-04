// Package internal contains common API functions and structures shared between lexer packages.
package internal

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/zhangdapeng520/zdpgo_pygments"
)

var (
	ignoredSuffixes = [...]string{
		// Editor backups
		"~", ".bak", ".old", ".orig",
		// Debian and derivatives apt/dpkg/ucf backups
		".dpkg-dist", ".dpkg-old", ".ucf-dist", ".ucf-new", ".ucf-old",
		// Red Hat and derivatives rpm backups
		".rpmnew", ".rpmorig", ".rpmsave",
		// Build system input/template files
		".in",
	}
)

// Registry of Lexers.
var Registry = struct {
	Lexers  zdpgo_pygments.Lexers
	byName  map[string]zdpgo_pygments.Lexer
	byAlias map[string]zdpgo_pygments.Lexer
}{
	byName:  map[string]zdpgo_pygments.Lexer{},
	byAlias: map[string]zdpgo_pygments.Lexer{},
}

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string {
	out := []string{}
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		out = append(out, config.Name)
		if withAliases {
			out = append(out, config.Aliases...)
		}
	}
	sort.Strings(out)
	return out
}

// Get a Lexer by name, alias or file extension.
func Get(name string) zdpgo_pygments.Lexer {
	if lexer := Registry.byName[name]; lexer != nil {
		return lexer
	}
	if lexer := Registry.byAlias[name]; lexer != nil {
		return lexer
	}
	if lexer := Registry.byName[strings.ToLower(name)]; lexer != nil {
		return lexer
	}
	if lexer := Registry.byAlias[strings.ToLower(name)]; lexer != nil {
		return lexer
	}

	candidates := zdpgo_pygments.PrioritisedLexers{}
	// Try file extension.
	if lexer := Match("filename." + name); lexer != nil {
		candidates = append(candidates, lexer)
	}
	// Try exact filename.
	if lexer := Match(name); lexer != nil {
		candidates = append(candidates, lexer)
	}
	if len(candidates) == 0 {
		return nil
	}
	sort.Sort(candidates)
	return candidates[0]
}

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) zdpgo_pygments.Lexer {
	matched := zdpgo_pygments.PrioritisedLexers{}
	for _, l := range Registry.Lexers {
		for _, lmt := range l.Config().MimeTypes {
			if mimeType == lmt {
				matched = append(matched, l)
			}
		}
	}
	if len(matched) != 0 {
		sort.Sort(matched)
		return matched[0]
	}
	return nil
}

// Match returns the first lexer matching filename.
func Match(filename string) zdpgo_pygments.Lexer {
	filename = filepath.Base(filename)
	matched := zdpgo_pygments.PrioritisedLexers{}
	// First, try primary filename matches.
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.Filenames {
			ok, err := filepath.Match(glob, filename)
			if err != nil { // nolint
				panic(err)
			} else if ok {
				matched = append(matched, lexer)
			} else {
				for _, suf := range &ignoredSuffixes {
					ok, err := filepath.Match(glob+suf, filename)
					if err != nil {
						panic(err)
					} else if ok {
						matched = append(matched, lexer)
						break
					}
				}
			}
		}
	}
	if len(matched) > 0 {
		sort.Sort(matched)
		return matched[0]
	}
	matched = nil
	// Next, try filename aliases.
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.AliasFilenames {
			ok, err := filepath.Match(glob, filename)
			if err != nil { // nolint
				panic(err)
			} else if ok {
				matched = append(matched, lexer)
			} else {
				for _, suf := range &ignoredSuffixes {
					ok, err := filepath.Match(glob+suf, filename)
					if err != nil {
						panic(err)
					} else if ok {
						matched = append(matched, lexer)
						break
					}
				}
			}
		}
	}
	if len(matched) > 0 {
		sort.Sort(matched)
		return matched[0]
	}
	return nil
}

// Analyse text content and return the "best" lexer..
func Analyse(text string) zdpgo_pygments.Lexer {
	var picked zdpgo_pygments.Lexer
	highest := float32(0.0)
	for _, lexer := range Registry.Lexers {
		if analyser, ok := lexer.(zdpgo_pygments.Analyser); ok {
			weight := analyser.AnalyseText(text)
			if weight > highest {
				picked = lexer
				highest = weight
			}
		}
	}
	return picked
}

// Register a Lexer with the global registry.
func Register(lexer zdpgo_pygments.Lexer) zdpgo_pygments.Lexer {
	config := lexer.Config()
	Registry.byName[config.Name] = lexer
	Registry.byName[strings.ToLower(config.Name)] = lexer
	for _, alias := range config.Aliases {
		Registry.byAlias[alias] = lexer
		Registry.byAlias[strings.ToLower(alias)] = lexer
	}
	Registry.Lexers = append(Registry.Lexers, lexer)
	return lexer
}

// PlaintextRules is used for the fallback lexer as well as the explicit
// plaintext lexer.
func PlaintextRules() zdpgo_pygments.Rules {
	return zdpgo_pygments.Rules{
		"root": []zdpgo_pygments.Rule{
			{`.+`, zdpgo_pygments.Text, nil},
			{`\n`, zdpgo_pygments.Text, nil},
		},
	}
}

// Fallback lexer if no other is found.
var Fallback zdpgo_pygments.Lexer = zdpgo_pygments.MustNewLazyLexer(&zdpgo_pygments.Config{
	Name:      "fallback",
	Filenames: []string{"*"},
}, PlaintextRules)
