#!/usr/bin/env python3
import argparse
import json
import re
import sys
from pathlib import Path

ICONS = {
    'done':          '✓',
    'failed':        '✗',
    'in-progress':   '~',
    'pending':       '·',
    'stale':         '↻',
    'clarification': '!',
    'corrupted':     '?',
    'inconsistent':  '%',
}


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


def resolve_features(args: argparse.Namespace, repo_root: Path) -> list[Path] | None:
    tasks_dir = repo_root / 'tasks'

    if args.all:
        return sorted(d for d in tasks_dir.iterdir() if d.is_dir() and d.name != '_backlog')

    if args.feature:
        pattern = args.feature
        glob_pattern = '*' + '*'.join(pattern.split('/')) + '*'
        matches = sorted(d for d in tasks_dir.glob(glob_pattern) if d.is_dir() and d.name != '_backlog')

        max_m = 1 if args.unique else args.max_matches
        if max_m is not None and len(matches) > max_m:
            names = ', '.join(d.name for d in matches)
            print(f"error: --feature '{pattern}' matched {len(matches)} tasks (max {max_m}): {names}", file=sys.stderr)
            sys.exit(1)
        if not matches:
            print(f"error: no task matching '{pattern}' in {tasks_dir}", file=sys.stderr)
            sys.exit(1)
        return matches

    feat_name = find_feature_from_cwd(Path.cwd(), tasks_dir)
    if feat_name:
        feat_dir = tasks_dir / feat_name
        if feat_dir.is_dir():
            return [feat_dir]

    return None


def get_run_number(content: str) -> int:
    runs = re.findall(r'run=(\d+)', content)
    return max((int(r) for r in runs), default=0)


def parse_status_log(log_path: Path) -> tuple[str, str | None]:
    """Returns (status, timestamp). status is the bare word without run=N."""
    if not log_path.exists():
        return 'corrupted', None

    content = log_path.read_text(encoding='utf-8')
    entries = re.findall(r'^# (\S+) — (.+)$', content, re.MULTILINE)
    if not entries:
        return 'pending', None

    last_dt, last_raw = entries[-1]
    status_word = re.sub(r'\s+run=\d+.*', '', last_raw).strip().split()[0]
    return status_word, last_dt


def check_inconsistent(step_dir: Path, status: str) -> tuple[str, str | None]:
    """Returns (effective_status, detail | None)."""
    log_path = step_dir / 'status-log.md'
    if not log_path.exists():
        return 'corrupted', None

    content = log_path.read_text(encoding='utf-8')
    log_run = get_run_number(content)

    def max_file_run(pattern: str) -> int:
        nums = []
        for f in step_dir.glob(pattern):
            m = re.match(r'(?:brief|report)-(\d+)\.md', f.name)
            if m:
                nums.append(int(m.group(1)))
        return max(nums, default=0)

    brief_files = list(step_dir.glob('brief-*.md'))
    report_files = list(step_dir.glob('report-*.md'))
    has_any_brief = bool(brief_files)
    has_any_report = bool(report_files)
    has_brief_N = (step_dir / f'brief-{log_run:03d}.md').exists()
    has_report_N = (step_dir / f'report-{log_run:03d}.md').exists()

    has_run_entries = bool(re.search(r'run=\d+', content))

    if has_run_entries:
        files_run = max(max_file_run('brief-*.md'), max_file_run('report-*.md'))
        if (has_any_brief or has_any_report) and log_run != files_run:
            return 'inconsistent', f'run counter mismatch: log={log_run}, files={files_run}'

        if status == 'in-progress':
            if not has_brief_N and not has_report_N:
                return 'inconsistent', f'brief-{log_run:03d}.md missing'
            if has_brief_N and has_report_N:
                return 'inconsistent', 'report written but end-step not called'
        elif status in ('done', 'failed'):
            if not has_brief_N:
                return 'inconsistent', f'brief-{log_run:03d}.md missing'
            if not has_report_N:
                return 'inconsistent', f'report-{log_run:03d}.md missing'

    if status in ('pending', 'stale') and (has_any_brief or has_any_report):
        return 'inconsistent', 'stale artifacts from previous run'

    return status, None


