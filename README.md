# cloudbees-openfeature-provider-go

[![a](https://img.shields.io/badge/slack-%40cncf%2Fopenfeature-brightgreen?style=flat&logo=slack)](https://cloud-native.slack.com/archives/C0344AANLA1)
[![OpenFeature Specification](https://img.shields.io/static/v1?label=OpenFeature%20Specification&message=v0.4.0&color=yellow)](https://github.com/open-feature/spec/tree/v0.4.0)
[![OpenFeature SDK](https://img.shields.io/static/v1?label=OpenFeature%20Golang%20SDK&message=v0.4.0&color=green)](https://github.com/open-feature/go-sdk)
[![CloudBees Rox SDK](https://img.shields.io/static/v1?label=Rox%20SDK&message=v5.0.2&color=green)](https://github.com/rollout/rox-go)

This is the [CloudBees](https://www.cloudbees.com/products/feature-management) provider implementation for [OpenFeature](https://openfeature.dev/) for the [Go SDK](https://github.com/open-feature/go-sdk).

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

Follow the instructions on the [Go SDK project](https://github.com/open-feature/go-sdk) for how to use the Go SDK.

You can configure the CloudBees provider by doing the following:

```go
package main

import (
	"github.com/open-feature/go-sdk/pkg/openfeature"
	"github.com/rollout/cloudbees-openfeature-provider-go/pkg/cloudbees"
)

func main() {
	appKey := "INSERT_APP_KEY_HERE"
	if provider, err := cloudbees.NewProvider(appKey); err == nil {
		openfeature.SetProvider(provider)
		client := openfeature.NewClient("app")
		value, err := client.BooleanValue("enableTutorial", false, openfeature.EvaluationContext{}, openfeature.EvaluationOptions{})
		fmt.Printf("flag value: %v, error: %v", value, err)
	} else {
		fmt.Printf("error creating client %v", err)
	}
}
```
