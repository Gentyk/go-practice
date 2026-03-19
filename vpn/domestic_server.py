#!/usr/bin/env python3
"""
Скрипт: чтение config.ini, проверка наличия скрипта, SSH-подключение и проверка файла xray config.json.
"""

import json
import os
import re
import sys

try:
    import paramiko
except ImportError:
    print("Установите paramiko: pip install paramiko")
    sys.exit(1)

from common import load_config, check_file_exists, load_domestic_template


CONFIG_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)), "config.ini")
REMOTE_FILE = "/usr/local/etc/xray/config.json"

PLACEHOLDER_PATTERN = re.compile(r"^<[^>]+>$")


def _path_str(path):
    """Преобразует путь (tuple ключей/индексов) в строку для вывода."""
    parts = []
    for p in path:
        parts.append(f"[{p}]" if isinstance(p, int) else f".{p}" if parts else str(p))
    return "".join(parts).lstrip(".")


def _get_by_path(obj, path):
    """Возвращает значение в obj по пути (tuple ключей/индексов)."""
    for key in path:
        obj = obj[key]
    return obj


def _format_remote_value(val):
    """Краткое описание значения для вывода (в т.ч. списки и объекты)."""
    if isinstance(val, list):
        return f"массив из {len(val)} элементов"
    if isinstance(val, dict):
        return f"объект с ключами: {', '.join(sorted(val.keys()))}"
    return val


def _walk(obj, path=()):
    """Рекурсивный обход JSON: выдаёт (path, value) для всех листовых значений."""
    if isinstance(obj, dict):
        for k, v in obj.items():
            yield from _walk(v, path + (k,))
    elif isinstance(obj, list):
        for i, v in enumerate(obj):
            yield from _walk(v, path + (i,))
    else:
        yield path, obj


def check_config(ssh, remote_file: str) -> None:
    """
    Скачивает конфиг с удалённого сервера, сравнивает с шаблоном.
    Выводит: заполнение полей-плейсхолдеров «<*>» и расхождения по остальным полям.
    """
    template = load_domestic_template()
    stdin, stdout, stderr = ssh.exec_command(f"cat {remote_file}")
    exit_code = stdout.channel.recv_exit_status()
    raw = (stdout.read() + stderr.read()).decode().strip()
    if exit_code != 0 or not raw:
        print(f"Не удалось прочитать файл на сервере: {remote_file}")
        return
    try:
        remote = json.loads(raw)
    except json.JSONDecodeError as e:
        print(f"На сервере невалидный JSON: {e}")
        return

    template_paths = dict(_walk(template))
    remote_paths = dict(_walk(remote))

    print("\n--- Заполнение полей-плейсхолдеров (<...>) ---")
    for path, t_val in template_paths.items():
        if isinstance(t_val, str) and PLACEHOLDER_PATTERN.match(t_val):
            try:
                r_val = _get_by_path(remote, path)
                if isinstance(r_val, (list, dict)):
                    r_val = _format_remote_value(r_val)
                print(f"  {_path_str(path)}: {t_val!r} -> {r_val!r}")
            except (KeyError, IndexError, TypeError):
                print(f"  {_path_str(path)}: плейсхолдер {t_val!r} -> не задан на сервере")

    print("\n--- Несовпадения (не плейсхолдеры) ---")
    found = False
    for path, t_val in template_paths.items():
        if isinstance(t_val, str) and PLACEHOLDER_PATTERN.match(t_val):
            continue
        r_val = remote_paths.get(path)
        if r_val is None and path not in remote_paths:
            print(f"  {_path_str(path)}: в шаблоне {t_val!r}, на сервере ключ отсутствует")
            found = True
        elif remote_paths.get(path) != t_val:
            print(f"  {_path_str(path)}: в шаблоне {t_val!r}, на сервере {remote_paths[path]!r}")
            found = True
    if not found:
        print("  нет")

    # Вывод routing.rules с сервера
    print("\n--- routing.rules (с сервера) ---")
    try:
        rules = _get_by_path(remote, ("routing", "rules"))
        if isinstance(rules, list):
            print(json.dumps(rules, ensure_ascii=False, indent=2))
        else:
            print(rules)
    except (KeyError, IndexError, TypeError):
        print("  отсутствует")


def main():
    # 1. Получаем данные из config.ini
    try:
        config = load_config(CONFIG_PATH)
    except FileNotFoundError as e:
        print(e)
        raise ValueError(f"Ошибка загрузки config.ini: {e}")

    host = config.get("RUS_SERVER_HOST")
    password = config.get("RUS_SERVER_PASSWORD")

    if not host or not password:
        print("В config.ini должны быть заданы RUS_SERVER_HOST и RUS_SERVER_PASSWORD")
        sys.exit(1)

    # 3. Подключаемся по SSH (root) и проверяем наличие файла на сервере
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    try:
        ssh.connect(
            hostname=host,
            port=22,
            username="root",
            password=password,
            timeout=10,
            allow_agent=False,
            look_for_keys=False,
        )
    except Exception as e:
        print(f"Ошибка подключения SSH: {e}")
        sys.exit(1)

    try:
        check_file_exists(ssh, REMOTE_FILE)
        check_config(ssh, REMOTE_FILE)
    finally:
        ssh.close()


if __name__ == "__main__":
    main()
