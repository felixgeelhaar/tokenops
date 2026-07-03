---
updated: 2026-07-03
---
## Now
- Let real usage accrue in ~/.tokenops/recovery/index.jsonl, then tune fmt learn thresholds (currently data-blocked).

## Next
- Consider a config-mutating MCP tool (add/tune formatters via MCP), extending tokenops_fmt_learn.
- Optional: per-subcommand JSON-aware cloud formatters (aws/gcloud/az currently pass JSON through untouched).

## Later
- Low-value RTK-parity tail (ls/cat/find/grep/diff/wget) — only if user demand; mostly signal.
- Per-subcommand JSON-aware formatters for cloud CLIs (aws/gcloud/az currently pass JSON through untouched).

## Done
- v0.26.0: fmt engine + 17 formatters + learning loop.
- v0.27.0: catalog complete — 46 commands / 51 tokens.
- v0.28.0: user-extensible (config formatters + learn --apply) + MCP tokenops_fmt_learn.
- v0.28.1: full docs for the fmt subsystem.
- 2026-07-03: proxy-plane validated via default-pipeline integration test; Agent OS memory committed.
- 2026-07-03: catalog fast-follow — +oc (kubectl alias), nomad, packer, gem, swift, nix → 51 formatters / 57 tokens. vault deferred (secret-bearing output, low compression value).
