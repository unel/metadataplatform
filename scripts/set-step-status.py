#!/usr/bin/env python3
"""Set step status in a feature's status-log.md with validation.

Usage:
    set-step-status.py <feature> <stage/step> <status> [--comment TEXT] [--dry-run]

Statuses and their validation:
    in-progress   brief-NNN.md exists; predecessor step is terminal (done/failed)
    done/failed   last status was in-progress run=N; report-NNN.md exists;
                  for 03-fix steps: report-NNN.md has in-response-to: in frontmatter
    stale         --comment required
    clarification --comment required
    pending       no validation

Run number is auto-detected from status-log.md.
Supports both legacy format "(run N)" and new format "run=N" when reading existing logs.
New entries are written in "run=N" format.
Predecessor for run 1 is read from previous-step frontmatter field.
Predecessor for run N>1 is read from next-fail-step frontmatter field (the fix step).
"""
import argparse
import re
import subprocess
import sys
from pathlib import Path


TERMINAL = {'done', 'failed'}
REQUIRES_COMMENT = {'stale', 'clarification'}
VALID_STATUSES = {'in-progress', 'done', 'failed', 'stale', 'clarification', 'pending'}
RUN_STATUSES = {'in-progress', 'done', 'failed'}


def get_datetime() -> str:
    return subprocess.run(
        ['date', '-u', '+%Y-%m-%dT%H:%M:%SZ'], capture_output=True, text=True,
    ).stdout.strip()


def find_repo_root(start: Path) -> Path | None:
    p = start
    while p != p.parent:
        if (p / 'tasks').is_dir():
            return p
        p = p.parent
    return None


def parse_frontmatter(text: str) -> dict:
    if not text.startswith('---'):
        return {}
    _, _, rest = text.partition('\n')
    block, sep, _ = rest.partition('\n---')
    if not sep:
        return {}
    meta = {}
    for line in block.splitlines():
        parts = line.split(':', 1)
        if len(parts) == 2 and parts[0].strip():
            meta[parts[0].strip()] = parts[1].strip()
    return meta


def _extract_run(run_eq: str, run_paren: str) -> int:
    """Extract run number from either 'run=N' or '(run N)' capture groups."""
    if run_eq:
        return int(run_eq)
    if run_paren:
        return int(run_paren)
    return 0


def get_last_status_and_run(log_path: Path) -> tuple[str, int]:
    """Returns (last_status, last_run_number). run=0 if no run entries yet.
    Handles both legacy '(run N)' and new 'run=N' formats."""
    if not log_path.exists():
        return 'pending', 0
    content = log_path.read_text(encoding='utf-8')
    # Match: status word, optionally followed by run=N or (run N)
    entries = re.findall(
        r'^# \S+ — (\S+)(?:\s+run=(\d+)|\s+\(run\s+(\d+)\))?',
        content, re.MULTILINE,
    )
    if not entries:
        return 'pending', 0
    last_status, run_eq, run_paren = entries[-1]
    return last_status, _extract_run(run_eq, run_paren)


def get_max_run(log_path: Path) -> int:
    """Returns highest run number seen in log (0 if none).
    Handles both 'run=N' and '(run N)' formats."""
    if not log_path.exists():
        return 0
    content = log_path.read_text(encoding='utf-8')
    runs_eq = [int(x) for x in re.findall(r'run=(\d+)', content)]
    runs_paren = [int(x) for x in re.findall(r'\(run\s+(\d+)\)', content)]
    return max(runs_eq + runs_paren, default=0)


def get_step_last_status(step_dir: Path) -> str:
    """Returns last status of a step, 'pending' if not started."""
    status, _ = get_last_status_and_run(step_dir / 'status-log.md')
    return status


def resolve_step_dir(stages_dir: Path, step_path: str) -> Path | None:
    parts = step_path.split('/', 1)
    if len(parts) != 2:
        return None
    d = stages_dir / parts[0] / parts[1]
    return d if d.is_dir() else None


def validate_in_progress(step_dir: Path, run: int, meta: dict, stages_dir: Path) -> list[str]:
    errors = []

    # brief-NNN.md must exist before marking in-progress
    brief = step_dir / f'brief-{run:03d}.md'
    if not brief.exists():
        errors.append(
            f"brief-{run:03d}.md not found.\n"
            f"  Action: create {brief.name} in {step_dir} with:\n"
            f"    - what inputs you received (links to upstream reports)\n"
            f"    - what you plan to do this run\n"
            f"  Then re-run this command."
        )

    # Predecessor check: first run vs re-run
    if run == 1:
        prev_step = meta.get('previous-step', '').strip()
        if prev_step:
            prev_dir = resolve_step_dir(stages_dir, prev_step)
            if prev_dir is None:
                errors.append(
                    f"previous step '{prev_step}' directory not found in stages/.\n"
                    f"  Check that the feature was scaffolded correctly."
                )
            else:
                prev_status = get_step_last_status(prev_dir)
                if prev_status not in TERMINAL:
                    errors.append(
                        f"previous step '{prev_step}' is '{prev_status}' — not terminal.\n"
                        f"  Action: complete '{prev_step}' (must reach done or failed) before starting this step."
                    )
    else:
        # Re-run: fix step must have completed
        fix_step = meta.get('next-fail-step', '').strip()
        if fix_step:
            fix_dir = resolve_step_dir(stages_dir, fix_step)
            if fix_dir is None:
                errors.append(
                    f"fix step '{fix_step}' directory not found in stages/.\n"
                    f"  Check that the feature was scaffolded correctly."
                )
            else:
                fix_status = get_step_last_status(fix_dir)
                if fix_status not in TERMINAL:
                    errors.append(
                        f"fix step '{fix_step}' is '{fix_status}' — not terminal.\n"
                        f"  Action: complete '{fix_step}' before re-running this step."
                    )

    return errors


