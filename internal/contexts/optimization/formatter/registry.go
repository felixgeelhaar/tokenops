// Registry maps command tokens to Formatters and applies the per-command
// loss policy. It is the single entry point the CLI wrapper and the proxy
// optimizer both call, so loss configuration and formatter dispatch live in
// one place.
package formatter

import "strings"

// LossPolicy carries the per-command loss configuration. Default is the
// level applied to any command without an explicit override; Overrides maps
// a command token to its level. This mirrors the "configurable per command"
// policy: operators can dial individual noisy commands more aggressively
// without loosening the global default.
type LossPolicy struct {
	Default   LossLevel
	Overrides map[string]LossLevel
}

// LevelFor resolves the loss level for a command token, honouring an
// override when present.
func (p LossPolicy) LevelFor(command string) LossLevel {
	if p.Overrides != nil {
		if lvl, ok := p.Overrides[strings.ToLower(command)]; ok {
			return lvl
		}
	}
	return p.Default
}

// Registry holds the formatter set and the loss policy.
type Registry struct {
	formatters map[string]Formatter
	policy     LossPolicy
}

// NewRegistry constructs a Registry with the given policy and formatters.
// A nil Overrides map is tolerated. Later formatters for the same command
// token replace earlier ones.
func NewRegistry(policy LossPolicy, formatters ...Formatter) *Registry {
	r := &Registry{
		formatters: make(map[string]Formatter, len(formatters)),
		policy:     policy,
	}
	for _, f := range formatters {
		r.formatters[strings.ToLower(f.Command())] = f
	}
	return r
}

// Formatter returns the formatter registered for command, or nil.
func (r *Registry) Formatter(command string) Formatter {
	return r.formatters[strings.ToLower(command)]
}

// Commands returns the registered command tokens, order unspecified.
func (r *Registry) Commands() []string {
	out := make([]string, 0, len(r.formatters))
	for c := range r.formatters {
		out = append(out, c)
	}
	return out
}

// Format looks up the formatter for the first token of argv, resolves its
// loss level from the policy, and compresses raw. When no formatter is
// registered for the command it falls back to the always-safe generic
// formatter at the policy level so unknown commands still shed pure noise.
// The returned bool reports whether a command-specific formatter handled
// the input (false = generic fallback), letting callers attribute savings.
func (r *Registry) Format(argv []string, raw []byte) (Result, bool) {
	command := firstToken(argv)
	level := r.policy.LevelFor(command)
	if f := r.formatters[strings.ToLower(command)]; f != nil {
		res, ok := f.Format(raw, level)
		if !ok {
			return rawPassthrough(raw, "declined"), true
		}
		return res, true
	}
	res, ok := generic.Format(raw, level)
	if !ok {
		return rawPassthrough(raw, "declined"), false
	}
	return res, false
}

// firstToken returns the command token argv[0] carries, ignoring a leading
// path (so "/usr/bin/git" keys as "git").
func firstToken(argv []string) string {
	if len(argv) == 0 {
		return ""
	}
	tok := argv[0]
	if i := strings.LastIndexAny(tok, "/\\"); i >= 0 {
		tok = tok[i+1:]
	}
	return tok
}

// generic is the package-level always-safe formatter used as the fallback
// for unregistered commands.
var generic = NewGeneric()
