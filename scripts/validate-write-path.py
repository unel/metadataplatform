#!/usr/bin/env python3
"""PreToolUse hook: проверяет что агент пишет только в разрешённые пути.

Использование (в frontmatter AGENT.md):
  hooks:
    PreToolUse:
      - matcher: "Write|Edit"
        hooks:
          - type: command
            command: "python3 scripts/validate-write-path.py <agent-slug>"

Claude Code передаёт JSON в stdin. Exit 2 = блокировать операцию.
"""
import sys
import json
import re
import yaml
from pathlib import Path


def load_write_access(agent: str) -> dict:
    root = Path(__file__).parent.parent
    agent_md = root / ".claude" / "agents" / agent / "AGENT.md"
    if not agent_md.exists():
        return {}
    content = agent_md.read_text()
    fm = re.match(r"^---\n(.*?)\n---", content, re.DOTALL)
    if not fm:
        return {}
    try:
        return yaml.safe_load(fm.group(1)).get("write_access", {})
    except Exception:
        return {}


def matches_any(file_path: str, patterns: list[str]) -> bool:
    p = Path(file_path.lstrip("./"))
    for pattern in patterns:
        # pathlib.Path.match поддерживает ** начиная с Python 3.12
        if p.match(pattern):
            return True
    return False


def main():
    agent = sys.argv[1] if len(sys.argv) > 1 else ""
    if not agent:
        sys.exit(0)

    write_access = load_write_access(agent)
    if not write_access:
        sys.exit(0)

    allow = write_access.get("allow", [])
    on_violation = write_access.get(
        "on_violation",
        f"[{agent}]: запись в {{path}} не разрешена этому агенту.",
    )

    try:
        data = json.loads(sys.stdin.read())
    except Exception:
        sys.exit(0)

    file_path = data.get("tool_input", {}).get("file_path", "")
    if not file_path:
        sys.exit(0)

    if matches_any(file_path, allow):
        sys.exit(0)

    print(on_violation.replace("{path}", file_path), file=sys.stderr)
    sys.exit(2)


if __name__ == "__main__":
    main()
