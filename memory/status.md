---
updated: 2026-07-03
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics) that ships `tokenops fmt` — a deterministic command-output compression subsystem now at **51 formatters / 57 command tokens**, user-extensible via config, self-tuning learn loop, MCP-exposed. Latest release **v0.29.0** (51 formatters). Agent OS memory committed. Proxy plane validated via default-pipeline integration test.

## Last Session Summary
Built `tokenops fmt` (5 releases v0.26.0→v0.28.1, 4 PRs). Then: validated the proxy plane (default-pipeline test), committed Agent OS memory, and did the catalog fast-follow — +oc/nomad/packer/gem/swift/nix (51 formatters). vault deferred (secret-bearing output).

## Next Session Should
Check `tokenops read-guard stats` — if the observe-mode reclaimable-token numbers look good over a few real sessions, flip the hook to --mode active in ~/.claude/settings.json to actually reclaim. Otherwise continue on demand.


## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
