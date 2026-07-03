---
updated: 2026-07-03
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics). Repo lives at **github.com/klarlabs-studio/tokenops**, module path **go.klarlabs.de/tokenops** (vanity); brew tap unchanged (`felixgeelhaar/tap/tokenops`). Latest release **v0.33.0**. Instruments AI usage on three planes: **proxy** (now **13 providers** — openai/anthropic/gemini/mistral/cohere + the OpenAI-compatible fleet), **passive read** (Claude Code + Codex JSONL, and now **opencode via SQLite**), and **MCP** (`tokenops serve`). Also ships `tokenops fmt` (deterministic command-output compression) and `read-guard` (installed here in ACTIVE mode). Honest boundaries: Gemini CLI has no local token log (proxy only), Bedrock needs SigV4 (no auth hook), hosted agents (Jules) out of reach.

## Last Session Summary
Expanded proxy providers 4 → 13 (OpenAI-compatible framework via `NewOpenAICompatible` + Cohere with its own normalizer). Built the opencode SQLite passive reader (verified 48,540 real turns). Hardened `provider set` (validation + presets). Added integration coverage docs + README matrix. Fixed the v0.32.0 goreleaser 307 (stale owner) and released **v0.33.0** cleanly. Verified **mnemos** vanity/module wiring end-to-end (`go get go.klarlabs.de/mnemos@latest` → v0.33.0, live).

## Next Session Should
Provider/integration work is complete and released. Pick a new direction — likely candidates: (a) a `pricing.path` override example for the multiplexers (fireworks/together/openrouter), (b) a Bedrock SigV4 auth-hook if remote/enterprise providers matter, or (c) unrelated new work. Nothing is blocking.

## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
