variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "aws_key_path" {}
variable "aws_key_name" {}

variable "aws_region" {
  description = "Region for the VPC"
  default = "eu-west-1"
}

variable "vpc_cidr" {
  description = "CIDR for the whole VPC"
  default = "10.0.0.0/16"
}

variable "public_subnet_cidr" {
  description = "CIDR for the Public Subnet"
  default = "10.0.0.0/24"
}

variable "private_subnet_cidr_1" {
  description = "CIDR for the Private Subnet"
  default = "10.0.1.0/24"
}

variable "private_subnet_cidr_2" {
  description = "CIDR for the Private Subnet"
  default = "10.0.2.0/24"
}