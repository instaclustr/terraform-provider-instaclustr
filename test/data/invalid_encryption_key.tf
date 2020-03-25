provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_encryption_key" "invalid" {
    alias = "%s"
    arn = "%s!@#$"
}
