---
updated: 2026-07-03
note: append-only log — never edit or delete entries; supersede with "→ superseded [date]"
---
- 2026-07-03: Adopted Agent OS memory system — persistent cross-session state via memory/ + wiki/ + cadence skills.
- 2026-07-03: `fmt` formatters guarantee critical-line survival in the ENGINE (`enforceCritical`), not per-formatter — so user config formatters are exactly as safe as built-ins. This is the moat: rules can be data, the invariant lives in code.
- 2026-07-03: Learning is offline + advisory + gated, never runtime self-modification. `fmt learn --apply` only writes loss-level overrides (safe — never touches critical rules); new-formatter candidates are printed, not auto-written. Preserves determinism.
- 2026-07-03: Cloud CLIs (aws/gcloud/az) pass JSON through untouched — the generic consecutive-dup scrub corrupts structured output. Only table/text is compressed.
- 2026-07-03: Config formatters excluded from the proxy content-sniff plane (no reliable content signature); they only run when dispatched by explicit command token.
- 2026-07-03: Skipped low-value RTK-parity commands (ls/cat/find/grep/diff/wget) — mostly signal, little to compress; the generic scrub covers them safely.
- 2026-07-03: Commit the Agent OS memory system (memory/ + wiki/ + AGENTS.md) to the repo rather than gitignore — cross-machine continuity for a solo-maintained OSS project; reversible if it becomes noise for contributors.
- 2026-07-03: Proxy-plane validation done via a default-pipeline integration test (real Anthropic tool_result → command_fmt event with savings), not live traffic — proves the wiring; live-traffic validation deferred, not blocking.
- 2026-07-03: read-guard observe-first, active opt-in — a PreToolUse hard-block on re-reads risks degrading agent flow, so default is observe (log only); active only after the user sees real reclaimable numbers. PreToolUse can only allow/deny (not modify input), so block-or-allow; conservative = unchanged full same-session re-reads only.
- 2026-07-03: The honest outcome of the whole token investigation: for this user, none of the three levers (caveman output prose <1%, fmt command output ~1-4% of Bash, Read re-reads ~0 reclaimable) yields much — because their usage is already efficient. The value delivered was measurement/diagnostics that proved where tokens go and honestly reported the absence of easy wins, not a fabricated saving.
