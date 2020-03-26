provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_encryption_key" "valid" {
    alias = "ic_test_key"
    arn = "%s"
}
