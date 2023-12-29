resource "aws_iam_role" "ecs_task_execution" {
  name = "${local.name_base}-ecs-task-execution"

  assume_role_policy = jsonencode(
    {
      Version = "2012-10-17"
      Statement = [
        {
          Action = "sts:AssumeRole"
          Effect = "Allow"
          Principal = {
            Service = "ecs-tasks.amazonaws.com"
          },
          Sid = ""
        },
      ]
    }
  )
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_managed" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = aws_iam_policy.ecs_task_execution.arn
}

resource "aws_iam_role" "ecs_task" {
  name = "${local.name_base}-ecs-task"

  assume_role_policy = jsonencode(
    {
      Version = "2012-10-17"
      Statement = [
        {
          Action = "sts:AssumeRole"
          Effect = "Allow"
          Principal = {
            Service = "ecs-tasks.amazonaws.com"
          },
          Sid = ""
        },
      ]
    }
  )
}

# resource "aws_iam_role_policy_attachment" "ecs_task" {
#   role       = aws_iam_role.ecs_task.name
#   policy_arn = aws_iam_policy.ecs_task.arn
# }
