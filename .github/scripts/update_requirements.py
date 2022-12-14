import os
import subprocess
import sys


def update_requirements(pyproject_file: str, requirements_file: str):
    pyproject_file = os.path.abspath(pyproject_file)
    requirements_file = os.path.abspath(requirements_file)
    if not os.path.exists(pyproject_file):
        raise FileNotFoundError(pyproject_file)

    # Read requirements
# [tool.poetry.dependencies]
# python = "^3.10"
# mkdocs = "^1.4.2"
# mkdocs-minify-plugin = "^0.6.2"
# mkdocs-material = "^8.5.11"
# pymdown-extensions = "^9.9"
# mkdocs-build-plantuml-plugin = "^1.7.4"
# mkdocs-mermaid2-plugin = "^0.6.0"
# mkdocs-material-extensions = "^1.1.1"
# mkdocs-git-revision-date-localized-plugin = "^1.1.0"

    print('Getting pip freeze')
    pip_freeze = {p.split('==')[0]: p.split('==')[1]
                  for p in subprocess.check_output(['pip', 'freeze']).decode(
        'utf-8').splitlines(keepends=False)}
    started = False
    packages = []
    print('Reading ', pyproject_file)
    with open(pyproject_file) as file:
        for line in file.readlines():
            line = line.strip()
            if line.startswith('[tool.poetry.dependencies]'):
                started = True
                continue
            if not started or line.startswith('['):
                started = False
                continue
            if started and '=' in line:
                package_name = line.split('=')[0].strip()
                packages.append(package_name)

    merged = {package: pip_freeze[package]
              for package in packages if package in pip_freeze}
    with open(requirements_file, 'w') as file:
        for package, version in merged.items():
            print(package, '==', version)
            file.write(f'{package}=={version}\n')

    print(requirements_file, 'saved')


if __name__ == '__main__':
    if len(sys.argv) < 3:
        raise RuntimeError(
            'Expected arguments [pyproject.toml, requirements.txt]')
    pyproject_file, requirements_file = sys.argv[1:3]
    update_requirements(pyproject_file, requirements_file)
