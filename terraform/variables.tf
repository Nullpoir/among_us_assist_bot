variable "app_name" {
  default = "among_us_assist_bot"
  type    = string
}
variable "bot_image_tag" {
  default = "latest"
  type    = string
}

variable "ecs_log_retention_period" {
  default = 7
  type    = number
}
