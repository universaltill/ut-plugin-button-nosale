# Code review: auto-tag-release.yml (release-on-merge automation rollout)

**Date:** 2026-09-08
**Card:** ut-docs#1700 (mechanical rollout, continued from ut-docs#1694)
**Author:** scrum-master pipeline (cloud cycle, `lane:cloud-41`), on behalf of Farshid Mirza

## What changed

Added `.github/workflows/auto-tag-release.yml`, copied byte-for-byte from
the canonical, independently-reviewed copy in `ut-plugin-tax-de`
(`docs/code-reviews/2026-09-07-auto-tag-release-workflow-1694.md`). No
repo-specific deviation was needed or made.

## Independent review (fresh-context Sonnet subagent, per `complexity:easy` routing)

Verdict: **SAFE TO MERGE**, no findings. The subagent independently ran
`diff` against the canonical `ut-plugin-tax-de` source (byte-identical),
parsed the copied YAML with `python3 -c "import yaml; yaml.safe_load(...)"`
(clean), confirmed `manifest.json` at repo root with a valid semver
`version` (`1.0.1`), confirmed `release.yml` is tag-triggered
(`push: tags: ["v*"]`) with `workflow_dispatch` inputs named exactly
`channel` (choice, includes `stable`) and `publish` (boolean) — no
adaptation needed — confirmed no conflicting tag/release automation
elsewhere in the repo, no `pull_request_target` exposure, and confirmed
`main` is this repo's actual default/integration branch (this repo-branch
check specifically is what the same review pass flagged as failing for
`ut-plugin-faq`, pulled from this batch as a result — see ut-docs#1700
for that finding).

Repo-specific note (non-blocking, informational): the existing tag
`v1.0.1` already matches the current manifest version, so this workflow
will simply no-op until the next version bump — the same no-drift proof
case the previous batch demonstrated for `ut-plugin-payment-demo`.

## Verification performed

- `diff` against canonical source: byte-identical.
- `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/auto-tag-release.yml'))"` — parses.
- Confirmed `manifest.json` version and `release.yml` dispatch input names via direct file inspection.
- DevOps to confirm no unexpected effect after merge (repo is already tag/version-synced, so no new tag/release run is expected on this merge itself — the workflow only acts on the *next* version bump).

## Non-goals confirmed out of scope

- Changing `release.yml`/`ci.yml` (no input-name mismatch found).
- The remaining `ut-plugin-*` repos in ut-docs#1700's scope (tracked on
  that issue; `ut-plugin-faq` specifically needs a repo-specific decision
  before it can be rolled out — see the finding filed there).
