---
updated: 2026-07-03
---
## [OPEN]

## [BLOCKED]
- fmt learn threshold tuning (MinRuns=5, GenericRatioFloor=0.5, AccessRateCeiling=0.10, AccessRateFloor=0.01) — defaults reviewed and sane; genuine empirical tuning is blocked on real usage telemetry accruing in ~/.tokenops/recovery/index.jsonl. Revisit when data exists.

## [WAITING]

## Resolved
- 2026-07-03: Validate command_fmt proxy plane — DONE. Added TestDefaultPipeline_CommandFmtCompressesToolOutput: a realistic Anthropic tool_result runs through the DEFAULT pipeline and surfaces a command_fmt event with real savings. (Live-traffic validation still ideal but the wiring is proven.)
- 2026-07-03: Commit vs gitignore memory system — DECIDED: commit (cross-machine continuity; reversible). Committed with the pipeline test.
