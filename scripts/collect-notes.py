#!/usr/bin/env python3
"""Print notes/complaints for one or more agents across all stages of a feature.

Usage:
    collect-notes.py --feature <slug> --agent <name> [--agent <name>...] [--type notes|complaints] [--include-retro]

Arguments:
    --feature       feature slug pattern (e.g. store-crud)
    --agent         agent name (english: ada, harley, grimm, tank, crowley, aziraphale, bo, herman, user);
                    repeat for multiple; use "all" for everyone
    --type          notes, complaints, or all (default: all)
    --include-retro also show 06-retro/01-recall files in a separate section

Output hierarchy (absolute heading levels):
    #      <Agent>         — only when multiple agents
    ##     Процессные /    — only when --include-retro
           Итоговые
    ###    <stage>
    ####   <stage/step>
    #####  Notes / Complaints
           <file contents>
"""
import argparse
import re
import sys
from collections import defaultdict
from pathlib import Path


AGENT_NAMES: dict[str, str] = {
    'ada': 'ада',
    'harley': 'харли',
    'grimm': 'гримм',
    'tank': 'танк',
    'crowley': 'кроули',
    'aziraphale': 'азирафаль',
    'bo': 'бо',
    'herman': 'герман',
    'user': 'user',
}
AGENT_NAMES_REVERSE: dict[str, str] = {v: k for k, v in AGENT_NAMES.items()}
LOCALIZED_NAMES: set[str] = set(AGENT_NAMES.values()) - {'user'}

AGENT_DISPLAY: dict[str, str] = {
    'ada': 'Ада',
    'harley': 'Харли',
    'grimm': 'Гримм',
    'tank': 'Танк',
    'crowley': 'Кроули',
    'aziraphale': 'Азирафаль',
    'bo': 'Бо',
    'herman': 'Герман',
    'user': 'User',
}


TYPE_DISPLAY: dict[str, str] = {
    'notes': 'Notes',
    'complaints': 'Complaints',
}


def strip_frontmatter(content: str) -> str:
    """Remove YAML frontmatter (--- ... ---) from the start of content."""
    if not content.startswith('---'):
        return content
    _, _, rest = content.partition('\n')
    _, sep, rest = rest.partition('\n---')
    return rest.lstrip('\n') if sep else content


def remap_headings(content: str, container_level: int) -> str:
    """Remap headings so they sit under container_level.

    If there is exactly one heading at the top (minimum) level — it's a file title,
    drop it and remap deeper levels starting at container_level+1.
    If there are multiple headings at the top level — they are structural sections,
    remap all of them to container_level+1 (and deeper ones accordingly).
    """
    lines = content.split('\n')
    heading_re = re.compile(r'^(#{1,6})( .+)')

    levels = [len(m.group(1)) for line in lines if (m := heading_re.match(line))]
    if not levels:
        return content

    min_level = min(levels)
    top_count = levels.count(min_level)
    single_title = top_count == 1

    # When dropping the lone title, deeper headings shift as if min_level didn't exist:
    # new = container + (level - min_level)         [single_title: gap closed]
    # new = container + (level - min_level) + 1     [multi: top maps to container+1]
    offset = 0 if single_title else 1

    result = []
    for line in lines:
        m = heading_re.match(line)
        if m:
            level = len(m.group(1))
            if level == min_level and single_title:
                continue  # drop lone file-title heading
            new_level = container_level + (level - min_level) + offset
            result.append('#' * new_level + m.group(2))
        else:
            result.append(line)

    return '\n'.join(result).strip()


def find_repo_root(start: Path) -> Path | None:
    p = start
    while p != p.parent:
        if (p / 'tasks').is_dir():
            return p
        p = p.parent
    return None


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


def resolve_agents(raw: list[str]) -> list[str]:
    """Return list of English agent names. Warns on localized input, errors on unknown."""
    resolved = []
    for a in raw:
        if a == 'all':
            resolved.append('all')
        elif a in AGENT_NAMES:
            resolved.append(a)
        elif a in LOCALIZED_NAMES:
            eng = AGENT_NAMES_REVERSE[a]
            print(f"warning: use English agent name '{eng}', not '{a}'", file=sys.stderr)
            resolved.append(eng)
        else:
            valid = ', '.join(sorted(AGENT_NAMES.keys()) + ['all'])
            print(f"error: unknown agent '{a}'; valid names: {valid}", file=sys.stderr)
            sys.exit(1)
    return resolved


def file_agent(filename: str) -> str | None:
    """Extract agent suffix from filename; normalize localized names to English."""
    for prefix in ('notes-', 'complaints-'):
        if filename.startswith(prefix):
            suffix = filename[len(prefix):].removesuffix('.md')
            return AGENT_NAMES_REVERSE.get(suffix, suffix)  # локализованное → английское
    return None


def file_type(filename: str) -> str | None:
    if filename.startswith('notes-'):
        return 'notes'
    if filename.startswith('complaints-'):
        return 'complaints'
    return None


