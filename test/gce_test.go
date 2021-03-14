package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestGCE(t *testing.T) {
	t.Parallel()
	actualGCEName := terraform.Output(t, terraformOptions, "gce_name")
	actualGCEMachineType := terraform.Output(t, terraformOptions, "gce_machine_type")
	actualGCEZone := terraform.Output(t, terraformOptions, "gce_zone")
	actualGCEDiskSize := terraform.Output(t, terraformOptions, "gce_boot_disk_size")
	actualGCEDiskImage := terraform.Output(t, terraformOptions, "gce_boot_disk_image")

	assert.Equal(t, actualGCEName, "sample")
	assert.Equal(t, actualGCEMachineType, "f1-micro")
	assert.Equal(t, actualGCEZone, "asia-northeast1-b")
	assert.Equal(t, actualGCEDiskSize, "20")
	assert.Contains(t, actualGCEDiskImage, "ubuntu-os-cloud/global/images/ubuntu-2004")
}
