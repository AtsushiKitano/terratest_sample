package test

import (
    "github.com/gruntwork-io/terratest/modules/terraform"
)

var terraformOptions = &terraform.Options{
	TerraformDir: "../src/",
}
