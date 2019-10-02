variable "base_cidr_block_public" {
  description = "A /16 CIDR range definition, such as 10.1.0.0/16, that the VPC will use"
  default = "10.1.1.0/20"
}

variable "base_cidr_block_private" {
  description = "A /16 CIDR range definition, such as 10.1.0.0/16, that the VPC will use"
  default = "10.1.100.0/20"
}

variable "aws_region" {
  description = "The AWS region"
  type = string
  default = "us-east-1"
}

variable "base_cidr_block" {
  description = "A /16 CIDR range definition, such as 10.1.0.0/16, that the VPC will use"
  default = "10.1.0.0/16"
}

variable "tag_name" {
  description = "a tag to apply to name the resource group"
  default = "terraform"
}

variable "availability_zones" {
  description = "A list of availability zones in which to create subnets"
  type = list(string)
  default = ["us-east-1a", "us-east-1b","us-east-1c","us-east-1d","us-east-1e"]
}