---
updated: 2026-07-03
---
## Current State
tokenops is a local-first MCP server + CLI for flat-rate AI subscriptions (rate-limit prediction, spend analytics) that ships `tokenops fmt` — a deterministic command-output compression subsystem now at **51 formatters / 57 command tokens**, user-extensible via config, self-tuning learn loop, MCP-exposed. Latest release **v0.28.1** on homebrew (fast-follow catalog on main, unreleased). Agent OS memory committed. Proxy plane validated via default-pipeline integration test.

## Last Session Summary
Built `tokenops fmt` (5 releases v0.26.0→v0.28.1, 4 PRs). Then: validated the proxy plane (default-pipeline test), committed Agent OS memory, and did the catalog fast-follow — +oc/nomad/packer/gem/swift/nix (51 formatters). vault deferred (secret-bearing output).

## Next Session Should
Cut a release (v0.29.0) for the fast-follow catalog if desired. Otherwise: let usage telemetry accrue, then tune fmt learn thresholds (data-blocked); or build a config-mutating MCP tool.

## Blocked / Waiting
- BLOCKED: fmt learn threshold tuning — needs real usage telemetry before empirical tuning.
