#!/bin/bash

set -e

# Detect context: PR or push
if [[ -n "$GITHUB_BASE_REF" && -n "$GITHUB_HEAD_REF" ]]; then
    BASE_BRANCH="$GITHUB_BASE_REF"
    HEAD_BRANCH="$GITHUB_HEAD_REF"
else
    # Fallback for push: compare with main
    BASE_BRANCH="main"
    HEAD_BRANCH="${GITHUB_REF##*/}"
fi

echo "Checking number of commits in PR branch (base: $BASE_BRANCH, head: $HEAD_BRANCH)..."

git fetch origin "$BASE_BRANCH"
git fetch origin "$HEAD_BRANCH"

BASE_SHA=$(git merge-base origin/"$BASE_BRANCH" origin/"$HEAD_BRANCH")
COMMIT_COUNT=$(git rev-list --count "$BASE_SHA"..origin/"$HEAD_BRANCH")

echo "Commits in PR: $COMMIT_COUNT"

if [ "$COMMIT_COUNT" -gt 1 ]; then
  echo "More than one commit. Please squash."
  exit 1
else
  echo "Only one commit found."
fi
