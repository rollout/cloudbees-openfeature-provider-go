package cloudbees_test

import (
	"errors"
	"github.com/open-feature/golang-sdk/pkg/openfeature"
	"github.com/rollout/cloudbees-openfeature-provider-go/pkg/cloudbees"
	"reflect"
	"testing"
)

const dashboardAppKey = "62bee5bbca1059d18808adad"
const stringProperty = "stringproperty"
const numberProperty = "numberproperty"
const booleanProperty = "booleanproperty"
const badProperty = "not specified"

func TestProvider_Metadata(t *testing.T) {
	tests := map[string]struct {
		want openfeature.Metadata
	}{
		"Given a CloudBees provider, then Metadata() will return CloudbeesProvider": {
			want: openfeature.Metadata{Name: "CloudbeesProvider"},
		},
	}
	p, _ := cloudbees.NewProvider(dashboardAppKey)
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := p.Metadata(); got != tt.want {
				t.Errorf("metadata = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_Constructor(t *testing.T) {
	tests := map[string]struct {
		appKey string
		want   error
	}{
		"Given a badly formatted app key, then NewProvider will return an error": {
			appKey: "badlyFormattedApiKey",
			want:   errors.New("Invalid rollout apikey"),
		},
		"Given an empty app key, then NewProvider will return an error": {
			appKey: "",
			want:   errors.New("Invalid rollout apikey"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			provider, err := cloudbees.NewProvider(tt.appKey)
			if got := err; tt.want.Error() != got.Error() {
				t.Errorf("error = %v, want %v", got, tt.want)
			}
			if provider != nil {
				t.Error("provider should be nil")
			}
		})
	}
}

func TestProvider_BooleanEvaluation(t *testing.T) {
	type args struct {
		flag         string
		defaultValue bool
		evalCtx      map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want openfeature.BoolResolutionDetail
	}{
		{
			name: "targeting turned on, value set to true",
			args: args{
				flag:         "boolean-static-true",
				defaultValue: false,
			},
			want: openfeature.BoolResolutionDetail{
				Value: true,
			},
		},
		{
			name: "targeting turned on, value set to false",
			args: args{
				flag:         "boolean-static-false",
				defaultValue: true,
			},
			want: openfeature.BoolResolutionDetail{
				Value: false,
			},
		},
		{
			name: "targeting turned off, using default value of false",
			args: args{
				flag:         "boolean-disabled",
				defaultValue: false,
			},
			want: openfeature.BoolResolutionDetail{
				Value: false,
			},
		},
		{
			name: "targeting turned off, using default value of true",
			args: args{
				flag:         "boolean-disabled",
				defaultValue: true,
			},
			want: openfeature.BoolResolutionDetail{
				Value: true,
			},
		},
		{
			name: "using a context, property is 'on'",
			args: args{
				flag:         "boolean-with-context",
				defaultValue: false,
				evalCtx:      map[string]interface{}{stringProperty: "on"},
			},
			want: openfeature.BoolResolutionDetail{
				Value: true,
			},
		},
		{
			name: "using a context, property is 'off'",
			args: args{
				flag:         "boolean-with-context",
				defaultValue: false,
				evalCtx:      map[string]interface{}{stringProperty: "off"},
			},
			want: openfeature.BoolResolutionDetail{
				Value: false,
			},
		},
	}
	provider, err := cloudbees.NewProvider(dashboardAppKey)
	if err != nil {
		t.Errorf("creating CloudBees provider %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := provider.BooleanEvaluation(tt.args.flag, tt.args.defaultValue, tt.args.evalCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name: %v; BooleanEvaluation() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestProvider_StringEvaluation(t *testing.T) {
	type args struct {
		flag         string
		defaultValue string
		evalCtx      map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want openfeature.StringResolutionDetail
	}{
		{
			name: "targeting turned on, value set to yes",
			args: args{
				flag:         "string-static-yes",
				defaultValue: "default",
			},
			want: openfeature.StringResolutionDetail{
				Value: "yes",
			},
		},
		{
			name: "targeting turned on, value set to no",
			args: args{
				flag:         "string-static-no",
				defaultValue: "default",
			},
			want: openfeature.StringResolutionDetail{
				Value: "no",
			},
		},
		{
			name: "targeting turned off, using default value",
			args: args{
				flag:         "string-disabled",
				defaultValue: "banana",
			},
			want: openfeature.StringResolutionDetail{
				Value: "banana",
			},
		},
		{
			name: "using a context, property is on evaluates to yes",
			args: args{
				flag:         "string-with-context",
				defaultValue: "default",
				evalCtx:      map[string]interface{}{stringProperty: "on"},
			},
			want: openfeature.StringResolutionDetail{
				Value: "yes",
			},
		},
		{
			name: "using a context, property is off evaluates to no",
			args: args{
				flag:         "string-with-context",
				defaultValue: "default",
				evalCtx:      map[string]interface{}{stringProperty: "off"},
			},
			want: openfeature.StringResolutionDetail{
				Value: "no",
			},
		},
		{
			name: "using a context, property is not defined results in dashboard",
			args: args{
				flag:         "string-with-context",
				defaultValue: "default",
				evalCtx:      map[string]interface{}{badProperty: "whatever"},
			},
			want: openfeature.StringResolutionDetail{
				Value: "not specified",
			},
		},
		{
			name: "using a context, no property results in dashboard",
			args: args{
				flag:         "string-with-context",
				defaultValue: "default",
				evalCtx:      map[string]interface{}{},
			},
			want: openfeature.StringResolutionDetail{
				Value: "not specified",
			},
		},
		{
			name: "using a context, nil context results in dashboard",
			args: args{
				flag:         "string-with-context",
				defaultValue: "default",
				evalCtx:      nil,
			},
			want: openfeature.StringResolutionDetail{
				Value: "not specified",
			},
		},
		{
			name: "using a context, no context results in dashboard",
			args: args{
				flag:         "string-with-context",
				defaultValue: "default",
			},
			want: openfeature.StringResolutionDetail{
				Value: "not specified",
			},
		},
	}
	provider, err := cloudbees.NewProvider(dashboardAppKey)
	if err != nil {
		t.Errorf("creating CloudBees provider %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := provider.StringEvaluation(tt.args.flag, tt.args.defaultValue, tt.args.evalCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name: %v; BooleanEvaluation() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestProvider_FloatEvaluation(t *testing.T) {
	type args struct {
		flag         string
		defaultValue float64
		evalCtx      map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want openfeature.FloatResolutionDetail
	}{
		{
			name: "targeting turned on, value set to 5",
			args: args{
				flag:         "integer-static-5",
				defaultValue: 5.0,
			},
			want: openfeature.FloatResolutionDetail{
				Value: 5.0,
			},
		},
		{
			name: "targeting turned off, using default value",
			args: args{
				flag:         "integer-disabled",
				defaultValue: 7.0,
			},
			want: openfeature.FloatResolutionDetail{
				Value: 7.0,
			},
		},
		{
			name: "using a context, property is set to 1",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1.0,
				evalCtx:      map[string]interface{}{stringProperty: "1"},
			},
			want: openfeature.FloatResolutionDetail{
				Value: 1.0,
			},
		},
		{
			name: "using a context, property is set to 5",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1.0,
				evalCtx:      map[string]interface{}{stringProperty: "5"},
			},
			want: openfeature.FloatResolutionDetail{
				Value: 5.0,
			},
		},
		{
			name: "using a context, no property results in dashboard",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1.0,
				evalCtx:      map[string]interface{}{},
			},
			want: openfeature.FloatResolutionDetail{
				Value: 10.0,
			},
		},
		{
			name: "using a context, nil context results in dashboard",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1.0,
				evalCtx:      nil,
			},
			want: openfeature.FloatResolutionDetail{
				Value: 10.0,
			},
		},
		{
			name: "using a context, no context results in dashboard",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1.0,
			},
			want: openfeature.FloatResolutionDetail{
				Value: 10.0,
			},
		},
	}
	provider, err := cloudbees.NewProvider(dashboardAppKey)
	if err != nil {
		t.Errorf("creating CloudBees provider %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := provider.FloatEvaluation(tt.args.flag, tt.args.defaultValue, tt.args.evalCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name: %v; BooleanEvaluation() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestProvider_IntEvaluation(t *testing.T) {
	type args struct {
		flag         string
		defaultValue int64
		evalCtx      map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want openfeature.IntResolutionDetail
	}{
		{
			name: "targeting turned on, value set to 5",
			args: args{
				flag:         "integer-static-5",
				defaultValue: 5,
			},
			want: openfeature.IntResolutionDetail{
				Value: 5,
			},
		},
		{
			name: "targeting turned off, using default value",
			args: args{
				flag:         "integer-disabled",
				defaultValue: 7,
			},
			want: openfeature.IntResolutionDetail{
				Value: 7,
			},
		},
		{
			name: "using a context, property is set to 1",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{stringProperty: "1"},
			},
			want: openfeature.IntResolutionDetail{
				Value: 1,
			},
		},
		{
			name: "using a context, property is set to 5",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{stringProperty: "5"},
			},
			want: openfeature.IntResolutionDetail{
				Value: 5,
			},
		},
		{
			name: "using a context, no property results in dashboard",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{},
			},
			want: openfeature.IntResolutionDetail{
				Value: 10,
			},
		},
		{
			name: "using a context, nil context results in dashboard",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1,
				evalCtx:      nil,
			},
			want: openfeature.IntResolutionDetail{
				Value: 10,
			},
		},
		{
			name: "using a context, no context results in dashboard",
			args: args{
				flag:         "integer-with-context",
				defaultValue: -1,
			},
			want: openfeature.IntResolutionDetail{
				Value: 10,
			},
		},
	}
	provider, err := cloudbees.NewProvider(dashboardAppKey)
	if err != nil {
		t.Errorf("creating CloudBees provider %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := provider.IntEvaluation(tt.args.flag, tt.args.defaultValue, tt.args.evalCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name: %v; BooleanEvaluation() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestProvider_ObjectEvaluation(t *testing.T) {
	type args struct {
		flag         string
		defaultValue interface{}
	}
	tests := []struct {
		name string
		args args
		want openfeature.ResolutionDetail
	}{
		{
			name: "not supported",
			args: args{
				flag:         "whatever",
				defaultValue: "default",
			},
			want: openfeature.ResolutionDetail{
				Value:     "default",
				ErrorCode: "Not implemented - CloudBees feature management does not support an object type. Only String, Number and Boolean",
				Reason:    openfeature.ERROR,
			},
		},
	}
	provider, err := cloudbees.NewProvider(dashboardAppKey)
	if err != nil {
		t.Errorf("creating CloudBees provider %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := provider.ObjectEvaluation(tt.args.flag, tt.args.defaultValue, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name: %v; BooleanEvaluation() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestProvider_IntEvaluation_WithDifferentlyTypedContextObjects(t *testing.T) {
	type args struct {
		flag         string
		defaultValue int64
		evalCtx      map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want openfeature.IntResolutionDetail
	}{
		// Test positive matches for supported types (string/number/boolean)
		{
			name: "stringproperty of one resolves to 1",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{stringProperty: "one"},
			},
			want: openfeature.IntResolutionDetail{
				Value: 1,
			},
		},
		{
			name: "numberproperty of 1 resolves to 1",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{numberProperty: 1},
			},
			want: openfeature.IntResolutionDetail{
				Value: 1,
			},
		},
		{
			name: "numberproperty of 1.0 resolves to 1",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{numberProperty: 1.0},
			},
			want: openfeature.IntResolutionDetail{
				Value: 1,
			},
		},
		{
			name: "booleanproperty of true resolves to 1",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{booleanProperty: true},
			},
			want: openfeature.IntResolutionDetail{
				Value: 1,
			},
		},

		// Test negative matches for supported types (string/number/boolean) - it should serve the default value
		{
			name: "stringproperty of no uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{stringProperty: "no"},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1,
			},
		},
		{
			name: "numberproperty of 0 uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{numberProperty: 0},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1,
			},
		},
		{
			name: "numberproperty of 0.0 uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{numberProperty: 0.0},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1,
			},
		},
		{
			name: "booleanproperty of false uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{booleanProperty: false},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1,
			},
		},

		// Unexpected/unsupported contexts
		{
			name: "badproperty uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{badProperty: "whatever"},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1, // default
			},
		},
		{
			name: "stringproperty as a list (slice) uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{stringProperty: make([]int, 0)},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1, // default
			},
		},
		{
			name: "stringproperty as a map uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{stringProperty: make(map[string]int)},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1, // default
			},
		},
		{
			name: "stringproperty as a number uses default value",
			args: args{
				flag:         "integer-with-complex-context",
				defaultValue: -1,
				evalCtx:      map[string]interface{}{stringProperty: 1},
			},
			want: openfeature.IntResolutionDetail{
				Value: -1, // default
			},
		},
	}
	provider, err := cloudbees.NewProvider(dashboardAppKey)
	if err != nil {
		t.Errorf("creating CloudBees provider %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := provider.IntEvaluation(tt.args.flag, tt.args.defaultValue, tt.args.evalCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name: %v; BooleanEvaluation() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
