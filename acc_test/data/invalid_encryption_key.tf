provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_encryption_key" "invalid" {
    alias = "ic_test_key"
    arn = "%s!@#$"
}
