# api

Schema of the API types that are served by Karmada.

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
