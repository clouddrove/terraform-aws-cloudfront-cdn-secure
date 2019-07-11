// Managed By : CloudDrove
// Description : This Terratest is used to test the Terraform VPC module.
// Copyright @ CloudDrove. All Right Reserved.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestCDN(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// Source path of Terraform directory.
		TerraformDir: "../../_example/domain-cdn",
	}

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	arn := terraform.OutputList(t, terraformOptions, "arn")
	// Verify we're getting back the outputs we expect
	assert.Contains(t, arn, "arn:aws:cloudfront")

}
