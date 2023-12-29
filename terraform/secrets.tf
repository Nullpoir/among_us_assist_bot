resource "aws_secretsmanager_secret" "client_id" { #tfsec:ignore:aws-ssm-secret-use-customer-key
  name = "${local.name_base}-client-id"
}

resource "aws_secretsmanager_secret" "client_secret" { #tfsec:ignore:aws-ssm-secret-use-customer-key
  name = "${local.name_base}-client-secret"
}

resource "aws_secretsmanager_secret" "access_token" { #tfsec:ignore:aws-ssm-secret-use-customer-key
  name = "${local.name_base}-access-token"
}
