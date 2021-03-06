// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package glacier

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/private/waiter"
)

func (c *Glacier) WaitUntilVaultExists(input *DescribeVaultInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVault",
		Delay:       3,
		MaxAttempts: 15,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "status",
				Argument: "",
				Expected: 200,
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "ResourceNotFoundException",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *Glacier) WaitUntilVaultNotExists(input *DescribeVaultInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVault",
		Delay:       3,
		MaxAttempts: 15,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "retry",
				Matcher:  "status",
				Argument: "",
				Expected: 200,
			},
			{
				State:    "success",
				Matcher:  "error",
				Argument: "",
				Expected: "ResourceNotFoundException",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
