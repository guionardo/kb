# KB
My personal knowledge-base

![GitHub](https://img.shields.io/github/license/guionardo/kb)
![GitHub top language](https://img.shields.io/github/languages/top/guionardo/kb)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/guionardo/kb/gh-pages)

[![GH Pages](https://img.shields.io/badge/GitHub%20Pages-guionardo%2Fkb-success?logo=markdown)](https://guionardo.github.io/kb/)

## Links

* [MkDocs](https://www.mkdocs.org/)
* [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/)

## Setup

- [x] Requirements: python3.10
- [x] Create virtual environment
    ```bash
    ❯ python3 -m venv .venv
    ❯ source .venv/bin/activate
    ```
- [x] Install poetry
    ```bash
    ❯ pip install poetry
    ```
- [x] Update packages
    ```bash
    ❯ poetry update
    ```
- [x] Setup hook to force update of last docs created
    ```bash
    ❯ make setup_hook
    ```
- [x] Create and edit files on ```docs``` folder
- [x] After push any commit, the github action pipeline will build and publish to github pages. 