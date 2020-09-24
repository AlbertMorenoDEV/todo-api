provider "aws" {
  version = "~> 2.0"
  region  = "eu-west-1"
}

resource "aws_ecr_repository" "todoapi_ecr_repo" {
  name = "todo-api"
}