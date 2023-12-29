resource "aws_iam_policy" "ecs_task_execution" {
  name        = "${local.name_base}-ecs-task-execution"
  description = "ECS Task Execution Policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "secretsmanager:GetSecretValue",
        ]
        Effect = "Allow"
        Resource = [
          aws_secretsmanager_secret.client_id.arn,
          aws_secretsmanager_secret.client_secret.arn,
          aws_secretsmanager_secret.access_token.arn,
        ]
      },
      {
        Action = [
          "ecr:GetAuthorizationToken",
        ]
        Effect = "Allow"
        Resource = [
          "*"
        ]
      }
    ]
  })
}
