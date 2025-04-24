variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-west-2"
}

variable "environment" {
  description = "Environment name (e.g., dev, staging, prod)"
  type        = string
}

variable "instance_count" {
  description = "Number of EC2 instances to create"
  type        = number
  default     = 1
}

variable "enable_monitoring" {
  description = "Whether to enable detailed monitoring for instances"
  type        = bool
  default     = false
}

variable "vpc_configuration" {
  description = "VPC configuration object"
  type = object({
    vpc_id             = string
    subnet_ids         = list(string)
    security_group_ids = list(string)
    enable_nat_gateway = bool
    single_nat_gateway = bool
  })
}

variable "instance_settings" {
  description = "EC2 instance settings"
  type = object({
    instance_type = string
    ami_id        = string
    key_name      = optional(string)
    root_volume = object({
      size = number
      type = string
    })
    ebs_volumes = list(object({
      size        = number
      type        = string
      device_name = string
      encrypted   = bool
    }))
  })
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
  default     = {}
}

variable "ingress_rules" {
  description = "List of ingress rules for the security group"
  type = list(object({
    from_port   = number
    to_port     = number
    protocol    = string
    cidr_blocks = list(string)
    description = optional(string)
  }))
  default = []
}

variable "ssm_parameters" {
  description = "Map of SSM parameters to create"
  type        = map(string)
  default     = {}
  sensitive   = true
}

variable "allowed_ports" {
  description = "List of allowed ports"
  type        = list(number)
  default     = [22, 80, 443]
}

variable "load_balancer_config" {
  description = "Load balancer configuration"
  type = object({
    name               = string
    internal           = bool
    load_balancer_type = string
    subnets            = list(string)
    listeners = list(object({
      port     = number
      protocol = string
      default_action = object({
        type             = string
        target_group_arn = string
      })
    }))
  })
  default = null
}

variable "db_settings" {
  description = "Database settings"
  type = map(object({
    engine         = string
    engine_version = string
    instance_class = string
    allocated_storage = number
    storage_encrypted = bool
    parameters = list(object({
      name  = string
      value = string
    }))
  }))
  default = {}
}

variable "autoscaling_settings" {
  description = "Autoscaling group settings"
  type = tuple([
    string,   # name
    number,   # min size
    number,   # max size
    list(string)  # availability zones
  ])
  default = ["my-asg", 1, 3, ["us-west-2a", "us-west-2b"]]
}
