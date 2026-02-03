package nats

import (
	ansiblecfg "github.com/EigerCode/openuem-ansible-config/ansible"
	"github.com/EigerCode/wingetcfg/wingetcfg"
)

type CfgProfiles struct {
	AgentID string `json:"agentID,omitempty"`
}

type ProfileConfig struct {
	ProfileID     int                           `yaml:"profileID"`
	Exclusions    []string                      `yaml:"exclusions"`
	Deployments   []string                      `yaml:"deployments"`
	WinGetConfig  *wingetcfg.WinGetCfg          `yaml:"config,omitempty"`
	AnsibleConfig []*ansiblecfg.AnsiblePlaybook `yaml:"ansible,omitempty"`
	NetBirdConfig []*NetbirdTask                `yaml:"netbird,omitempty"`
}