def get_step_info(step_dir: Path) -> dict:
    status, timestamp = parse_status_log(step_dir / 'status-log.md')

    if status not in ('stale', 'clarification', 'corrupted'):
        status, detail = check_inconsistent(step_dir, status)
    else:
        detail = None

    return {'status': status, 'timestamp': timestamp, 'detail': detail}


def determine_active_step(steps: list[dict]) -> str | None:
    for priority in (('in-progress',), ('failed', 'clarification'), ('stale',), ('pending',)):
        for s in steps:
            if s['status'] in priority:
                return s['step']
    return None


def collect_feature_data(feat_dir: Path) -> dict:
    stages_dir = feat_dir / 'stages'
    feature_name = feat_dir.name

    if not stages_dir.is_dir():
        return {'feature': feature_name, 'error': 'no stages/ directory'}

    steps = []
    for log in sorted(stages_dir.rglob('status-log.md')):
        step_dir = log.parent
        step_name = str(step_dir.relative_to(stages_dir))
        info = get_step_info(step_dir)
        steps.append({'step': step_name, **info})

    return {'feature': feature_name, 'steps': steps, 'active_step': determine_active_step(steps)}


def format_verbose(r: dict) -> str:
    lines = [r['feature']]
    active = r.get('active_step')
    for s in r['steps']:
        is_active = s['step'] == active
        prefix = '→' if is_active else ' '
        icon = ICONS.get(s['status'], '*')
        label = s['status']
        if s.get('detail'):
            label += f" ({s['detail']})"
        ts = s['timestamp'] or ''
        lines.append(f"{prefix} {s['step']:<38} {icon} {label:<14} {ts}".rstrip())
    return '\n'.join(lines)


def format_short(r: dict) -> str:
    active = r.get('active_step')
    if active is None:
        return f"{r['feature']}  ✓  (all done)"
    for s in r['steps']:
        if s['step'] == active:
            ts = s['timestamp'] or ''
            return f"{r['feature']}  →  {active}  [{s['status']}]  {ts}".rstrip()
    return r['feature']


def format_json_obj(r: dict) -> dict:
    if 'error' in r:
        return {'feature': r['feature'], 'error': r['error']}
    step_list = []
    active = r.get('active_step')
    for s in r['steps']:
        obj: dict = {
            'step': s['step'],
            'status': s['status'],
            'timestamp': s['timestamp'],
            'active': s['step'] == active,
        }
        if s.get('detail'):
            obj['detail'] = s['detail']
        step_list.append(obj)
    return {'feature': r['feature'], 'active_step': active, 'steps': step_list}


def main() -> None:
    parser = argparse.ArgumentParser(description='Show pipeline status for a feature task.')
    parser.add_argument('--feature', metavar='PATTERN')
    parser.add_argument('--all', action='store_true')
    parser.add_argument('--short', action='store_true')
    parser.add_argument('--json', action='store_true')
    parser.add_argument('--dry-run', action='store_true')
    parser.add_argument('--max-matches', type=int, metavar='N')
    parser.add_argument('--unique', action='store_true')
    args = parser.parse_args()

    repo_root = find_repo_root(Path.cwd())
    if repo_root is None:
        print("error: cannot find repository root (no tasks/ directory in path)", file=sys.stderr)
        sys.exit(1)

    feat_dirs = resolve_features(args, repo_root)
    if feat_dirs is None:
        parser.print_help()
        sys.exit(0)
    results = [collect_feature_data(d) for d in feat_dirs]

    if args.json:
        if args.all or len(results) > 1:
            print(json.dumps([format_json_obj(r) for r in results], ensure_ascii=False, indent=2))
        else:
            print(json.dumps(format_json_obj(results[0]), ensure_ascii=False, indent=2))
        return

    parts = []
    for r in results:
        if 'error' in r:
            parts.append(f"{r['feature']}: error — {r['error']}")
        elif args.short:
            parts.append(format_short(r))
        else:
            parts.append(format_verbose(r))

    sep = '\n\n' if (args.all or len(results) > 1) else '\n'
    print(sep.join(parts))


if __name__ == '__main__':
    main()
