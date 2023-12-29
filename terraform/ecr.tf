resource "aws_ecr_repository" "bot" { #tfsec:ignore:aws-ecr-repository-customer-key tfsec:ignore:aws-ecr-enforce-immutable-repository
  name                 = "${local.name_base}-bot"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
