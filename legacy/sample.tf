
resource "kubernetes_deployment" "sample_deployments" {
  metadata {
    name      = "test-deployment"
    namespace = kubernetes_namespace.app.metadata[0].name
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "test-deployment"
      }
    }
    template {
      metadata {
        labels = {
          app = "test-deployment"
        }
      }
      spec {
        container {
          image = "nginxdemos/hello"
          name  = "nginx-hello"
          port {
            container_port = 80
          }
          resources {
            limits = {
              memory = "512M"
              cpu    = "1"
            }
            requests = {
              memory = "256M"
              cpu    = "50m"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "sample_services" {
  metadata {
    name      = "test-service"
    namespace = kubernetes_namespace.app.metadata[0].name
  }
  spec {
    selector = {
      app = "test-deployment"
    }
    port {
      port = 80
    }
  }
}
