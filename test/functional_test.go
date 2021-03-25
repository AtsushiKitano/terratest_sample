package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestConnection(t *testing.T) {
	t.Parallel()

	ipAddressList := terraform.OutputList(t, terraformOptions, "gce_global_ip")
	url := fmt.Sprintf("http://%s:80", ipAddressList[0])

	statusCode, _ := http_helper.HTTPDo(t, "GET", url, nil, nil, nil)
	if expectedCode := 200; statusCode != expectedCode {
		t.Errorf("handler returned wrong status code: got %v want %v", statusCode, expectedCode)
	}
}
