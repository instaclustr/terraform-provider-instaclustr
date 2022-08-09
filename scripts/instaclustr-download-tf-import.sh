#!/usr/bin/env bash

set -e

INSTACLUSTR_USERNAME="$INSTACLUSTR_TF_IMPORT_USERNAME"
INSTACLUSTR_API_KEY="$INSTACLUSTR_TF_IMPORT_API_KEY"
DEST_FILE_NAME="$1"

if [ -z "$DEST_FILE_NAME" ] || [ -z "$INSTACLUSTR_USERNAME" ] || [ -z "$INSTACLUSTR_API_KEY" ]
then
  echo

  if [ -z "$INSTACLUSTR_TF_IMPORT_USERNAME" ]
  then
    echo "Missing required environment variable 'INSTACLUSTR_TF_IMPORT_USERNAME'"
    echo
  fi

  if [ -z "$INSTACLUSTR_TF_IMPORT_API_KEY" ]
  then
    echo "Missing required environment variable 'INSTACLUSTR_TF_IMPORT_API_KEY'"
    echo
  fi

  echo "Usage: ./$0 <path_to_output_file>"
  echo "This script also depends on existence of 2 environment variables - 'INSTACLUSTR_TF_IMPORT_USERNAME' and 'INSTACLUSTR_TF_IMPORT_API_KEY' which should contain the Instaclustr username and Provisioning API Key respectively."
  echo "Example - ./$0 instaclustr/terraform-import.zip"
  echo
  exit 0
fi

INSTACLUSTR_API_URL="$INSTACLUSTR_TF_IMPORT_API_URL"

if [ -z "$INSTACLUSTR_API_URL" ]
then
  INSTACLUSTR_API_URL="https://api.instaclustr.com"
fi

curl $INSTACLUSTR_API_URL/cluster-management/v2/operations/terraform-import -u "$INSTACLUSTR_USERNAME:$INSTACLUSTR_API_KEY" --output "$DEST_FILE_NAME" --fail

echo "Terraform files downloaded to '$DEST_FILE_NAME'"
