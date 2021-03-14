package helloworld

import(
    "testing"

    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
    t.Parallel()

    terraformOption := &terraform.Options{
        TerraformDir: ".",
    }

    actual := terraform.Output(t, terraformOption, "helloworld")
    expected := "Hello World"

    assert.Equal(t, expected, actual)
}
