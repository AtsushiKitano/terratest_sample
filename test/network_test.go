package test

import (
    "testing"

    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/gruntwork-io/terratest/modules/http-helper"
    "github.com/stretchr/testify/assert"
)

var terraformOptions = &terraform.Options{
    TerraformDir: "../src/",
}

func TestNetwork(t *testing.T) {
    t.Parallel()
    testNetworkName := terraform.Output(t, terraformOptions, "network")
    assert.Equal(t, "sample", testNetworkName)
}

func TestConnection(t *testing.T) {
    t.Parallel()
    ip := terraform.Output(t, terraformOptions, "gce_static_ip")
    url := "http://" + ip + ":80"

    statusCode, _ := http_helper.HTTPDo(t, "GET", url, nil, nil, nil)
    expectedCode := 200

    if statusCode != expectedCode {
        t.Errorf("handler returned wrong status code: got %v want %v", statusCode, expectedCode)
    }
}
