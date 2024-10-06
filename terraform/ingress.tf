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
      host = "feedme.benosborn.tech"

      http {
        path {
          backend {
            service {
              name = "apigw-service"
              port {
                number = 3000
              }
            }
          }
          path = "/"
        }
      }
    }

    tls {
      secret_name = "letsencrypt"
      hosts       = ["feedme.benosborn.tech"]
    }
  }
}


