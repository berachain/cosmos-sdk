# PR Title Update Required

## Issue
The PR #333 title does not follow the Conventional Commits format, causing the "Lint PR" workflow to fail.

## Current Title
```
Fix CodeQL workflow conflicts with default setup
```

## Required Title
```
ci: fix CodeQL workflow conflicts with default setup
```

## Why This Matters
The repository requires all PR titles to follow the [Conventional Commits specification](https://www.conventionalcommits.org/). This PR is about CI configuration changes, so it should use the `ci:` prefix.

## How to Fix
Please update the PR title through the GitHub UI to:
**ci: fix CodeQL workflow conflicts with default setup**

This will allow the "Lint PR" workflow to pass.
