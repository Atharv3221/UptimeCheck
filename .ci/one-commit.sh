#!/bin/bash

set -e

echo "Checking number of commits in PR branch (base: $GITHUB_BASE_REF)..."

git fetch origin "$GITHUB_BASE_REF"
git fetch origin "$GITHUB_HEAD_REF"

BASE_SHA=$(git merge-base origin/"$GITHUB_BASE_REF" origin/"$GITHUB_HEAD_REF")

COMMIT_COUNT=$(git rev-list --count "$BASE_SHA"..origin/"$GITHUB_HEAD_REF")

echo "Commits in PR: $COMMIT_COUNT"

if [ "$COMMIT_COUNT" -gt 1 ]; then
  echo "More than one commit. Please squash."
  exit 1
else
  echo "Only one commit found."
fi
