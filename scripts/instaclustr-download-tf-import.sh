#!/usr/bin/env bash

set -e

if [ "$#" -lt 3 ]
then
  echo
  echo "Usage: $0 <instaclustr_username> <instaclustr_provisioning_api_key> <path_to_output_file>"
  echo "Example - $0 johndoe 0a1b2c3daabbccdd00112233e4f5g6h7 instaclustr/terraform-import.zip"
  echo
  exit 0
fi

INSTACLUSTR_USERNAME="$1"
INSTACLUSTR_API_KEY="$2"
DEST_FILE_NAME="$3"

curl https://api.instaclustr.com/cluster-management/v2/operations/terraform-import -u "$INSTACLUSTR_USERNAME:$INSTACLUSTR_API_KEY" --output "$DEST_FILE_NAME" --fail

echo "Terraform files downloaded to '$DEST_FILE_NAME'"
