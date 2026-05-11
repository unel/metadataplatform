#!/usr/bin/env python3
"""Create the stages/ skeleton for a new feature task.

Usage: scaffold-feature.py <feature-slug> [--flow <name>] [--dry-run]

Reads the flow graph from docs/standards/v2/flows/<name>.json.
Discovers steps from docs/standards/v2/ (authoritative step list).
Creates tasks/<slug>/stages/<stage>/<step>/ with:
  - status-log.md  — frontmatter (next-success-step, next-fail-step, previous-step) + pending
  - README.md      — copied from standards, navigation section injected from flow
"""
import argparse
import json
import subprocess
import sys
from pathlib import Path


def find_repo_root(start: Path) -> Path | None:
    p = start
    while p != p.parent:
        if (p / 'tasks').is_dir():
            return p
        p = p.parent
    return None


def load_flow(flows_dir: Path, name: str) -> dict[str, dict]:
    """Returns {step: {next-success, next-fail, previous}}."""
    flow_file = flows_dir / f'{name}.json'
    if not flow_file.exists():
        print(f"error: flow not found: {flow_file}", file=sys.stderr)
        sys.exit(1)
    data = json.loads(flow_file.read_text(encoding='utf-8'))
    return {s['step']: s for s in data['steps']}


def step_exists_in_standards(standards_dir: Path, step: str) -> bool:
    stage, step_name = step.split('/', 1)
    return (standards_dir / stage / step_name).is_dir()


def get_datetime() -> str:
    return subprocess.run(
        ['date', '-u', '+%Y-%m-%dT%H:%M:%SZ'], capture_output=True, text=True,
    ).stdout.strip()


def build_status_log(entry: dict, timestamp: str) -> str:
    return '\n'.join([
        '---',
        f'next-success-step: {entry.get("next-success", "")}',
        f'next-fail-step: {entry.get("next-fail", "")}',
        f'previous-step: {entry.get("previous", "")}',
        '---',
        '',
        f'# {timestamp} — pending',
        '',
    ])


def nav_value(step: str) -> str:
    return f'`{step}`' if step else '—'


def build_navigation_section(entry: dict) -> str:
    ns = entry.get('next-success', '')
    nf = entry.get('next-fail', '')
    prev = entry.get('previous', '')
    rows = [
        f'| Успех | {nav_value(ns)} |',
        f'| Провал | {nav_value(nf)} |',
        f'| Предыдущий шаг | {nav_value(prev)} |',
    ]
    return '## Навигация\n\n| Исход | Шаг |\n|---|---|\n' + '\n'.join(rows)


def build_readme(src: Path, task_path: str, generated: str, source_rel: str, nav_section: str) -> str:
    original = src.read_text(encoding='utf-8')
    meta = '\n'.join([
        '---',
        f'task: {task_path}',
        f'generated: {generated}',
        f'source: {source_rel}',
        '---',
        '',
    ])
    # Strip existing frontmatter
    if original.startswith('---\n'):
        _, _, rest = original.partition('\n---\n')
        original = rest.lstrip('\n')

    # Inject navigation before ## Артефакты процесса (or at end)
    marker = '\n## Артефакты процесса'
    if marker in original:
        original = original.replace(marker, f'\n{nav_section}\n{marker}')
    else:
        original = original.rstrip('\n') + f'\n\n{nav_section}\n'

    return meta + original


def main() -> None:
    parser = argparse.ArgumentParser(
        description='Create stages/ skeleton for a new feature task.',
    )
    parser.add_argument('slug', nargs='?', metavar='FEATURE-SLUG',
                        help='feature directory name under tasks/ (e.g. 0003-my-feature)')
    parser.add_argument('--flow', default='default', metavar='NAME',
                        help='flow definition name in docs/standards/v2/flows/ (default: default)')
    parser.add_argument('--dry-run', action='store_true')
    args = parser.parse_args()

    if not args.slug:
        parser.print_help()
        sys.exit(0)

    repo_root = find_repo_root(Path.cwd())
    if repo_root is None:
        print("error: cannot find repository root", file=sys.stderr)
        sys.exit(1)

    standards_dir = repo_root / 'docs' / 'standards' / 'v2'
    flows_dir = standards_dir / 'flows'
    if not standards_dir.is_dir():
        print(f"error: standards directory not found: {standards_dir}", file=sys.stderr)
        sys.exit(1)

    flow = load_flow(flows_dir, args.flow)
    steps = list(flow.keys())
    feature_dir = repo_root / 'tasks' / args.slug
    stages_dir = feature_dir / 'stages'

    if stages_dir.exists():
        print(f"error: {stages_dir.relative_to(repo_root)} already exists", file=sys.stderr)
        sys.exit(1)

    task_path = f"tasks/{args.slug}"
    timestamp = get_datetime()
    warnings: list[str] = []

    if args.dry_run:
        print(f"dry-run: {len(steps)} steps → {task_path}/stages/  (flow: {args.flow})")
        for step in steps:
            stage, step_name = step.split('/', 1)
            step_dir = stages_dir / stage / step_name
            print(f"  {(step_dir / 'status-log.md').relative_to(repo_root)}")
            src = standards_dir / stage / step_name / 'README.md'
            rm = '' if src.exists() else ' [!] source README missing'
            print(f"  {(step_dir / 'README.md').relative_to(repo_root)}{rm}")
        return

    created = 0
    for step in steps:
        entry = flow[step]

        stage, step_name = step.split('/', 1)
        if not step_exists_in_standards(standards_dir, step):
            warnings.append(f"warning: '{step}' not found in standards, skipping")
            continue
        step_dir = stages_dir / stage / step_name
        step_dir.mkdir(parents=True, exist_ok=True)

        (step_dir / 'status-log.md').write_text(
            build_status_log(entry, timestamp), encoding='utf-8',
        )
        print(f"  {(step_dir / 'status-log.md').relative_to(repo_root)}")

        src_readme = standards_dir / stage / step_name / 'README.md'
        if src_readme.exists():
            source_rel = str(src_readme.relative_to(repo_root))
            nav = build_navigation_section(entry)
            (step_dir / 'README.md').write_text(
                build_readme(src_readme, task_path, timestamp, source_rel, nav),
                encoding='utf-8',
            )
            print(f"  {(step_dir / 'README.md').relative_to(repo_root)}")
        else:
            warnings.append(f"warning: source README.md not found for '{step}', skipped")

        created += 1

    for w in warnings:
        print(w, file=sys.stderr)

    print(f"\n{task_path}/stages/ — {created} шагов создано  (flow: {args.flow})")


if __name__ == '__main__':
    main()
