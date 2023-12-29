resource "aws_ecs_service" "bot" {
  name            = "${local.name_base}-bot"
  cluster         = module.ecs_cluster.cluster_id
  task_definition = aws_ecs_task_definition.bot.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  enable_execute_command = true

  network_configuration {
    subnets          = module.vpc.private_subnets
    security_groups  = [module.security-group_app.security_group_id, module.security-group_app.security_group_id]
    assign_public_ip = false
  }
}
