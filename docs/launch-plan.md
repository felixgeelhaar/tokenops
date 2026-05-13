# Launch Plan — TokenOps Wedge (v0.8.x → v0.9.0)

GTM-skill recommendation: post one Loom + one Show HN + four Discord posts + 25 founder DMs. Target 10 installs and 3 discovery-call bookings inside a 7-day window. No new features ship until at least one real-user signal lands.

## Day 1 — Loom + Show HN

### Loom outline (90 seconds)

**Scene 1 — 0:00–0:20 — The struggling moment**
Open in Claude Code mid-refactor. Type a long prompt. Cut to a screen recording of an actual Anthropic rate-limit error message ("You've reached your usage limit. Try again in 2h 14m."). Voice-over:
> "If you've used Claude Max 20x for real work, you've seen this screen. There's no warning before it hits. TokenOps predicts it from inside your agent."

**Scene 2 — 0:20–0:55 — The fix**
Terminal split next to Claude Code:
1. `brew install felixgeelhaar/tap/tokenops` (skip if reviewer impatience; cut to installed prompt).
2. `tokenops init` (one line of output).
3. `tokenops plan set anthropic claude-max-20x` ("set plans.anthropic = claude-max-20x").
4. In Claude Code, ask the agent: "How much headroom do I have right now?" Agent calls `tokenops_session_budget`. Response renders inline:
   ```
   window: 71 / 200 (35.5%) — resets in 5h
   recommended_action: continue
   confidence: high
   ```

**Scene 3 — 0:55–1:30 — The honest part**
Voice-over:
> "TokenOps observes MCP traffic, not your raw Claude turns — yet. Every response says so. When you're 80% through your window with 90 minutes left, it says `slow_down`. When you're cooked, it says `wait_for_reset`."
Cut to the response payload with `signal_quality.level: low` highlighted.
> "Open source. Local-first. Brew install above. Github link in the description."

### Show HN post

**Title** (≤80 chars):
> Show HN: Predict your Claude Max rate-limit cutoffs from inside Claude Code

**Body**:
> Hi HN — I built TokenOps because I was tired of hitting Claude Max 20x rate limits mid-refactor with no warning. It's a local Go daemon + MCP server. After `tokenops init` and `tokenops plan set anthropic claude-max-20x`, the agent can call `tokenops_session_budget` and get back `{continue|slow_down|switch_model|wait_for_reset}` with a confidence band and the time-until-cap.
>
> Honesty note: the v0.8.1 release explicitly types `signal_quality` on every response. Right now the math runs on MCP-tool activity, not on your actual Claude turns (those don't flow through TokenOps). So the signal is a heuristic. The product says so. Vendor `/usage` API ingestion is queued.
>
> What's in for v0.9: empty-state scorecard with a first-week checklist instead of a useless F grade; demo-data isolation so seeded events don't contaminate real rollups; hot-reload on `tokenops plan set` so MCP doesn't need a manual reconnect.
>
> Repo: https://github.com/felixgeelhaar/tokenops
>
> Looking for: 5 people on Claude Max 5x/20x or ChatGPT Plus who'd hop on a 15-minute discovery call. Reply or DM @felixgeelhaar.

**First comment (plant immediately, by author)**:
> Three things I'd love feedback on:
> 1. Is `recommended_action` the right output, or do you just want the raw "N hours until reset"?
> 2. Anyone using GitHub Copilot Premium-request math? I haven't built that surface yet — would value the use case.
> 3. The MCP-ping heuristic vs. proxying client traffic — which would you actually wire up?

## Day 1 — Discord cross-post

Identical message, adapted to each community's norms:

### Claude Builders #show-and-tell
> Hey 👋 just shipped a small tool that predicts Max plan rate-limit cutoffs from inside Claude Code (MCP server, local-only, brew install). 90s Loom below. Would love 5 of you to install it and trade 15 min on a call so I can hear what's missing. [Loom link] [Repo link]

### Cursor Discord #showcase
> Built a TokenOps MCP server that gives Cursor agents a `session_budget` tool for plan headroom. Works for Cursor Pro and Claude Max 20x today; ChatGPT Plus and Copilot in catalog. Looking for 3 beta users to trade install for feedback. [Loom] [Repo]

### aider Discord
> Aider users on Claude subscriptions: shipped a CLI + MCP server that predicts rate-limit cutoffs from session activity. Works alongside aider's existing budget tracking. Free / open source. Loom: [link]. Github: [link]. Discovery calls open this week.

### AI Engineer Foundation Discord
> Released TokenOps v0.8.1 — MCP-resident headroom prediction for Claude Max / ChatGPT Plus / Copilot / Cursor. Honest about the signal-quality gap (response carries a `signal_quality` field). Looking for 5 power users to install and chat. [Loom] [Repo]

## Day 2–5 — 25 founder DMs

Find 25 solo AI-native founders / staff engineers on X or Bluesky who recently posted about AI workflow pain. Adapt this template per person:

> Hi {name} — saw your post about {their specific complaint, paraphrased to prove you read it}. I shipped a small CLI that predicts Claude Max 20x rate-limit cutoffs from inside Claude Code: {Loom link}. Open source, brew install, no signup. If it's a 5-second "no" — totally fine, ignore. If it'd be useful, I'd love 15 minutes to watch you install it and see what breaks.

Rules:
- Five DMs per day max. Quality > volume.
- Personalise the first sentence — never generic.
- Loom link first, repo link second. Watching is lower-friction than installing.
- If they reply: book a Calendly slot inside 24 hours. Use `docs/customer-discovery.md` interview script.

## Day 4 — Long-form blog post

Title (working): "How I predict Claude Max rate-limit cutoffs without hitting the Anthropic API."

Word count: 1200–1500. Sections:
1. The struggling moment (your own story).
2. Why vendor dashboards don't help (after-the-fact, no prediction).
3. The MCP-ping heuristic — what it actually counts and what it doesn't.
4. The `signal_quality` honesty contract.
5. Install in 90 seconds. Drop-in for Claude Code / Cursor / aider.
6. What's broken / on the roadmap (vendor `/usage` ingestion).

Cross-post: dev.to, Hashnode, personal blog. Submit to HN under "Show HN" only if Day 1's Show HN didn't fire. Otherwise just regular submission.

## Success criteria (by Day 7)

| Metric | Target |
|---|---|
| Loom views | ≥200 |
| Show HN front-page (≥80 points) | yes / no |
| GitHub stars added | ≥30 |
| `tokenops` brew installs (homebrew analytics) | ≥10 |
| DM replies | ≥5 |
| Discovery calls booked | ≥3 |
| Calls completed + tracker rows filled | ≥3 |

If targets miss by >50% on day 7, **the wedge hypothesis is in question.** Do not ship more code. Re-read the interview transcripts (even if only 1 or 2). Decide whether to pivot, narrow, or pause.

## What NOT to do

- No paid ads. No SEO marathon. No conference submissions. No category-creation pitch.
- Do not ship a new feature this week. The product is good enough to test demand.
- Do not write a "What we believe about AI" manifesto. Nobody on day 1 wants that.
- Do not respond to negative HN comments defensively. Note them in `docs/launch-tracker.md`, move on.

## Post-launch retrospective (Day 8)

Update `docs/customer-discovery.md` synthesis section with the actual interview data. File a Roady feature per top-three product gap surfaced. Kill any in-flight feature that doesn't trace to an interview quote.
