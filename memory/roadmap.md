---
updated: 2026-07-03
---
## Now
- Validate `command_fmt` proxy-plane optimizer (toolfmt) on real agent traffic.

## Next
- Catalog fast-follow: evaluate more differentiator commands (oc, nomad, packer, vault, gem, swift, nix).
- `fmt learn` telemetry accrual → revisit auto-tuning thresholds on real usage.
- Consider a config-mutating MCP tool (add/tune formatters via MCP), extending tokenops_fmt_learn.

## Later
- Low-value RTK-parity tail (ls/cat/find/grep/diff/wget) — only if user demand; mostly signal.
- Per-subcommand JSON-aware formatters for cloud CLIs (aws/gcloud/az currently pass JSON through untouched).

## Done
- v0.26.0: fmt engine + 17 formatters + learning loop.
- v0.27.0: catalog complete — 46 commands / 51 tokens.
- v0.28.0: user-extensible (config formatters + learn --apply) + MCP tokenops_fmt_learn.
- v0.28.1: full docs for the fmt subsystem.
