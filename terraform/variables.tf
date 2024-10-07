variable "do_token" {
  type      = string
  sensitive = true
}

variable "do_user" {
  type      = string
  sensitive = true
}

variable "do_region" {
  type    = string
  default = "syd1"
}

variable "tld" {
  type    = string
  default = "api.feedme.benosborn.tech"
}
