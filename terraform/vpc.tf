module "vpc" { # tfsec:ignore:aws-ec2-require-vpc-flow-logs-for-all-vpcs
  source  = "terraform-aws-modules/vpc/aws"
  version = "4.0.1"

  name = "${local.name_base}-vpc"
  cidr = "10.0.0.0/16"

  azs = ["ap-northeast-1a", "ap-northeast-1c"]

  private_subnets     = ["10.0.1.0/24", "10.0.2.0/24"]
  database_subnets    = ["10.0.21.0/24", "10.0.22.0/24"]
  elasticache_subnets = ["10.0.31.0/24", "10.0.32.0/24"]
  public_subnets      = ["10.0.101.0/24", "10.0.102.0/24"]

  enable_nat_gateway      = true
  single_nat_gateway      = false
  one_nat_gateway_per_az  = true
  map_public_ip_on_launch = false

  manage_default_security_group  = true
  default_security_group_egress  = []
  default_security_group_ingress = []
}
