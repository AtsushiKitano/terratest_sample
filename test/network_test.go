package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNetwork(t *testing.T) {
	t.Parallel()

	actualNetworkName := terraform.Output(t, terraformOptions, "network_name")
	assert.Equal(t, "sample", actualNetworkName)

	actualSubnetName := terraform.Output(t, terraformOptions, "subnetwork_name")
	assert.Equal(t, "sample", actualSubnetName)

	actualSubnetRegion := terraform.Output(t, terraformOptions, "subnetwork_region")
	assert.Equal(t, "asia-northeast1", actualSubnetRegion)

	actualSubnetCidr := terraform.Output(t, terraformOptions, "subnetwork_cidr")
	assert.Equal(t, "192.168.10.0/24", actualSubnetCidr)
}

func TestFirewall(t *testing.T) {
	t.Parallel()

	actualFirewallName := terraform.Output(t, terraformOptions, "firewall_name")
	assert.Equal(t, "ingress-sample", actualFirewallName)

	actualFirewallDirection := terraform.Output(t, terraformOptions, "firewall_direction")
	assert.Equal(t, "INGRESS", actualFirewallDirection)

	actualFirewallPriority := terraform.Output(t, terraformOptions, "firewall_priority")
	assert.Equal(t, "1000", actualFirewallPriority)

	actualFirewallProtocol := terraform.Output(t, terraformOptions, "firewall_allow_rules_protocol")
	assert.Equal(t, "tcp", actualFirewallProtocol)

	expectedSrcRanges := []string{"0.0.0.0/0"}
	actualFirewallSrcRanges := terraform.OutputList(t, terraformOptions, "firewall_src_ranges")
	assert.Equal(t, expectedSrcRanges, actualFirewallSrcRanges)

	expectedFirewallPorts := []string{"80"}
	actualFirewallPorts := terraform.OutputList(t, terraformOptions, "firewall_allow_rules_ports")
	assert.Equal(t, expectedFirewallPorts, actualFirewallPorts)
}
