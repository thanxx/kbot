#!/usr/bin/env bash
set -e

cp ./pre-commit ../.git/hooks
chmod +x ../.git/hooks/pre-commit

git config --bool hooks.gitleaks true

echo "GitLeaks pre-commit hook installed and enabled!"
