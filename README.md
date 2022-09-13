# cloudbees-openfeature-provider-node

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
[![a](https://img.shields.io/badge/slack-%40cncf%2Fopenfeature-brightgreen?style=flat&logo=slack)](https://cloud-native.slack.com/archives/C0344AANLA1)
[![OpenFeature Specification](https://img.shields.io/static/v1?label=OpenFeature%20Specification&message=v0.4.0&color=yellow)](https://github.com/open-feature/spec/tree/v0.4.0)
[![OpenFeature SDK](https://img.shields.io/static/v1?label=OpenFeature%20Golang%20SDK&message=v0.2.0&color=green)](https://github.com/open-feature/golang-sdk/tree/v0.2.0)
[![Version](https://badge.fury.io/go/github.com%2Frollout%2Fcloudbees-openfeature-provider-go.svg)](https://github.com/rollout/cloudbees-openfeature-provider-go)
[![CloudBees Rox SDK](https://img.shields.io/static/v1?label=Rox%20SDK&message=v5.0.2&color=green)](https://github.com/rollout/rox-go)
[![Known Vulnerabilities](https://snyk.io/test/github/rollout/cloudbees-openfeature-provider-go/badge.svg)](https://snyk.io/test/github/rollout/cloudbees-openfeature-provider-go)

This is the [CloudBees](https://www.cloudbees.com/products/feature-management) provider implementation for [OpenFeature](https://openfeature.dev/) for the [Golang SDK](https://github.com/open-feature/golang-sdk).

OpenFeature provides a vendor-agnostic abstraction layer on Feature Flag management.

This provider allows the use of CloudBees Feature Management as a backend for Feature Flag configurations.

## Requirements
- go 17 or above

## Installation

### Add it to your build

```bash
go get github.com/rollout/cloudbees-openfeature-provider-go
```

### Configuration

Follow the instructions on the [Go SDK project](https://github.com/open-feature/golang-sdk) for how to use the Golang SDK.

You can configure the CloudBees provider by doing the following:

```go
package main

import (
	"github.com/open-feature/golang-sdk/pkg/openfeature"
	"github.com/rollout/cloudbees-openfeature-provider-go/cloudbees
)

func main() {
    appKey := "INSERT_APP_KEY_HERE"
	openfeature.SetProvider(cloudbees.NewProvider(appKey))
	client := openfeature.NewClient("app")
	value, err := client.BooleanValue("v2_enabled", false, openfeature.EvaluationContext{}, openfeature.EvaluationOptions{})
}```
