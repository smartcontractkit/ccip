#!/bin/bash
set -e

# Detect if in CI to skip hooks
# https://docs.github.com/en/actions/learn-github-actions/variables#default-environment-variables
if [[ $CI == "true" ]]; then
    exit 0
fi

cd ../
chmod +x contracts/.husky/*.sh
pnpm husky contracts/
