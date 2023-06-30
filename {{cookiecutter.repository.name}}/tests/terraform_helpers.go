package tests

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

var (
	AdditionalRetryableTerraformErrors = map[string]string{}
)

// An ExampleConfiguration is the runtime lifecycle of an example Terraform
// configuration for this Terraform module, and should manage the init, apply,
// and destroy stages of the configuration, as well as accessing the Outputs
type ExampleConfiguration struct {
	dir  string `type:"string"`
	tag  string `type:"string"`
	envs map[string]string
	vars map[string]interface{}

	applyOptions *terraform.Options
	basicOptions *terraform.Options
}

// Create an ExampleConfiguration object which is used to manage the state and
// interactions with the example configuration being tested
func NewExampleConfiguration(t *testing.T, dir string, tag string) *ExampleConfiguration {
	cfg := &ExampleConfiguration{
		dir:  dir,
		tag:  tag,
		envs: map[string]string{},
		vars: map[string]interface{}{},
	}

	cfg.prepare(t)
	return cfg
}

// Update the working directory for the example Terraform configuration
// Note: This should not be done once Init() has been run, as you will need to
// re-Init() the configuration, and will loose access to all previous state
func (e *ExampleConfiguration) SetDir(t *testing.T, dir string) {
	e.dir = dir
	e.prepare(t)
}

// Update the tag for the example Terraform configuration
// Note: This should not be done once Apply() has been run, as it will force the
// resources to be re-created upon another Apply() (although it will have no
// effect on a Destroy() call)
func (e *ExampleConfiguration) SetTag(t *testing.T, tag string) {
	e.tag = tag
	e.prepare(t)
}

// Add an Environment Variable to the Terraform runtime when interacting with
// the example Terraform configuration
func (e *ExampleConfiguration) AddEnvironmentVariable(t *testing.T, name string, value string) {
	e.envs[name] = value
	e.prepare(t)
}

// Add a Terraform Variable via the command-line to the Terraform runtime when
// interacting with the example Terraform configuration
func (e *ExampleConfiguration) AddTerraformVariable(t *testing.T, name string, value string) {
	e.vars[name] = value
	e.prepare(t)
}

// Take the directory, tag, and variables provided, and set up the Terraform
// environment for the respective runtime types
func (e *ExampleConfiguration) prepare(t *testing.T) {
	basicOptions := &terraform.Options{
		TerraformDir:             e.dir,
		MaxRetries:               3,
		TimeBetweenRetries:       10 * time.Second,
		Lock:                     true,
		EnvVars:                  e.envs,
		RetryableTerraformErrors: AdditionalRetryableTerraformErrors,
	}

	// Take a copy of basicOptions and add the Vars configuration; this has to be
	// done separately or TerraTest will try to add -vars arguments to init and
	// output commands, which will throw errors, as such these should only be used
	// on apply and destroy commands
	applyOptions, err := basicOptions.Clone()
	require.NoError(t, err)
	e.vars["tag"] = e.tag // Make sure this is set/updated beforehand
	applyOptions.Vars = e.vars

	e.basicOptions = basicOptions
	e.applyOptions = applyOptions
}

// Initialise the example Terraform configuration
func (e ExampleConfiguration) Init(t *testing.T) {
	terraform.Init(t, terraform.WithDefaultRetryableErrors(t, e.basicOptions))
}

// Apply the example Terraform configuration, and then check for idempotence in
// the configuration by re-running the apply command
func (e ExampleConfiguration) Apply(t *testing.T) {
	terraform.ApplyAndIdempotent(t, terraform.WithDefaultRetryableErrors(t, e.applyOptions))
}

// Destroy the applied configuration for this Terraform example
func (e ExampleConfiguration) Destroy(t *testing.T) {
	terraform.Destroy(t, terraform.WithDefaultRetryableErrors(t, e.applyOptions))
}

// Fetch the Output called output from the applied configuration for this
// Terraform example
func (e ExampleConfiguration) Output(t *testing.T, output string) string {
	return terraform.Output(t, e.basicOptions, output)
}
