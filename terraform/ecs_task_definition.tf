resource "aws_ecs_task_definition" "bot" {
  family                   = "${local.name_base}-bot"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution.arn
  task_role_arn            = aws_iam_role.ecs_task.arn
  container_definitions = jsonencode([
    {
      name      = "${local.name_base}-bot"
      image     = "${aws_ecr_repository.bot.repository_url}:${var.bot_image_tag}"
      essential = true
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.bot.name
          "awslogs-region"        = data.aws_region.current.name
          "awslogs-stream-prefix" = "bot"
        }
      }
      command = [
        "python",
        "application.py"
      ],
      secrets : [
        {
          name      = "CLIENT_ID"
          valueFrom = aws_secretsmanager_secret.client_id.arn
        },
        {
          name      = "CLIENT_SECRET"
          valueFrom = aws_secretsmanager_secret.client_secret.arn
        },
        {
          name      = "ACCESS_TOKEN"
          valueFrom = aws_secretsmanager_secret.access_token.arn
        },
      ],
      environment : [
        {
          name  = "KVS_HOST",
          value = aws_elasticache_cluster.redis.cache_nodes[0].address
        },
        {
           name = "KVS_PORT",
           value = "6379"
        }
      ]
      portMappings = [
        {
          containerPort = 3000
          hostPort      = 3000
        }
      ]
    }
  ])

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }
}
