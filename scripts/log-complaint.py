#!/usr/bin/env python3
import importlib.util
import os
import sys

spec = importlib.util.spec_from_file_location(
    'log_note',
    os.path.join(os.path.dirname(os.path.abspath(__file__)), 'log-note.py'),
)
mod = importlib.util.module_from_spec(spec)
spec.loader.exec_module(mod)

mod.FILE_PREFIX = 'complaints'
mod.main()
