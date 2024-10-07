resource "kubernetes_ingress_v1" "default_cluster_ingress" {
  depends_on = [
    helm_release.nginx_ingress_chart,
  ]

  metadata {
    name      = "ingress"
    namespace = kubernetes_namespace.app.metadata[0].name
    annotations = {
      "kubernetes.io/ingress.class"          = "nginx"
      "ingress.kubernetes.io/rewrite-target" = "/"
      "cert-manager.io/cluster-issuer"       = "letsencrypt"
    }
  }

  spec {
    rule {
      host = var.tld

      http {
        path {
          backend {
            service {
              name = "test-service"
              port {
                number = 80
              }
            }
          }
          path = "/"
        }
      }
    }

    tls {
      secret_name = "letsencrypt"
      hosts       = [var.tld]
    }
  }
}


