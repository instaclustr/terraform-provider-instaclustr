resource "instaclustr_vpc_peer" "myvpcpeer" {
  cdc_id              = "b8129e68-2ee9-4d5e-a29b-74d233b7b4bb"
  peer_aws_account_id = "123456789123"
  peer_vpc_id         = "vpc-aaaa1234"
  peer_subnets        = ["174.16.0.0/20", "172.16.0.0/20", "173.16.0.0/20"]
  peer_region         = "US_EAST_1"
}
