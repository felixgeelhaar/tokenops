---
updated: 2026-07-03
---
## [OPEN]

## [OPEN]
- fmt learn threshold tuning — telemetry now STARTING to accrue (dogfooded v0.29.0: 14 runs, learn loop produced a `go raise` hint + `printf` next-formatter candidate). Was BLOCKED on data; now unblocking. Revisit once a realistic volume of real command runs exists; verify hints stay sensible (printf was a spurious test artifact — expected, human filters).

## [WAITING]

## Resolved
- 2026-07-03: Validate command_fmt proxy plane — DONE. Added TestDefaultPipeline_CommandFmtCompressesToolOutput: a realistic Anthropic tool_result runs through the DEFAULT pipeline and surfaces a command_fmt event with real savings. (Live-traffic validation still ideal but the wiring is proven.)
- 2026-07-03: Commit vs gitignore memory system — DECIDED: commit (cross-machine continuity; reversible). Committed with the pipeline test.

- read-guard: installed in OBSERVE mode (~/.claude/settings.json). Waiting on the user to watch `tokenops read-guard stats` over a few sessions, then decide whether to flip to --mode active. Backup: ~/.claude/settings.json.pre-readguard.bak.
