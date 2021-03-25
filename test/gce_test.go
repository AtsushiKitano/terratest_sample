package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestGCE(t *testing.T) {
	t.Parallel()

	actualGCEName := terraform.Output(t, terraformOptions, "gce_name")
	assert.Equal(t, "sample", actualGCEName)

	actualGCEMachineType := terraform.Output(t, terraformOptions, "gce_machine_type")
	assert.Equal(t, "f1-micro", actualGCEMachineType)

	actualGCEZone := terraform.Output(t, terraformOptions, "gce_zone")
	assert.Equal(t, "asia-northeast1-b", actualGCEZone)

	actualGCEDiskSize := terraform.Output(t, terraformOptions, "gce_boot_disk_size")
	assert.Equal(t, "20", actualGCEDiskSize)

	actualGCEDiskImage := terraform.Output(t, terraformOptions, "gce_boot_disk_image")
	assert.Contains(t, "ubuntu-os-cloud/ubuntu-2004-lts", actualGCEDiskImage)
}
