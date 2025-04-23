variable "instance_count" {
  description = "Number of instances to create"
  type        = number
  validation {
    condition     = var.instance_count > 0
    error_message = "Instance count must be greater than 0."
  }
}

variable "instance_name" {
  description = "Name of the instance"
  type        = string
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
  default     = {}
}

variable "enable_monitoring" {
  description = "Enable detailed monitoring"
  type        = bool
  default     = false
}

variable "subnet_ids" {
  description = "List of subnet IDs"
  type        = list(string)
}

variable "complex_settings" {
  description = "Complex configuration settings"
  type = object({
    timeout     = number
    retries     = number
    logging     = bool
    performance = object({
      cpu_allocation    = number
      memory_allocation = number
    })
  })
  default = {
    timeout     = 30
    retries     = 3
    logging     = true
    performance = {
      cpu_allocation    = 1
      memory_allocation = 2
    }
  }
}

variable "multi_line_type" {
  description = "Variable with multi-line type definition"
  type = list(
    object({
      name    = string,
      enabled = bool,
      config  = map(string)
    })
  )
  default = []
}