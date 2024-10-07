resource "digitalocean_domain" "domain" {
  name = var.tld
}

resource "digitalocean_record" "a_records" {
  domain = digitalocean_domain.domain.id
  type   = "A"
  ttl    = 60
  name   = "@"
  value  = digitalocean_loadbalancer.ingress_load_balancer.ip
}
