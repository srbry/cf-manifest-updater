# cf-manifest-updater

[![codecov.io](https://codecov.io/github/srbry/cf-manifest-updater/coverage.svg?branch=master)](https://codecov.io/github/srbry/cf-manifest-updater?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/srbry/cf-manifest-updater)](https://goreportcard.com/report/github.com/srbry/cf-manifest-updater)
[![Build Status](https://travis-ci.org/srbry/cf-manifest-updater.svg?branch=master)](https://travis-ci.org/srbry/cf-manifest-updater)

`cf-manifest-updater` is a simple CLI tool for upgrading your CF manifest to help with deprecations as described [here](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html#route-attribute).

As other features of manifests get altered this tool will aim to provide easy migration paths where possible.

## Install

`go get github.com/srbry/cf-manifest-updater`

## Usage

```sh
cf-manifest-updater <manifest_file>
```
