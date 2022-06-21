package networkfirewall_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/networkfirewall"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccNetworkFirewallFirewallDataSource_basic(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_networkfirewall_firewall.test"
	dataSourceName := "data.aws_networkfirewall_firewall.test"
	policyResourceName := "aws_networkfirewall_firewall_policy.test"
	subnetResourceName := "aws_subnet.test"
	vpcResourceName := "aws_vpc.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t); testAccPreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, networkfirewall.EndpointsID),
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFirewallExists(resourceName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "network-firewall", fmt.Sprintf("firewall/%s", rName)),
					resource.TestCheckResourceAttr(dataSourceName, "delete_protection", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "description", ""),
					resource.TestCheckResourceAttrPair(dataSourceName, "firewall_policy_arn", policyResourceName, "arn"),
					resource.TestCheckResourceAttr(dataSourceName, "firewall_status.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "firewall_status.0.sync_states.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "firewall_status.0.sync_states.*.availability_zone", subnetResourceName, "availability_zone"),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceName, "firewall_status.0.sync_states.*", map[string]*regexp.Regexp{
						"attachment.0.endpoint_id": regexp.MustCompile(`vpce-`),
					}),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "firewall_status.0.sync_states.*.attachment.0.subnet_id", subnetResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttrPair(dataSourceName, "vpc_id", vpcResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "subnet_mapping.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "subnet_mapping.*.subnet_id", subnetResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "update_token"),
				),
			},
		},
	})
}

func TestAccNetworkFirewallFirewallDataSource_name(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_networkfirewall_firewall.test"
	dataSourceName := "data.aws_networkfirewall_firewall.test"
	policyResourceName := "aws_networkfirewall_firewall_policy.test"
	subnetResourceName := "aws_subnet.test"
	vpcResourceName := "aws_vpc.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t); testAccPreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, networkfirewall.EndpointsID),
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallDataSourceConfig_name(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFirewallExists(resourceName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "network-firewall", fmt.Sprintf("firewall/%s", rName)),
					resource.TestCheckResourceAttr(dataSourceName, "delete_protection", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "description", ""),
					resource.TestCheckResourceAttrPair(dataSourceName, "firewall_policy_arn", policyResourceName, "arn"),
					resource.TestCheckResourceAttr(dataSourceName, "firewall_status.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "firewall_status.0.sync_states.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "firewall_status.0.sync_states.*.availability_zone", subnetResourceName, "availability_zone"),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceName, "firewall_status.0.sync_states.*", map[string]*regexp.Regexp{
						"attachment.0.endpoint_id": regexp.MustCompile(`vpce-`),
					}),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "firewall_status.0.sync_states.*.attachment.0.subnet_id", subnetResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttrPair(dataSourceName, "vpc_id", vpcResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "subnet_mapping.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "subnet_mapping.*.subnet_id", subnetResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "update_token"),
				),
			},
		},
	})
}

func TestAccNetworkFirewallFirewallDataSource_both(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_networkfirewall_firewall.test"
	dataSourceName := "data.aws_networkfirewall_firewall.test"
	policyResourceName := "aws_networkfirewall_firewall_policy.test"
	subnetResourceName := "aws_subnet.test"
	vpcResourceName := "aws_vpc.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t); testAccPreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, networkfirewall.EndpointsID),
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallDataSourceConfig_name(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFirewallExists(resourceName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "network-firewall", fmt.Sprintf("firewall/%s", rName)),
					resource.TestCheckResourceAttr(dataSourceName, "delete_protection", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "description", ""),
					resource.TestCheckResourceAttrPair(dataSourceName, "firewall_policy_arn", policyResourceName, "arn"),
					resource.TestCheckResourceAttr(dataSourceName, "firewall_status.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "firewall_status.0.sync_states.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "firewall_status.0.sync_states.*.availability_zone", subnetResourceName, "availability_zone"),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceName, "firewall_status.0.sync_states.*", map[string]*regexp.Regexp{
						"attachment.0.endpoint_id": regexp.MustCompile(`vpce-`),
					}),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "firewall_status.0.sync_states.*.attachment.0.subnet_id", subnetResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttrPair(dataSourceName, "vpc_id", vpcResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "subnet_mapping.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(dataSourceName, "subnet_mapping.*.subnet_id", subnetResourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "update_token"),
				),
			},
		},
	})
}

func testAccFirewallDataSourceDependenciesConfig(rName string) string {
	return fmt.Sprintf(`
data "aws_availability_zones" "available" {
  state = "available"

  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

resource "aws_vpc" "test" {
  cidr_block = "192.168.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block        = cidrsubnet(aws_vpc.test.cidr_block, 8, 0)
  vpc_id            = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkfirewall_firewall_policy" "test" {
  name = %[1]q
  firewall_policy {
    stateless_fragment_default_actions = ["aws:drop"]
    stateless_default_actions          = ["aws:pass"]
  }
}
`, rName)
}

func testAccFirewallDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(
		testAccFirewallDataSourceDependenciesConfig(rName),
		fmt.Sprintf(`
resource "aws_networkfirewall_firewall" "test" {
  name                = %[1]q
  firewall_policy_arn = aws_networkfirewall_firewall_policy.test.arn
  vpc_id              = aws_vpc.test.id

  subnet_mapping {
    subnet_id = aws_subnet.test.id
  }
}

data "aws_networkfirewall_firewall" "test" {
  arn = aws_networkfirewall_firewall.test.arn
}
`, rName))
}

func testAccFirewallDataSourceConfig_name(rName string) string {
	return acctest.ConfigCompose(
		testAccFirewallDataSourceDependenciesConfig(rName),
		fmt.Sprintf(`
resource "aws_networkfirewall_firewall" "test" {
  name                = %[1]q
  firewall_policy_arn = aws_networkfirewall_firewall_policy.test.arn
  vpc_id              = aws_vpc.test.id

  subnet_mapping {
    subnet_id = aws_subnet.test.id
  }
}

data "aws_networkfirewall_firewall" "test" {
  name = %[1]q

  depends_on = [aws_networkfirewall_firewall.test]
}
`, rName))
}

func testAccFirewallDataSourceConfig_both(rName string) string {
	return acctest.ConfigCompose(
		testAccFirewallDataSourceDependenciesConfig(rName),
		fmt.Sprintf(`
resource "aws_networkfirewall_firewall" "test" {
  name                = %[1]q
  firewall_policy_arn = aws_networkfirewall_firewall_policy.test.arn
  vpc_id              = aws_vpc.test.id

  subnet_mapping {
    subnet_id = aws_subnet.test.id
  }
}

data "aws_networkfirewall_firewall" "test" {
  arn = aws_networkfirewall_firewall.test.arn
  name = %[1]q

  depends_on = [aws_networkfirewall_firewall.test]
}
`, rName))
}
