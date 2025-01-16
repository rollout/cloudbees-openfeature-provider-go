package cloudbees

import (
	"context"

	"github.com/open-feature/go-sdk/pkg/openfeature"
	roxcontext "github.com/rollout/rox-go/v6/core/context"
	"github.com/rollout/rox-go/v6/core/model"
	"github.com/rollout/rox-go/v6/server"
)

const providerName = "CloudBeesProvider"

// Provider implements the FeatureProvider interface and provides functions for evaluating flags using CloudBees Feature Management
type Provider struct {
	rox *server.Rox
}

// NewProvider creates a new Provider with the specified appKey
func NewProvider(appKey string) (*Provider, error) {
	return NewProviderWithOptions(appKey, server.NewRoxOptions(server.RoxOptionsBuilder{}))
}

// NewProviderWithOptions creates a new Provider specified appKey and RoxOption
func NewProviderWithOptions(appKey string, options model.RoxOptions) (*Provider, error) {
	rox := server.NewRox()
	err := <-rox.Setup(appKey, options)
	if err != nil {
		return nil, err
	}
	return &Provider{
		rox: rox,
	}, nil
}

func (p Provider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{Name: providerName}
}

// BooleanEvaluation returns a boolean flag.
func (p Provider) BooleanEvaluation(_ context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	value := p.rox.DynamicAPI().IsEnabled(flag, defaultValue, roxcontext.NewContext(evalCtx))
	return openfeature.BoolResolutionDetail{
		Value: value,
	}
}

// StringEvaluation returns a string flag.
func (p Provider) StringEvaluation(_ context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	value := p.rox.DynamicAPI().Value(flag, defaultValue, []string{}, roxcontext.NewContext(evalCtx))
	return openfeature.StringResolutionDetail{
		Value: value,
	}
}

// FloatEvaluation returns a float flag.
func (p Provider) FloatEvaluation(_ context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	value := p.rox.DynamicAPI().GetDouble(flag, defaultValue, []float64{}, roxcontext.NewContext(evalCtx))
	return openfeature.FloatResolutionDetail{
		Value: value,
	}
}

// IntEvaluation returns an int flag.
func (p Provider) IntEvaluation(_ context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	value := int64(p.rox.DynamicAPI().GetInt(flag, int(defaultValue), []int{}, roxcontext.NewContext(evalCtx)))
	return openfeature.IntResolutionDetail{
		Value: value,
	}
}

// ObjectEvaluation returns an object flag
func (p Provider) ObjectEvaluation(_ context.Context, _ string, defaultValue interface{}, _ openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewInvalidContextResolutionError("Not implemented - CloudBees feature management does not support an object type. Only String, Number and Boolean"),
		},
	}
}

// Hooks returns hooks
func (p Provider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}
