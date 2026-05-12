#!/usr/bin/env python3
import argparse
import json
import subprocess
import sys
from pathlib import Path

FILE_PREFIX = 'notes'


def get_datetime() -> str:
    return subprocess.run(
        ['date', '-u', '+%Y-%m-%dT%H:%M:%SZ'], capture_output=True, text=True,
    ).stdout.strip()


def next_entry_number(target: Path) -> int:
    if not target.exists():
        return 1
    content = target.read_text(encoding='utf-8')
    return content.count('\n## ') + (1 if content.startswith('## ') else 0) + 1


def find_repo_root(start: Path) -> Path | None:
    p = start
    while p != p.parent:
        if (p / 'tasks').is_dir():
            return p
        p = p.parent
    return None


def find_feature_from_cwd(cwd: Path, tasks_dir: Path) -> str | None:
    p = cwd
    while p != p.parent:
        try:
            rel = p.relative_to(tasks_dir)
            parts = rel.parts
            if parts and parts[0] != '_backlog':
                return parts[0]
        except ValueError:
            pass
        p = p.parent
    return None


def find_step_from_cwd(cwd: Path, tasks_dir: Path) -> Path | None:
    """Returns step dir if CWD is inside stages/<stage>/<step>/."""
    p = cwd
    while p != p.parent:
        try:
            rel = p.relative_to(tasks_dir)
            parts = rel.parts
            # tasks/<feat>/stages/<stage>/<step>/...
            if len(parts) >= 4 and parts[1] == 'stages':
                step_dir = tasks_dir / parts[0] / 'stages' / parts[2] / parts[3]
                if step_dir.is_dir():
                    return step_dir
        except ValueError:
            pass
        p = p.parent
    return None


def get_active_step_via_script(repo_root: Path) -> str | None:
    """Calls feature-status.py --json and returns active_step."""
    script = repo_root / 'scripts' / 'feature-status.py'
    try:
        result = subprocess.run(
            [sys.executable, str(script), '--json'],
            capture_output=True, text=True, cwd=Path.cwd(),
        )
        if result.returncode != 0:
            return None
        data = json.loads(result.stdout)
        return data.get('active_step')
    except Exception:
        return None


def resolve_step_dir(args: argparse.Namespace, repo_root: Path) -> Path:
    tasks_dir = repo_root / 'tasks'
    cwd = Path.cwd()

    if args.step:
        feat_name = find_feature_from_cwd(cwd, tasks_dir)
        if not feat_name:
            print("error: cannot determine feature from CWD", file=sys.stderr)
            sys.exit(1)
        stage, step = args.step.split('/', 1) if '/' in args.step else (None, None)
        if not stage or not step:
            print(f"error: --step must be <stage>/<step>, got '{args.step}'", file=sys.stderr)
            sys.exit(1)
        step_dir = tasks_dir / feat_name / 'stages' / stage / step
        if not step_dir.is_dir():
            print(f"error: step not found: {step_dir}", file=sys.stderr)
            sys.exit(1)
        return step_dir

    # CWD inside a step?
    step_dir = find_step_from_cwd(cwd, tasks_dir)
    if step_dir:
        return step_dir

    # CWD inside a feature but not in a step — ask feature-status.py
    feat_name = find_feature_from_cwd(cwd, tasks_dir)
    if feat_name:
        active = get_active_step_via_script(repo_root)
        if active:
            stage, step = active.split('/', 1)
            return tasks_dir / feat_name / 'stages' / stage / step
        # All done — fall back to retro/01-recall
        fallback = tasks_dir / feat_name / 'stages' / '06-retro' / '01-recall'
        fallback.mkdir(parents=True, exist_ok=True)
        print("warning: active step not found, writing to retro/01-recall", file=sys.stderr)
        return fallback

    print("error: cannot determine step from CWD; use --step", file=sys.stderr)
    sys.exit(1)


def main() -> None:
    parser = argparse.ArgumentParser(
        description=f'Append a line to {FILE_PREFIX}-<agent>.md for the current step.',
    )
    agent_group = parser.add_mutually_exclusive_group()
    agent_group.add_argument('--agent', metavar='AGENT')
    agent_group.add_argument('--user', action='store_true', help='shorthand for --agent user')
    parser.add_argument('--message', metavar='TEXT')
    parser.add_argument('--step', metavar='STAGE/STEP')
    parser.add_argument('--dry-run', action='store_true')
    args = parser.parse_args()

    if not args.agent and not args.user:
        parser.print_help()
        sys.exit(0)
    if not args.message:
        parser.print_help()
        sys.exit(0)

    if args.user:
        args.agent = 'user'

    repo_root = find_repo_root(Path.cwd())
    if repo_root is None:
        print("error: cannot find repository root (no tasks/ directory in path)", file=sys.stderr)
        sys.exit(1)

    step_dir = resolve_step_dir(args, repo_root)
    target = step_dir / f'{FILE_PREFIX}-{args.agent}.md'

    n = next_entry_number(target)
    ts = get_datetime()

    if args.dry_run:
        print(f"dry-run: would append to {target}:")
        print(f"  ## {n} — {ts}")
        print(f"  {args.message}")
        return

    with target.open('a', encoding='utf-8') as f:
        f.write(f"## {n} — {ts}\n{args.message}\n\n")

    print(f"→ {target.relative_to(repo_root)}")


if __name__ == '__main__':
    main()
