---
updated: 2026-07-03
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics) that now also ships `tokenops fmt` — a deterministic command-output compression subsystem (46 built-in formatters, user-extensible via config, self-tuning learn loop, MCP-exposed). Latest release **v0.28.1** is live on homebrew. Main is clean, 0 open PRs, all CI green. Agent OS memory system scaffolded this session (untracked, pending commit decision).

## Last Session Summary
Built `tokenops fmt` from scratch to match/beat rtk-ai/rtk: engine + critical-line-survival invariant, 46 formatters, 3 planes (CLI/hook/proxy), user config formatters, `fmt learn --apply` local self-tuning, MCP `tokenops_fmt_learn`. Shipped 5 releases (v0.26.0→v0.28.1), 4 PRs (#112–#115), full docs. Then scaffolded Agent OS memory.

## Next Session Should
Catalog fast-follow decision: add more differentiator commands (oc/nomad/packer/vault/gem/swift/nix) or stop at 46. Optionally revisit fmt learn thresholds once real telemetry has accrued in ~/.tokenops/recovery/index.jsonl.

## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
