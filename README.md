# api

Schema of the API types that are served by Karmada.

[![LICENSE](https://img.shields.io/github/license/karmada-io/api.svg)](/LICENSE)
[![Releases](https://img.shields.io/github/v/release/karmada-io/api)](https://github.com/karmada-io/api/releases/latest)
[![Slack](https://img.shields.io/badge/slack-join-brightgreen)](https://slack.cncf.io)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/5301/badge)](https://bestpractices.coreinfrastructure.org/projects/5301)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/karmada-io/api/badge)](https://securityscorecards.dev/viewer/?uri=github.com/karmada-io/api)
[![Go Report Card](https://goreportcard.com/badge/github.com/karmada-io/api)](https://goreportcard.com/report/github.com/karmada-io/api)
[![Artifact HUB](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/karmada)](https://artifacthub.io/packages/krew/krew-index/karmada)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B28176%2Fgithub.com%2Fkarmada-io%2Fkarmada.svg?type=shield)](https://app.fossa.com/projects/custom%2B28176%2Fgithub.com%2Fkarmada-io%2Fkarmada?ref=badge_shield)

## Purpose

This library is the canonical location of the Karmada API definition.Most likely interaction with this repository is as a dependency of client-go.

## Recommended Use

We recommend using the go types in this repo. You may serialize them directly to JSON.

## Compatibility

Branches and tags track `karmada-io/karmada` branches and are compatible with that repo.

## Where does it come from?

This `api` is synced from [https://github.com/karmada-io/karmada/tree/master/pkg/apis](https://github.com/karmada-io/karmada/tree/master/pkg/apis).
Code changes are made in that location, merged into `karmada-io/karmada` and later synced here.

## Things you should NOT do

[https://github.com/karmada-io/karmada/tree/master/pkg/apis](https://github.com/karmada-io/karmada/tree/master/pkg/apis) is synced to here.
All changes must be made in the former. The latter is read-only.

## Governance

This repo is governed by the Karmada community. For details, see [Karmada Community Governance](https://github.com/karmada-io/community/blob/main/GOVERNANCE.md).

## Code of Conduct

Please refer to our [Karmada Community Code of Conduct](https://github.com/karmada-io/community/blob/main/CODE_OF_CONDUCT.md).

## Contributing

If you're interested in being a contributor and want to get involved in developing the Karmada code, please see [CONTRIBUTING](https://github.com/karmada-io/community/blob/main/CONTRIBUTING.md) for details on submitting patches and the contribution workflow.

## Security

Please refer to our [Karmada Security Policy](https://github.com/karmada-io/community/blob/main/security-team/SECURITY.md).

## Roadmap

The community defines a high level roadmap for Karmada development and upcoming releases. See the [Karmada Roadmap](https://github.com/karmada-io/karmada/blob/master/ROADMAP.md) file for details.
