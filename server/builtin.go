package server

import (
	"github.com/avenbreaks/raindrops/chain"
	"github.com/avenbreaks/raindrops/consensus"
	consensusDev "github.com/avenbreaks/raindrops/consensus/dev"
	consensusDummy "github.com/avenbreaks/raindrops/consensus/dummy"
	consensusIBFT "github.com/avenbreaks/raindrops/consensus/ibft"
	consensusPolyBFT "github.com/avenbreaks/raindrops/consensus/polybft"
	"github.com/avenbreaks/raindrops/forkmanager"
	"github.com/avenbreaks/raindrops/secrets"
	"github.com/avenbreaks/raindrops/secrets/awsssm"
	"github.com/avenbreaks/raindrops/secrets/gcpssm"
	"github.com/avenbreaks/raindrops/secrets/hashicorpvault"
	"github.com/avenbreaks/raindrops/secrets/local"
	"github.com/avenbreaks/raindrops/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

var forkManagerFactory = map[ConsensusType]ForkManagerFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerFactory,
}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
