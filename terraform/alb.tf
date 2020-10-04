resource "aws_alb" "application_load_balancer" {
  name               = "todoapi"
  load_balancer_type = "application"
  subnets = [
    aws_subnet.eu-west-1a-private.id,
    aws_subnet.eu-west-1b-private.id
  ]

  security_groups = [aws_security_group.load_balancer_security_group.id]
}

resource "aws_security_group" "load_balancer_security_group" {
  vpc_id = aws_vpc.default.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}