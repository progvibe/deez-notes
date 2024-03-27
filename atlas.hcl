variable "token" {
  type    = string
  default = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MTAwODE3NjksImlkIjoiNjYzOGJmMzctMjM0MS00ZDNlLWExYTItZmYyZGI3NWQxMjhkIn0.jGUXVm_q17lyoIc3FdYaAKkPR3ngWGFe-AFpLZFonpEdp2b9w8anziDcqpyS5vrgBv3UE0-VDDQgHujOeHuIAA"
}

env "turso" {
  url     = "libsql+wss://deez-notes-progvibe.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
}

env "local" {
  url = "libsql+ws://127.0.0.1:8080"
}