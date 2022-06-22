resource "instaclustr_gcp_vpc_peer_v2" "myvpcpeer" {
  cdc_id                  = "b8129e68-2ee9-4d5e-a29b-74d233b7b4bb"
  name                    = "peering-b8129e68-2ee9-4d5e-a29b-74d233b7b4bb"
  peer_project_id         = "example-project123"
  peer_subnets            = ["10.1.0.0/16", "10.2.0.0/16"]
  peer_vpc_network_name   = "network-aabb1122"
}
