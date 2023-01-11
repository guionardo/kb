requirements:
	python .github/scripts/update_requirements.py ./pyproject.toml ./requirements.txt

last_changed:
	python3 .github/scripts/last_changed.py

setup_hook:
	cp .github/hooks/pre-commit .git/hooks/pre-commit