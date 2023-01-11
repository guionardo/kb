import glob
import os
import subprocess
from datetime import datetime


def get_md():
    return glob.glob("docs/**/*.md", recursive=True)


def read_last_change(filename: str) -> datetime:
    process = subprocess.run(
        ['git', '--no-pager', 'log', '-p', '-1', '--', filename], capture_output=True)
    lines = process.stdout.decode('utf-8').splitlines(keepends=False)

    for line in lines:
        if line.startswith('Date:'):
            # Date:   Thu Dec 15 10:57:15 2022 -0300
            ds = line.split(':', 1)[1].strip()
            d = datetime.strptime(ds, '%a %b %d %H:%M:%S %Y %z')
            break
    else:
        ltz = datetime.now().astimezone()
        d = datetime.fromtimestamp(os.path.getmtime(
            filename)).replace(tzinfo=ltz.tzinfo)

    return d


def parse_file_title(filename: str):
    with open(filename) as file:
        lines = file.readlines()
    name, _ = os.path.splitext(filename)
    if not lines:
        return name
    if lines[0].startswith('---'):
        for line in lines[1:]:
            if line.startswith('title:'):
                return line.split(':', 1)[1].strip()
            if line.startswith('---'):
                break
    for line in lines:
        if line.startswith('#'):
            while line.startswith('#'):
                line = line.removeprefix('#')
            return line.strip()

    return name


def parse_snake_case(text: str):
    words = text.split('_')
    return ' '.join([w.capitalize() for w in words])


def main():
    files = []
    for file in glob.iglob("docs/**/*.md", recursive=True):
        files.append((file, read_last_change(file)))

    sorted_files = sorted(files, key=lambda f: f[1], reverse=True)
    last_changed_file = os.path.abspath('./LAST_CHANGE.md')
    with open(last_changed_file, 'w') as last:
        last.write('# Last changes\n\n')
        last.write('| Group | Doc | Update |\n')
        last.write('|-------|-----|--------|\n')
        for (file, change_date) in sorted_files[:10]:
            title = parse_file_title(file)
            file_link, _ = os.path.splitext(file.removeprefix('docs/'))
            group = parse_snake_case(
                os.path.dirname(file.removeprefix('docs/'))) or 'Home'

            last.write(
                f'| {group} | [{title}](kb/{file_link}) | {change_date:%Y-%m-%d} |\n')


if __name__ == '__main__':
    main()
