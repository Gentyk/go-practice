"""Общие функции: конфиг, проверка файлов по SSH."""

import json
import os
from typing import Optional

DOMESTIC_TEMPLATE_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)), "domestic_template.json")

try:
    import paramiko
except ImportError:
    paramiko = None


def load_config(path: str) -> dict:
    """Читает config.ini и возвращает словарь ключ -> значение."""
    if not os.path.isfile(path):
        raise FileNotFoundError(f"Конфиг не найден: {path}")
    config = {}
    with open(path, "r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line or line.startswith(";"):
                continue
            if "=" in line:
                key, _, value = line.partition("=")
                config[key.strip()] = value.strip()
    return config


def load_domestic_template(path: Optional[str] = None) -> dict:
    """Читает domestic_template.json и возвращает словарь (конфиг xray)."""
    if path is None:
        path = DOMESTIC_TEMPLATE_PATH
    if not os.path.isfile(path):
        raise FileNotFoundError(f"Шаблон не найден: {path}")
    with open(path, "r", encoding="utf-8") as f:
        return json.load(f)


def check_file_exists(ssh: "paramiko.SSHClient", file_path: str) -> bool:
    stdin, stdout, stderr = ssh.exec_command(f"test -f {file_path} && echo EXISTS || echo NOT_FOUND")
    exit_code = stdout.channel.recv_exit_status()
    out = (stdout.read() + stderr.read()).decode().strip()
    if "EXISTS" in out:
        print(f"Файл на сервере найден: {file_path}")
        return True
    else:
        print(f"Файл на сервере отсутствует: {file_path}")
        raise FileNotFoundError(f"Файл на сервере отсутствует: {file_path}")