def validate_done_failed(
    step_dir: Path, run: int, last_status: str, step_path: str, repo_root: Path,
) -> list[str]:
    errors = []

    # Must currently be in-progress
    if last_status != 'in-progress':
        errors.append(
            f"current status is '{last_status}', expected 'in-progress'.\n"
            f"  Action: call 'in-progress' first, then write report-{run:03d}.md, then call done/failed."
        )

    # report-NNN.md must exist
    report = step_dir / f'report-{run:03d}.md'
    if not report.exists():
        errors.append(
            f"report-{run:03d}.md not found.\n"
            f"  Action: write {report.name} in {step_dir} summarising what was done,\n"
            f"  then re-run this command."
        )
    elif '03-fix' in step_path:
        # Fix steps must have in-response-to in report frontmatter
        content = report.read_text(encoding='utf-8')
        fm = parse_frontmatter(content)
        in_response_to = fm.get('in-response-to', '').strip()
        if not in_response_to:
            errors.append(
                f"report-{run:03d}.md is missing 'in-response-to:' in frontmatter.\n"
                f"  Action: add to the frontmatter of {report.name}:\n"
                f"    in-response-to: tasks/<feature>/stages/<group>/02-review/report-NNN.md\n"
                f"  pointing to the review report this fix responds to."
            )
        else:
            ref = Path(in_response_to)
            if not ref.is_absolute():
                ref = repo_root / in_response_to
            if not ref.exists():
                errors.append(
                    f"in-response-to path does not exist: '{in_response_to}'.\n"
                    f"  Action: correct the path in {report.name} frontmatter."
                )

    return errors


def resolve_feature_dir(tasks_dir: Path, pattern: str) -> Path:
    glob_pattern = '*' + '*'.join(pattern.split('/')) + '*'
    matches = sorted(
        d for d in tasks_dir.glob(glob_pattern)
        if d.is_dir() and d.name != '_backlog'
    )
    if not matches:
        print(f"error: no task matching '{pattern}' in {tasks_dir}", file=sys.stderr)
        sys.exit(1)
    if len(matches) > 1:
        names = ', '.join(d.name for d in matches)
        print(f"error: '{pattern}' matched multiple tasks: {names}", file=sys.stderr)
        sys.exit(1)
    return matches[0]


def main() -> None:
    parser = argparse.ArgumentParser(
        description='Set step status in status-log.md with pre-condition validation.',
    )
    parser.add_argument('feature', metavar='FEATURE', help='feature slug pattern')
    parser.add_argument('step', metavar='STAGE/STEP', help='e.g. 01-spec/02-review')
    parser.add_argument('status', choices=sorted(VALID_STATUSES))
    parser.add_argument('--comment', metavar='TEXT', help='comment appended to the log entry')
    parser.add_argument('--dry-run', action='store_true')
    args = parser.parse_args()

    repo_root = find_repo_root(Path.cwd())
    if repo_root is None:
        print("error: cannot find repository root (no tasks/ directory in path)", file=sys.stderr)
        sys.exit(1)

    feat_dir = resolve_feature_dir(repo_root / 'tasks', args.feature)
    stages_dir = feat_dir / 'stages'

    step_parts = args.step.split('/', 1)
    if len(step_parts) != 2:
        print(f"error: step must be <stage>/<step>, got '{args.step}'", file=sys.stderr)
        sys.exit(1)
    step_dir = stages_dir / step_parts[0] / step_parts[1]
    if not step_dir.is_dir():
        print(f"error: step directory not found: {step_dir}", file=sys.stderr)
        sys.exit(1)

    log_path = step_dir / 'status-log.md'
    meta = parse_frontmatter(log_path.read_text(encoding='utf-8')) if log_path.exists() else {}

    last_status, last_run = get_last_status_and_run(log_path)
    max_run = get_max_run(log_path)

    # Determine run number for this transition
    status = args.status
    if status == 'in-progress':
        run = max_run + 1
    else:
        run = max_run if max_run > 0 else 1

    # Validation
    errors = []

    if status in REQUIRES_COMMENT and not args.comment:
        errors.append(f"--comment is required for status '{status}'")

    if status == 'in-progress':
        errors.extend(validate_in_progress(step_dir, run, meta, stages_dir))
    elif status in ('done', 'failed'):
        errors.extend(validate_done_failed(step_dir, run, last_status, args.step, repo_root))

    if errors:
        for e in errors:
            print(f"error: {e}", file=sys.stderr)
        sys.exit(1)

    # Build log entry
    ts = get_datetime()
    header = f"# {ts} — {status} run={run}" if status in RUN_STATUSES else f"# {ts} — {status}"

    lines = [header]
    if status in ('done', 'failed') and (step_dir / f'report-{run:03d}.md').exists():
        lines.append(f"→ report-{run:03d}.md")
    if args.comment:
        lines.append(args.comment)

    entry = '\n'.join(lines) + '\n'

    if args.dry_run:
        rel = log_path.relative_to(repo_root)
        print(f"dry-run: would append to {rel}:")
        print(entry)
        return

    with log_path.open('a', encoding='utf-8') as f:
        f.write('\n' + entry)

    rel = log_path.relative_to(repo_root)
    run_label = f" run={run}" if status in RUN_STATUSES else ""
    print(f"→ {rel}: {status}{run_label}")


if __name__ == '__main__':
    main()
