module "security-group_app" { #tfsec:ignore:aws-ec2-no-public-egress-sgr
  source  = "terraform-aws-modules/security-group/aws"
  version = "4.17.2"

  name        = "${local.name_base}-app"
  description = "Security group for APP"
  vpc_id      = module.vpc.vpc_id

  egress_cidr_blocks = ["0.0.0.0/0"]
  egress_rules       = ["all-all"]
}

module "security-group_kvs" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "4.17.2"

  name        = "${local.name_base}-kvs"
  description = "Security group for KVS"
  vpc_id      = module.vpc.vpc_id

  computed_ingress_with_source_security_group_id = [
    {
      rule                     = "redis-tcp"
      source_security_group_id = module.security-group_app.security_group_id
    },
  ]
  number_of_computed_ingress_with_source_security_group_id = 1
}
