package ssm

import (
	"github.com/concourse/concourse/atc/creds"
	"github.com/go-viper/mapstructure/v2"
	flags "github.com/jessevdk/go-flags"
)

type ssmManagerFactory struct{}

func init() {
	creds.Register("ssm", NewSsmManagerFactory())
}

func NewSsmManagerFactory() creds.ManagerFactory {
	return &ssmManagerFactory{}
}

func (factory *ssmManagerFactory) AddConfig(group *flags.Group) creds.Manager {
	manager := &SsmManager{}
	subGroup, err := group.AddGroup("AWS SSM Credential Management", "", manager)
	if err != nil {
		panic(err)
	}

	subGroup.Namespace = "aws-ssm"
	return manager
}

func (factory *ssmManagerFactory) NewInstance(config any) (creds.Manager, error) {
	manager := &SsmManager{
		TeamSecretTemplate:     DefaultTeamSecretTemplate,
		PipelineSecretTemplate: DefaultPipelineSecretTemplate,
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &manager,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return manager, nil
}
