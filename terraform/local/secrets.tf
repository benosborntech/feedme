resource "kubernetes_secret" "secrets" {
  metadata {
    name      = "secrets"
    namespace = kubernetes_namespace.dev.metadata[0].name
  }

  data = {
    "google_client_id" : var.google_client_id
    "google_client_secret" : var.google_client_secret
    "server_secret": var.server_secret
    "mysql_dsn": var.mysql_dsn
  }

  type = "Opaque"
}