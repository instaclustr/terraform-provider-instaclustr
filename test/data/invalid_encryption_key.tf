provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_encryption_key" "invalid" {
    alias = "ic_test_key"
    arn = "%s!@#$"
}