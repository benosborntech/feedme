resource "digitalocean_domain" "domain" {
  name = var.tld
}

resource "digitalocean_record" "a_record" {
  domain = digitalocean_domain.domain.id
  type   = "A"
  ttl    = 60
  name   = "@"
  value  = digitalocean_loadbalancer.ingress_load_balancer.ip
}

resource "digitalocean_record" "cname_record" {
  domain = digitalocean_domain.domain.id
  type   = "CNAME"
  ttl    = 60
  name   = "www"
  value  = "@"
}
