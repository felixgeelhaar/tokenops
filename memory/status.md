---
updated: 2026-07-03
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics) that ships `tokenops fmt` — a deterministic command-output compression subsystem now at **51 formatters / 57 command tokens**, user-extensible via config, self-tuning learn loop, MCP-exposed. Latest release **v0.30.0** — adds the self-wiring JSONL analyzer (`fmt analyze`/`learn` mine Claude Code logs), a Read-side trimming diagnostic, and `read-guard` (a Claude Code PreToolUse hook that reclaims redundant re-reads). read-guard installed in OBSERVE mode on this machine.

## Last Session Summary
Removed caveman (proved <1% of tokens on real data). Built the self-wiring JSONL analyzer (fmt analyze/learn), the Read-side diagnostic, and read-guard (Claude Code PreToolUse hook). Released v0.30.0 + v0.30.1; installed read-guard in observe mode. Honest finding: for this user, re-read waste is ~0 reclaimable (already reads via ranges) — the analyzers measured it rather than assuming.


## Next Session Should
Run `tokenops read-guard stats` — if reclaimable is still ~0 across more sessions, leave observe (or remove the hook); do NOT flip to active (nothing to reclaim). Otherwise go active. Then the fmt subsystem work is effectively complete; pick a new direction.


## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
