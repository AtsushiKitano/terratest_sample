package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestConnection(t *testing.T) {
	t.Parallel()
	ip := terraform.OutputList(t, terraformOptions, "gce_global_ip")[0]
	url := "http://" + ip + ":80"

	statusCode, _ := http_helper.HTTPDo(t, "GET", url, nil, nil, nil)
	expectedCode := 200

	if statusCode != expectedCode {
		t.Errorf("handler returned wrong status code: got %v want %v", statusCode, expectedCode)
	}
}
