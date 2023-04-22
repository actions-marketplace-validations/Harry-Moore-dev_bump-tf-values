locals {
  # pin the target versions of the code
  other_code_version = "3.3.3.3"
  code_version       = "v2.55.4"
}

output "test_version_string" {
  value = var.other_code_version
}

output "test_version_number" {
  value = var.code_version
}
