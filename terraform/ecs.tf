resource "aws_ecr_repository" "todoapi_ecr_repo" {
  name = "todo-api"
}

resource "aws_ecs_cluster" "todoapi_ecs_cluster" {
  name = "todo-api"
}

resource "aws_ecs_task_definition" "todoapi_app" {
  family                   = "todo-api-app"
  container_definitions    = <<DEFINITION
  [
    {
      "name": "todo-api-app",
      "image": "${aws_ecr_repository.todoapi_ecr_repo.repository_url}",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 3000,
          "hostPort": 3000
        }
      ],
      "memory": 512,
      "cpu": 256
    }
  ]
  DEFINITION
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = 512
  cpu                      = 256
  execution_role_arn       = aws_iam_role.ecsTaskExecutionRole.arn
}

resource "aws_iam_role" "ecsTaskExecutionRole" {
  name               = "ecsTaskExecutionRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "ecsTaskExecutionRole_policy" {
  role       = aws_iam_role.ecsTaskExecutionRole.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_service" "todoapi" {
  name            = "todo-api-app"
  cluster         = aws_ecs_cluster.todoapi_ecs_cluster.id
  task_definition = aws_ecs_task_definition.todoapi_app.arn
  launch_type     = "FARGATE"
  desired_count   = 3

  network_configuration {
    subnets          = [aws_subnet.eu-west-1a-private.id, aws_subnet.eu-west-1b-private.id]
    assign_public_ip = true
  }
}