---
updated: 2026-07-04
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics). Repo `github.com/klarlabs-studio/tokenops`, module `go.klarlabs.de/tokenops` (vanity); brew tap `felixgeelhaar/tap/tokenops`. Latest release **v0.35.0** (2026-07-04); `main` is clean. Three planes: **proxy** (**17 providers**), **passive read** (Claude Code/Codex JSONL, opencode SQLite, quota scrapers), **MCP**. OpenAI uses an **exact tiktoken tokenizer**. Core prediction reads the **vendor's own rate-limit meter** across session_budget + plan_headroom. read-guard ACTIVE + agent-scoped. `fmt analyze --svg` emits bar + **weekly over-time charts**, selectable via **`--charts`**. Two klarlabs blog posts live: "800 to 1" and "The tool was guessing" (both from real tokenops output/commits).

## Last Session Summary
Big multi-part session (3 capture blocks). 4-agent review → four directions (vendor-meter predictions, exact tokenizer, optimizer honesty, 17 providers) + follow-ups; cut v0.34.0. Then built over-time charts (svgchart Lines/StackedArea + weekly bucketing) + a `--charts` selector; cut v0.35.0. Wrote and published a second blog post, "The tool was guessing" (the rate-limit prediction bug). Fixed a periodLabel axis-label bug + the "six months"→"several weeks" prose (verified against the events store). Everything released + deployed, CI-green.

## Next Session Should
Nothing pressing — clean stopping point. Optional: (a) user live-verifies an OpenAI-compat provider via the hand-off script (flips 9 providers unit→live verified); (b) plan-catalog tiers (need sourced caps) or Bedrock SigV4; (c) a third post if inspiration strikes (candidate angles noted in the session log).

## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
- WAITING: user to live-verify an OpenAI-compat provider (hand-off script provided); would flip 9 providers from unit- to live-verified.
