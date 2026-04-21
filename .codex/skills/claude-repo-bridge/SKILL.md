---
name: "claude-repo-bridge"
description: "Use when the repository contains a `.claude/` directory or the user asks to continue Claude Code workflows, reuse Claude agents, or interpret Claude command files. Read the local Claude project files, map Claude agents to Codex behaviors, and follow the repository workflow without assuming Claude runtime features exist."
---

# Claude Repo Bridge

Use this skill when a repo contains Claude Code artifacts such as `.claude/project.yaml`, `.claude/commands/*.md`, `.claude/agents/*.md`, or `CLAUDE.md`.

## Quick start

1. Read `CLAUDE.md`.
2. Read `.claude/project.yaml`.
3. Read only the relevant files from `.claude/commands/` or `.claude/agents/`.
4. Treat Claude command files as workflow instructions, not executable runtime features.
5. Treat Claude agent files as role prompts that can be mapped onto Codex behavior.

## Runtime boundary

Claude-specific runtime features do not exist here unless explicitly recreated:

- No native Claude `Skill` tool.
- No native Claude `Agent` tool contract.
- No automatic `SendMessage` continuation semantics from Claude sessions.
- No assumption that Claude MCP servers are available just because `.claude/settings.local.json` references them.

Use the files as instructions and process artifacts, not as proof that the runtime exists.

## Mapping

Map Claude roles to Codex behavior like this:

- `harley`: orchestration. In Codex, keep this local unless the user explicitly asks for subagents. Use plans, task sequencing, and status tracking.
- `herman`: repository exploration. Use local codebase inspection. If subagents are explicitly allowed, prefer an `explorer` agent.
- `bo`: external research. Use web browsing with primary sources. If subagents are explicitly allowed and the task is sidecar research, use an `explorer` agent.
- `ada`: implementation. Work locally by default. If subagents are explicitly allowed and the task is parallelizable, use a `worker` agent with a clear file ownership boundary.
- `grimm`: review. Use a code-review mindset focused on bugs, regressions, and missing tests. Findings first.
- `tank`: specification and requirements. Write or revise task specs, edge cases, and NFRs.
- `crowley`: adversarial tests. Focus on failure modes, limits, and negative paths.
- `aziraphale`: positive-path tests. Focus on happy path, contracts, and acceptance coverage.

## Workflow rules

If the repo uses `tasks/<slug>/status.md`, preserve that workflow:

- Respect step ordering from `.claude/project.yaml`.
- Treat `needs-recheck` as not done.
- If an artifact is rewritten after its review passed, reset downstream statuses consistently.
- When a Claude command says to read standards or templates, load only the relevant files.

## Command handling

When the user references a Claude command such as `/code-write`, `/code-review`, or `/spec-review`:

1. Open the corresponding file in `.claude/commands/`.
2. Execute its intent using Codex tools and constraints.
3. Preserve any file-writing conventions in `tasks/<slug>/`.
4. Ignore Claude-only tool names and replace them with the nearest Codex-native behavior.

## What to read

Read only what is needed:

- `CLAUDE.md` for repo-level orchestration rules.
- `.claude/project.yaml` for transitions, retries, and proficiency.
- `.claude/agents/<name>.md` only for the role being used.
- `.claude/commands/<name>.md` only for the requested workflow step.
- `docs/standards/` only when a command or task actually depends on standards.

## Output style

- Be explicit when you are following a Claude workflow by interpretation rather than native runtime support.
- Prefer concrete paths and status implications over persona roleplay.
- Keep the Claude flavor subordinate to correctness and completion.
