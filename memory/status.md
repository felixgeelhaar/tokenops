---
updated: 2026-07-04
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics). Repo `github.com/klarlabs-studio/tokenops`, module `go.klarlabs.de/tokenops` (vanity); brew tap `felixgeelhaar/tap/tokenops`. Latest RELEASE is still **v0.33.0** — the post-v0.33.0 work on `main` (through commit c07cdd0) is UNRELEASED. Three planes: **proxy** (now **17 providers** — +Ollama/LM Studio/LiteLLM/Vercel), **passive read** (Claude Code/Codex JSONL, opencode SQLite, quota scrapers), **MCP**. OpenAI now uses an **exact tiktoken tokenizer** (o200k_base, offline); others heuristic. The core prediction now reads the **vendor's own rate-limit meter** (not a message count) across session_budget + plan_headroom (window + monthly). read-guard is ACTIVE + agent-scoped.

## Last Session Summary
Big session. read-guard cross-agent fix (agent_id-scoped ledger). Ran a 4-agent review → delivered all 4 chosen directions: (1) core prediction now reads the authoritative vendor quota % — the headline gap; (2) exact tiktoken OpenAI tokenizer; (3) optimizer honesty (real savings estimate, retrieval_prune scored below the gate, dedupe doc de-hyped); (4) 13→17 providers. Then follow-ups: plan_headroom window + monthly authoritative (Copilot/Cursor get a real risk), dedupe footgun hardened. All CI-green through c07cdd0.

## Next Session Should
Consider cutting **v0.34.0** — a lot of unreleased value sits on main (17 providers, exact tokenizer, the authoritative-meter prediction rewrite, optimizer honesty). Verify the changelog covers it, then tag. Optionally: user live-verifies a provider via the hand-off script; or tackle plan-catalog tiers (need sourced caps) / Bedrock SigV4.

## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
- WAITING: user to live-verify an OpenAI-compat provider (hand-off script provided); would flip 9 providers from unit- to live-verified.
