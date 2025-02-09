// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import (
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/baidu"
)

func init() {
	type VpcListOptions struct {
	}
	shellutils.R(&VpcListOptions{}, "vpc-list", "list vpcs", func(cli *baidu.SRegion, args *VpcListOptions) error {
		vpcs, err := cli.GetVpcs()
		if err != nil {
			return err
		}
		printList(vpcs)
		return nil
	})
	type VpcIdOptions struct {
		ID string `help:"ID of vpc to show"`
	}
	shellutils.R(&VpcIdOptions{}, "vpc-show", "list vpcs", func(cli *baidu.SRegion, args *VpcIdOptions) error {
		vpc, err := cli.GetVpc(args.ID)
		if err != nil {
			return errors.Wrap(err, "GetVpc")
		}
		printObject(vpc)
		return nil
	})

	shellutils.R(&VpcIdOptions{}, "vpc-delete", "delete vpc", func(cli *baidu.SRegion, args *VpcIdOptions) error {
		return cli.DeleteVpc(args.ID)
	})

	shellutils.R(&cloudprovider.VpcCreateOptions{}, "vpc-create", "create vpc", func(cli *baidu.SRegion, args *cloudprovider.VpcCreateOptions) error {
		vpc, err := cli.CreateVpc(args)
		if err != nil {
			return err
		}
		printObject(vpc)
		return nil
	})

}
