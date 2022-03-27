data "instaclustr_clusters_v2" "my_clusters" {}

data "instaclustr_cassandra_cluster_v2_instance" "cassandra_clusters" {
  for_each = {
  for cluster in data.instaclustr_clusters_v2.my_clusters.clusters :
  cluster.id => cluster if cluster.application == "APACHE_CASSANDRA"
  }
  id = each.value.id
}

resource "your_resource" "resource_based_on_existing_cassandra_clusters_example" {
  for_each      = data.instaclustr_cassandra_cluster_v2_instance.cassandra_clusters
  //.....
  your_property = each.value.<cassandra_cluster_property>
  //.....
}
