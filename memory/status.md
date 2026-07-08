---
updated: 2026-07-08
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics). Repo `github.com/klarlabs-studio/tokenops`, module `go.klarlabs.de/tokenops` (vanity); brew tap `felixgeelhaar/tap/tokenops`. Latest release **v0.39.0** (2026-07-07); `main` is clean. **Pricing is now researched + effective-dated (ADR 0002)** — pluggable LiteLLM source, timestamped/sourced snapshots (`tokenops pricing refresh/show/diff/lint` + a consistency guard), and a time-aware cost engine (`spend.NewDatedEngine`/`ComputeAt`) pricing each event at the rate in effect *then* (embedded `pricing.yaml` = baseline); **Opus 4.x = $5/$25/$0.50** (Anthropic cut Opus at 4.5; a v0.38.0 'correction' to $15 was WRONG and reverted in v0.39.0 — LiteLLM confirmed the original $5 value). A live LiteLLM snapshot is adopted. **`tokenops status` now flags stale ingestion** — an enabled vendor-usage source with 0 events in 48h → soft `warnings` + remediation, `state` ready→degraded (v0.37.0, #131). **Opt-in usage-coaching hooks (ADR 0001)** shipped: a `coach-hook` Claude Code **Stop** hook that tracks cumulative per-session **$-budget** and fires graduated latched alerts (50/75/100% + over-budget escalation, default $50) nudging `/compact`; wired via `tokenops hooks install`. Joins `read-guard` (PreToolUse dedup) as the second Claude Code hook. Three planes: **proxy** (**17 providers**), **passive read** (Claude Code/Codex JSONL, opencode SQLite, quota scrapers), **MCP**. OpenAI uses an **exact tiktoken tokenizer**. Core prediction reads the **vendor's own rate-limit meter** across session_budget + plan_headroom. read-guard ACTIVE + agent-scoped. `fmt analyze --svg` emits bar + **weekly over-time charts**, selectable via **`--charts`**. Two klarlabs blog posts live: "800 to 1" and "The tool was guessing" (both from real tokenops output/commits).

## Last Session Summary
2026-07-08 (Session 2): fixed **snapshot-vs-baseline precedence** — the runtime effective engine was letting adopted LiteLLM snapshots override the vendor-verified baseline (deepseek-chat $0.14→$0.28, mistral-small $0.15→$0.06 at runtime). Added **verified-row pinning** (#151): `verified: true` catalog rows are stripped from snapshot overrides so the baseline wins; unpinned rows still auto-adopt. Pinned the non-Anthropic vendor-verified rows (the gap the Anthropic-only guard leaves). _Session 1, 2026-07-08: worked through the drift the all-provider refresh surfaced and **corrected the catalog against vendor pages** (#149, cross-checked every row — never single-sourced): Mistral keys → current -latest gen (Large 3 $0.50/$1.50, Medium 3.5 $1.50/$7.50, Small 4 $0.15/$0.60), deepseek-chat/reasoner → v4-flash $0.14/$0.28 cache $0.0028, o1 +cache $7.50, grok-3 +cache $0.75; codestral/gpt-3.5-turbo/gemini-1.5-flash confirmed FALSE drift, left alone. Then fixed **two adapter false-drift bugs** (#150) the refresh exposed — OpenAI MMDD misorder (gated 4-digit dates to YYMM) + distinct-SKU bleed (grok-3-fast/-mini no longer fold into grok-3). Both admin-merged (unreleased). Residual diff is now pure LiteLLM staleness (deepseek alias, mistral-small Small 3.2), not catalog error; `pricing lint` clean over **300 models**. Prior (2026-07-07): all-provider source (#147) + collision fix (#148); v0.40.0 (sonnet-5). See sessions/2026-07-08.md + 2026-07-07.md.

_Earlier today:_ 2026-07-07: "look at our usage" → shipped a feature. Parsed real transcripts:
~$50.6k API-equiv/7d, **79% cache-read from long sessions** (biggest: 7–9k turns /
~$2,400 each). tokenops's own data was stale ($0 cost — provider unset;
claude_code_jsonl poller idle since ~June 30) + a 0.25.1/0.30.1/0.35.0 version
skew (stale long-lived MCP `serve` process) — both fixed. Wrote ADR 0001 (opt-in
coaching hooks), built the `coach-hook` Stop nudge + `hooks install` (#127), then
replaced the flat per-turn threshold with a **cumulative $-budget + graduated
alerts** (#128, the GitHub-Actions-budget idea — per-turn missed long-flat
sessions). Released **v0.36.0** (#129), brew-upgraded, both hooks unified on the
homebrew binary (no mismatch). Then (Session 2) closed the "tokenops undercounts
its own usage" thread with a **stale-ingestion health warning** in `tokenops
status` (#131), released **v0.37.0** (#132), brew-upgraded + verified live (flags
`anthropic-cookie`, correctly not `claude_code_jsonl`). See sessions/2026-07-07.md.

## Next Session Should
Snapshot-vs-baseline is RESOLVED (verified-row pinning, #151). Consider **releasing**
#147–#151 (catalog now vendor-accurate + framework hardening) if a version is wanted.
Fast-follow: **pricing show/diff pin-awareness** (they render the raw snapshot, so
pinned rows still display the stale source value + show as drift — annotate so a human
isn't tempted to "fix" the baseline toward stale). Optional/older: gemini-2.5 cache-read
verified pass; Phase 2+ of ADR 0001 (SessionStart brief, UserPromptSubmit guardrail,
scorecard digest, Codex/Cursor Stop parity); user live-verifies an OpenAI-compat provider
(flips 9 providers unit→live); plan-catalog tiers / Bedrock SigV4; CI path-filter gap.

## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
- WAITING: user to live-verify an OpenAI-compat provider (hand-off script provided); would flip 9 providers from unit- to live-verified.
