package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformMasterBranch(t *testing.T) {
    terraformOptions := &terraform.Options{
        // Path to the Terraform code that will be tested
        TerraformDir: "../",
        
        // Variables to pass to the Terraform code using -var options
        Vars: map[string]interface{}{
            "example_variable": "example_value",
        },
        
        // Disable colors in Terraform commands so its easier to parse stdout/stderr
        NoColor: true,
    }

    // Clean up resources with "terraform destroy" at the end of the test
    defer terraform.Destroy(t, terraformOptions)

    // Run "terraform init" and "terraform apply". Fail the test if there are any errors.
    terraform.InitAndApply(t, terraformOptions)

    // Run `terraform output` to get the value of an output variable
    output := terraform.Output(t, terraformOptions, "example_output")

    // Verify we're getting back the output we expect
    assert.Equal(t, "expected_output_value", output)
}
