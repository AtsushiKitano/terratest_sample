package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNetwork(t *testing.T) {
	t.Parallel()
	actualNetworkName := terraform.Output(t, terraformOptions, "network_name")
    actualSubnetName := terraform.Output(t, terraformOptions, "subnetwork_name")
    actualSubnetRegion := terraform.Output(t, terraformOptions, "subnetwork_region")
    actualSubnetCidr := terraform.Output(t, terraformOptions, "subnetwork_cidr")

	assert.Equal(t, actualNetworkName, "sample")
    assert.Equal(t, actualSubnetName, "sample")
    assert.Equal(t, actualSubnetRegion, "asia-northeast1")
    assert.Equal(t, actualSubnetCidr, "192.168.10.0/24")
}

func TestFirewall(t *testing.T) {
	t.Parallel()
	actualFirewallName := terraform.Output(t, terraformOptions, "firewall_name")
	actualFirewallDirection := terraform.Output(t, terraformOptions, "firewall_direction")
	actualFirewallPriority := terraform.Output(t, terraformOptions, "firewall_priority")
	actualFirewallProtocol := terraform.Output(t, terraformOptions, "firewall_allow_rules_protocol")
	actualFirewallPorts := terraform.OutputList(t, terraformOptions, "firewall_allow_rules_ports")
	actualFirewallSrcRanges := terraform.OutputList(t, terraformOptions, "firewall_src_ranges")

	expected_src_ranges := []string{"0.0.0.0/0"}
	expected_firewall_ports := []string{"80"}

	assert.Equal(t, actualFirewallName, "ingress-sample")
	assert.Equal(t, actualFirewallDirection, "INGRESS")
	assert.Equal(t, actualFirewallSrcRanges, expected_src_ranges)
	assert.Equal(t, actualFirewallPorts, expected_firewall_ports)
	assert.Equal(t, actualFirewallPriority, "1000")
	assert.Equal(t, actualFirewallProtocol, "tcp")
}
