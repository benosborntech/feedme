variable "google_client_id" {
  description = "Google client id"
  type        = string
  sensitive   = true
}

variable "google_client_secret" {
  description = "Google client secret"
  type        = string
  sensitive   = true
}

variable "server_secret" {
  description = "Server secret"
  type        = string
  sensitive   = true
}

variable "mysql_dsn" {
  description = "MySQL DSN"
  type        = string
  sensitive   = true
}

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
