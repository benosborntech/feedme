resource "digitalocean_project" "project" {
  name      = "feedme"
  resources = [digitalocean_kubernetes_cluster.cluster.urn, digitalocean_loadbalancer.ingress_load_balancer.urn, digitalocean_domain.domain.urn]
}
