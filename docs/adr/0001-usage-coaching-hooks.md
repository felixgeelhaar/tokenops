# ADR 0001 — Usage-coaching hooks for in-client token reclamation

- Status: Accepted
- Date: 2026-07-07

Phase 1 implemented in this PR (feat/coach-hook-phase1).

## Context

tokenops reclaims tokens on the request path via the proxy, but a large class
of users runs Claude Code on a subscription: their traffic never reaches the
tokenops proxy, so the proxy can neither observe nor reclaim it. The reclaimable
waste still exists — it is just inside the client. Two levers dominate a long
Claude Code session:

1. **Redundant re-reads** — the same file read in full, unchanged, more than
   once in a session. Already addressed by `read-guard` (a PreToolUse hook).
2. **Cache-read drag** — every turn re-bills the entire accumulated context at
   the cache-read rate. A session that has grown a large prefix pays that toll
   on *every* turn until it is compacted or restarted. This is usually the
   single largest recurring cost in a long session, and the operator has a
   cheap lever (`/compact`, or a fresh session) but no signal telling them when
   to pull it.

Claude Code exposes lifecycle **hooks** (PreToolUse, Stop, …) that run local
commands with a JSON payload on stdin. This lets tokenops act *inside* the
client, where the tokens actually are, without touching the model traffic.

## Decision

Ship a family of Claude Code hooks that observe and coach in-client usage,
installable through a single `tokenops hooks install` scaffolder.

**Phase 1 (this PR):**

- **`coach-hook`** — a **Stop** hook. After each turn it reads the *tail* of the
  session transcript jsonl (never the whole multi-MB file, never anything
  off-machine), extracts the most recent turn's `cache_read_input_tokens`, and
  when that crosses a threshold surfaces a **non-blocking** nudge to `/compact`
  or start a fresh session. It uses Claude Code's `systemMessage` output channel
  (user-facing, does not force the agent to continue — unlike `decision:"block"`)
  and a per-session cooldown so it never nags every turn.
- **`tokenops hooks install / uninstall / status`** — idempotently merges the
  selected hook entries into `~/.claude/settings.json`, backing up the prior
  file and writing atomically, without clobbering unrelated hooks.

**Fail-open is sacred.** A coach must never disrupt a session. Every error path
(unreadable transcript, malformed payload, missing usage) results in exit 0 with
no output — the turn proceeds untouched.

**Pricing** reuses the existing `spend` pricing engine (embedded `pricing.yaml`
catalog) to price cache-read load for Opus-family models; a const mirrors the
catalog rate as a fallback when a model can't be resolved. Non-Opus models show
tokens without a dollar figure rather than a misleading one.

## Consequences

- Subscription users get real, actionable reclamation signal for the first time.
- The hook reads only local transcript tails — no data leaves the machine, and
  the read is cheap regardless of transcript size.
- The `hooks install` scaffolder becomes the shared install path for the whole
  hook family (`read-guard` included), replacing per-command copy-paste of
  settings blocks.

## Future phases (not in this PR)

- Active levers beyond nudging (e.g. auto-compact suggestions with structured
  output), additional signals (turn-over-turn growth, tool-call fan-out), and a
  consolidated `tokenops coach` dashboard over the combined ledgers.