def collect(
    stages_dir: Path,
    agents: list[str],
    type_filter: str,
    include_retro: bool,
) -> dict:
    """Returns nested dict: agent → section → stage → step → type → content.
    section is 'pipeline' or 'retro'."""
    data: dict = defaultdict(lambda: defaultdict(lambda: defaultdict(lambda: defaultdict(dict))))

    for f in sorted(stages_dir.rglob('*.md')):
        name = f.name
        agent = file_agent(name)
        ftype = file_type(name)
        if agent is None or ftype is None:
            continue
        if 'all' not in agents and agent not in agents:
            continue
        if type_filter != 'all' and ftype != type_filter:
            continue

        rel = f.relative_to(stages_dir)
        parts = rel.parts  # ('01-spec', '01-write', 'notes-ада.md')
        if len(parts) < 3:
            continue

        stage = parts[0]
        step = f"{parts[0]}/{parts[1]}"
        section = 'retro' if stage == '06-retro' else 'pipeline'

        if section == 'retro' and not include_retro:
            continue

        content = f.read_text(encoding='utf-8').strip()
        data[agent][section][stage][step][ftype] = content

    return data


CONTENT_LEVEL = 5  # ##### Notes / Complaints


def print_stages(stages: dict) -> None:
    """Print ### stage → #### step → ##### Notes/Complaints."""
    for stage in sorted(stages):
        print(f"### {stage}\n")
        for step in sorted(stages[stage]):
            print(f"#### {step}\n")
            for ftype in ('notes', 'complaints'):
                if ftype not in stages[stage][step]:
                    continue
                raw = stages[stage][step][ftype]
                content = remap_headings(strip_frontmatter(raw), CONTENT_LEVEL)
                print(f"##### {TYPE_DISPLAY[ftype]}\n")
                print(content if content else "_(пусто)_")
                print()


def make_glob_pattern(agent: str, file_type: str) -> str:
    """agent is always an English name here."""
    if agent == 'all':
        agent_part = '*'
    elif agent == 'user':
        agent_part = 'user'
    else:
        loc = AGENT_NAMES.get(agent, agent)
        agent_part = f"{{{agent},{loc}}}" if loc != agent else agent
    if file_type == 'all':
        name = f"{{notes,complaints}}-{agent_part}.md"
    else:
        name = f"{file_type}-{agent_part}.md"
    return f"stages/**/{name}  (excluding 06-retro/)"


def print_agent_block(sections: dict, include_retro: bool, agent: str, file_type: str) -> None:
    """Print sections for one agent (no # agent header)."""
    if include_retro:
        print("## Процессные\n")
        if 'pipeline' in sections:
            print_stages(sections['pipeline'])
        else:
            print("_(файлы отсутствуют)_\n")
            print(f"_glob: `{make_glob_pattern(agent, file_type)}`_\n")
        if 'retro' in sections:
            print("## Итоговые\n")
            print_stages(sections['retro'])
    else:
        if 'pipeline' in sections:
            print_stages(sections['pipeline'])
        else:
            print("_(файлы отсутствуют)_\n")
            print(f"_glob: `{make_glob_pattern(agent, file_type)}`_\n")


def main() -> None:
    parser = argparse.ArgumentParser(
        description='Show notes/complaints for an agent across all stages.',
    )
    parser.add_argument('--feature', required=True, metavar='FEATURE',
                        help='feature slug pattern (e.g. store-crud)')
    parser.add_argument('--agent', metavar='NAME', action='append', default=[],
                        help='agent name (repeat for multiple, or "all")')
    parser.add_argument('--type', choices=['notes', 'complaints', 'all'], default='all',
                        dest='file_type', metavar='TYPE',
                        help='notes, complaints, or all (default: all)')
    parser.add_argument('--include-retro', action='store_true',
                        help='also show 06-retro/01-recall in a separate "## Итоговые" section')
    args = parser.parse_args()

    if not args.agent:
        print("error: --agent is required (or use --agent all)", file=sys.stderr)
        sys.exit(1)

    repo_root = find_repo_root(Path.cwd())
    if repo_root is None:
        print("error: cannot find repository root (no tasks/ directory in path)", file=sys.stderr)
        sys.exit(1)

    feat_dir = resolve_feature_dir(repo_root / 'tasks', args.feature)
    stages_dir = feat_dir / 'stages'

    if not stages_dir.is_dir():
        print(f"error: stages/ not found in {feat_dir.name}", file=sys.stderr)
        sys.exit(1)

    agents = resolve_agents(args.agent)
    data = collect(stages_dir, agents, args.file_type, args.include_retro)

    multi_agent = len(agents) > 1 or 'all' in agents

    for agent in (sorted(data) if 'all' in agents else agents):
        if multi_agent:
            print(f"# {AGENT_DISPLAY.get(agent, agent)}\n")
        if agent in data:
            print_agent_block(data[agent], args.include_retro, agent, args.file_type)
        else:
            print_agent_block({}, args.include_retro, agent, args.file_type)


if __name__ == '__main__':
    main()
