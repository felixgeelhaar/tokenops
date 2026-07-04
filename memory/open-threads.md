---
updated: 2026-07-03
---
## [OPEN]
- fmt learn threshold tuning — telemetry now STARTING to accrue (dogfooded v0.29.0: 14 runs, learn loop produced a `go raise` hint + `printf` next-formatter candidate). Was BLOCKED on data; now unblocking. Revisit once a realistic volume of real command runs exists; verify hints stay sensible (printf was a spurious test artifact — expected, human filters).
- read-guard: ACTIVE mode (v0.30.1, ~/.claude/settings.json; backups .pre-readguard.bak + .pre-active.bak). Flipped observe→active after real reclaimable appeared: at 101 reads/3 sessions, 45 repeat reads = 4 reclaimable (~5.4k tok) + 39 ranged + 2 post-edit. Watching `tokenops read-guard stats` — the `blocked`/`reclaimed` line should climb over more sessions. If the agent ever fights a needed block, revert to observe or restore a backup.

## [WAITING]
- 2026-07-04: User to live-verify an OpenAI-compat provider (OpenRouter) via the hand-off script — env sandbox blocks reading opencode's key + external call from my side. Would flip 9 new providers from unit-verified to live-verified.

## Resolved
- 2026-07-03: Validate command_fmt proxy plane — DONE. Added TestDefaultPipeline_CommandFmtCompressesToolOutput: a realistic Anthropic tool_result runs through the DEFAULT pipeline and surfaces a command_fmt event with real savings. (Live-traffic validation still ideal but the wiring is proven.)
- 2026-07-03: Commit vs gitignore memory system — DECIDED: commit (cross-machine continuity; reversible). Committed with the pipeline test.
- 2026-07-03: Move tokenops to klarlabs-studio — DONE. Repo transferred, module → go.klarlabs.de/tokenops (vanity), blog links updated, v0.32.0 goreleaser 307 (stale release owner) fixed.
- 2026-07-03: Full provider coverage — DONE. 4→13 proxy providers (OpenAI-compatible fleet via NewOpenAICompatible + Cohere), opencode SQLite passive reader (verified 48,540 real turns), `provider set` validation/presets, coverage docs + README matrix. Released v0.33.0 (assets + brew 0.33.0). Honest boundaries documented (Gemini CLI proxy-only, Bedrock SigV4, hosted out of reach).
- 2026-07-03: mnemos vanity/module wiring — VERIFIED end-to-end (`go get go.klarlabs.de/mnemos@latest` → v0.33.0 live). Nothing to fix.
- 2026-07-04: read-guard cross-agent block — FIXED (agent_id-scoped ledger). The AGENTS.md Known Failure Mode about it is now resolved in code (lands next release; installed brew binary still has the bug until upgrade).
- 2026-07-04: Core prediction ignored the vendor meter — FIXED across session_budget + plan_headroom (window + monthly). The single biggest accuracy gap the review found.
- 2026-07-04: Tokenizer accuracy — exact tiktoken OpenAI tokenizer wired (others still heuristic).
- 2026-07-04: Optimizer honesty (canary estimate, retrieval_prune quality, dedupe doc) — FIXED.
- 2026-07-04: Provider coverage 13→17 (local/gateway). Plan-catalog tiers DEFERRED (need sourced caps).
- 2026-07-04: v0.34.0 RELEASED (assets + brew 0.34.0). Everything from the review + follow-ups shipped.
- 2026-07-04: Over-time charts — built in tokenops (`fmt analyze --svg`) + embedded composition + tokens in the 800:1 blog post; klarlabs site deployed. Kept 800:1 (no number on chart); charts framed as recent-weeks zooms.
- 2026-07-04: "six months" prose loose end — FIXED to "several weeks" (verified events store spans ~5-6 weeks; token numbers all confirmed).
- 2026-07-04: Second blog post "The tool was guessing" — WRITTEN + PUBLISHED (live, deployed). Self-critical rate-limit-bug story.
- 2026-07-04: `--charts` selector + v0.35.0 RELEASED (assets + brew 0.35.0). periodLabel axis-label bug fixed + live blog charts regenerated.
