# resource "helm_release" "prometheus" {
#   name       = "prometheus"
#   repository = "https://prometheus-community.github.io/helm-charts"
#   chart      = "prometheus"
#   namespace  = kubernetes_namespace.monitoring.id
# }

# resource "helm_release" "loki" {
#   name       = "loki"
#   repository = "https://grafana.github.io/helm-charts"
#   chart      = "loki-stack"
#   namespace  = kubernetes_namespace.monitoring.id

#   set {
#     name  = "grafana.enabled"
#     value = "true"
#   }

#   set {
#     name  = "grafana.adminPassword"
#     value = "admin"
#   }
# }
