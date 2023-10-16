package testconfig

import (
	"encoding/base64"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/networks"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

const (
	DefaultConfigFilename = "config.toml"

	ErrReadConfig           = "failed to read TOML config"
	ErrUnmarshalConfig      = "failed to unmarshal TOML config"
	ErrInvalidOverrideGroup = "invalid override key, please check ApplyOverrides() in config and override with env vars"
)

type TestConfig interface {
	ApplyOverrides(from *Config) error
	ReadSecrets() error
	Validate() error
}

type Config struct {
	CCIP ProductTestConfig
}

type ProductTestConfig struct {
	EnvUser   string                `toml:",omitempty"`
	TTL       models.Duration       `toml:",omitempty"`
	Chainlink *Chainlink            `toml:",omitempty"`
	Networks  []string              `toml:",omitempty"`
	Tests     map[string]TestConfig `toml:",omitempty"`
}

func (c *Config) readSecrets() error {
	// read secrets for all products and test types

	return nil
}

func (c *Config) EVMNetworks() []blockchain.EVMNetwork {
	return networks.SetNetworks(c.Networks)
}

func (c *Config) ApplyOverrides(from *Config) error {

	return nil
}

type Chainlink struct {
	Common    *Node   `toml:",omitempty"`
	NoOfNodes *int    `toml:",omitempty"`
	Nodes     []*Node `toml:",omitempty"` // to be mentioned only if diff nodes follow diff configs; not required if all nodes follow CommonConfig
}

func (c *Chainlink) ApplyOverrides(from *Chainlink) {
	if from.NoOfNodes != nil {
		c.NoOfNodes = from.NoOfNodes
	}
	if from.Common != nil {
		c.Common.ApplyOverrides(from.Common)
	}
}

type Node struct {
	Name       *string `toml:",omitempty"`
	Image      *string `toml:",omitempty"`
	Tag        *string `toml:",omitempty"`
	NodeConfig *string `toml:",omitempty"`
	DBImage    *string `toml:",omitempty"`
	DBTag      *string `toml:",omitempty"`
}

func (n *Node) ApplyOverrides(from *Node) {
	if from.Name != nil {
		n.Name = from.Name
	}
	if from.Image != nil {
		n.Image = from.Image
	}
	if from.Tag != nil {
		n.Tag = from.Tag
	}
	if from.DBImage != nil {
		n.DBImage = from.DBImage
	}
	if from.DBTag != nil {
		n.DBTag = from.DBTag
	}
	if from.NodeConfig != nil {
		n.NodeConfig = from.NodeConfig
	}
}

func NewConfig() (*Config, error) {
	var cfg *Config
	var override *Config
	// load config from default file
	d, err := os.ReadFile(DefaultConfigFilename)
	if err != nil {
		return nil, errors.Wrap(err, ErrReadConfig)
	}
	err = toml.Unmarshal(d, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, ErrUnmarshalConfig)
	}
	// load config from env var if specified
	rawConfig, _ := utils.GetEnv("TEST_CONFIG_OVERRIDE")
	if rawConfig != "" {
		d, err = base64.StdEncoding.DecodeString(rawConfig)
		err = toml.Unmarshal(d, &override)
		if err != nil {
			return nil, errors.Wrap(err, ErrUnmarshalConfig)
		}
	}
	if override != nil {
		if err := cfg.ApplyOverrides(override); err != nil {
			return cfg, err
		}
	}
	if err := cfg.readSecrets(); err != nil {
		return nil, err
	}
	log.Debug().Interface("Config", cfg).Msg("Parsed config")
	return cfg, nil
}
