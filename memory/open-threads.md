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

- read-guard: OBSERVE mode (v0.30.1, ~/.claude/settings.json; backup .pre-readguard.bak). Early live data (41 reads / 2 sessions): 7 repeat reads, ALL ranged → 0 reclaimable. Consistent with the user already being 54% ranged (good read hygiene). Emerging honest conclusion: little/nothing to reclaim for this user; likely NOT worth flipping to active. Check `tokenops read-guard stats` after more sessions; only go active if a real reclaimable number appears.
