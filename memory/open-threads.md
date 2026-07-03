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

- read-guard: ACTIVE mode (v0.30.1, ~/.claude/settings.json; backups .pre-readguard.bak + .pre-active.bak). Flipped observe→active after real reclaimable appeared: at 101 reads/3 sessions, 45 repeat reads = 4 reclaimable (~5.4k tok) + 39 ranged + 2 post-edit. Small but real lever. Active applied live via Claude Code's settings file watcher (no restart). Now watching `tokenops read-guard stats` — the `blocked`/`reclaimed` line should climb over more sessions. If the agent ever fights a needed block, revert to observe or restore a backup.
