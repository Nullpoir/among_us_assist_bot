resource "aws_cloudwatch_log_group" "bot" { #tfsec:ignore:aws-cloudwatch-log-group-customer-key
  name              = "/${var.app_name}/${terraform.workspace}/bot"
  retention_in_days = var.ecs_log_retention_period
}
