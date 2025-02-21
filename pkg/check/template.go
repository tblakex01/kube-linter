package check

import (
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
)

// A Func is a specific lint-check, which runs on a specific objects, and emits diagnostics if problems are found.
// Checks have access to the entire LintContext, with all the objects in it, but must only report problems for the
// object passed in the second argument.
type Func func(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic

// A Template is a template for a check.
type Template struct {
	// HumanName is a human-friendly name for the template.
	// It is to be used ONLY for documentation, and has no
	// semantic relevance.
	HumanName            string
	Key                  string
	Description          string
	SupportedObjectKinds config.ObjectKindsDesc

	Parameters             []ParameterDesc
	ParseAndValidateParams func(params map[string]interface{}) (interface{}, error)
	Instantiate            func(parsedParams interface{}) (Func, error)
}
