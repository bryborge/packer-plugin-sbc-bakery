package builder

// A custom builder must implement the Builder interface.
//
// type Builder interface {
//   ConfigSpec() hcldec.ObjectSpec
//   Prepare(...interface{}) ([]string, []string, error)
//   Run(context.Context, ui Ui, hook Hook) (Artifact, error)
// }
//
// Docs: https://developer.hashicorp.com/packer/docs/plugins/creation/custom-builders#the-interface

import (
	"context"

	cfg "github.com/bryborge/sbc-bakery/config"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig  `mapstructure:",squash"`
	cfg.RemoteFileConfig `mapstructure:",squash"`

	ctx interpolate.Context
}

type Builder struct {
	config  Config
	context context.Context
	cancel  context.CancelFunc

	// TODO: Actually do something ...
	// runner multistep.Runner
}

// Returns a new instance of the builder.
func NewBuilder() *Builder {
	ctx, cancel := context.WithCancel(context.Background())
	return &Builder{
		context: ctx,
		cancel:  cancel,
	}
}

// Returns the configuration specification for the builder.
// Docs: https://developer.hashicorp.com/packer/docs/plugins/creation/custom-builders#the-configspec-method
func (b *Builder) ConfigSpec() hcldec.ObjectSpec {
	return b.config.FlatMapstructure().HCL2Spec()
}

func (b *Builder) InitConfig(ctx *interpolate.Context) (warnings []string, errors []error) {
	var (
		warns []string
		errs  []error
	)

	warns, errs = b.config.RemoteFileConfig.Prepare(ctx)
	warnings = append(warnings, warns...)
	errors = append(errors, errs...)

	return warnings, errors
}

// Processes the input configuration and returns any warnings and errors.
// Docs: https://developer.hashicorp.com/packer/docs/plugins/creation/custom-builders#the-prepare-method
func (b *Builder) Prepare(args ...interface{}) ([]string, []string, error) {
	var (
		errs     *packer.MultiError
		warnings []string
	)

	if err := config.Decode(&b.config, &config.DecodeOpts{
		Interpolate:       true,
		InterpolateFilter: &interpolate.RenderFilter{},
	}, args...); err != nil {
		return nil, nil, err
	}

	fileWarns, fileErrs := b.InitConfig(&b.config.ctx)
	warnings = append(fileWarns, fileWarns...)
	errs = packer.MultiErrorAppend(errs, fileErrs...)

	if errs != nil && len(errs.Errors) > 0 {
		return nil, warnings, errs
	}

	return nil, warnings, nil
}

// Executes the build and returns an artifact.
// Docs: https://developer.hashicorp.com/packer/docs/plugins/creation/custom-builders#the-run-method
func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	// TODO: Actually do something ...
	return nil, nil
}
