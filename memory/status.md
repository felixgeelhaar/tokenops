---
updated: 2026-07-04
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics). Repo `github.com/klarlabs-studio/tokenops`, module `go.klarlabs.de/tokenops` (vanity); brew tap `felixgeelhaar/tap/tokenops`. Latest release **v0.34.0** (released 2026-07-04). Post-v0.34.0 on `main` (through 8626669): the over-time chart feature. Three planes: **proxy** (**17 providers**), **passive read** (Claude Code/Codex JSONL, opencode SQLite, quota scrapers), **MCP**. OpenAI uses an **exact tiktoken tokenizer** (o200k_base, offline). Core prediction reads the **vendor's own rate-limit meter** across session_budget + plan_headroom (window + monthly). read-guard ACTIVE + agent-scoped. `fmt analyze --svg` now also emits **weekly over-time charts** (tokens / volume / composition), used in the klarlabs "800:1" blog post (live).

## Last Session Summary
Delivered a 4-agent review's four directions (vendor-meter predictions, exact tokenizer, optimizer honesty, 17 providers) + follow-ups, cut v0.34.0, added AGENTS.md failure modes, then built over-time charts (svgchart Lines/StackedArea + weekly jsonlfmt bucketing) and embedded two of them (composition + tokens) in the 800:1 blog post. Site deployed. All CI-green.

## Next Session Should
Nothing pressing. Optional: (a) make the blog's "roughly six months" prose exact (~5 weeks of on-disk timestamped logs); (b) user live-verifies an OpenAI-compat provider via the hand-off script; (c) plan-catalog tiers (need sourced caps) or Bedrock SigV4. Consider a v0.35.0 only once the over-time feature has more to ride with.

## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
- WAITING: user to live-verify an OpenAI-compat provider (hand-off script provided); would flip 9 providers from unit- to live-verified.
