# Container registry
resource "digitalocean_container_registry" "container_registry" {
  name                   = "feedmecr"
  subscription_tier_slug = "basic"
}

resource "kubernetes_secret" "docker_registry_secret" {
  metadata {
    name      = "dockercred"
    namespace = kubernetes_namespace.app.metadata[0].name
  }

  data = {
    ".dockerconfigjson" = jsonencode({
      auths = {
        "registry.digitalocean.com" = {
          auth = base64encode("${var.do_user}:${var.do_token}")
        }
      }
    })
  }

  type = "kubernetes.io/dockerconfigjson"
}

# Cluster
resource "digitalocean_kubernetes_cluster" "cluster" {
  name    = "feedme-cluster"
  region  = var.do_region
  version = "1.29.1-do.0"

  node_pool {
    name       = "feedme-worker-pool"
    size       = "s-1vcpu-2gb"
    auto_scale = true
    min_nodes  = 2
    max_nodes  = 2
  }
}

# Load balancer
resource "digitalocean_loadbalancer" "ingress_load_balancer" {
  name   = "feedme-lb"
  region = var.region
  size   = "lb-small"

  forwarding_rule {
    entry_port     = 80
    entry_protocol = "http"

    target_port     = 80
    target_protocol = "http"

  }

  lifecycle {
    ignore_changes = [
      forwarding_rule,
    ]
  }
}

resource "helm_release" "nginx_ingress_chart" {
  name       = "nginx-ingress-controller"
  namespace  = "default"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "nginx-ingress-controller"
  set {
    name  = "service.type"
    value = "LoadBalancer"
  }
  set {
    name  = "service.annotations.kubernetes\\.digitalocean\\.com/load-balancer-id"
    value = digitalocean_loadbalancer.ingress_load_balancer.id
  }
  depends_on = [
    digitalocean_loadbalancer.ingress_load_balancer,
  ]
}
