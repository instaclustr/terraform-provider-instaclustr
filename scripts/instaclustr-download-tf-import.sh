#!/usr/bin/env bash

set -e

INSTACLUSTR_USERNAME="$IC_USERNAME"
INSTACLUSTR_API_KEY="$IC_API_KEY"
DEST_FILE_NAME="$1"

if [ -z "$DEST_FILE_NAME" ] || [ -z "$INSTACLUSTR_USERNAME" ] || [ -z "$INSTACLUSTR_API_KEY" ]
then
  echo

  if [ -z "$IC_USERNAME" ]
  then
    echo "Missing required environment variable 'IC_USERNAME'"
    echo
  fi

  if [ -z "$IC_API_KEY" ]
  then
    echo "Missing required environment variable 'IC_API_KEY'"
    echo
  fi

  echo "Usage: $0 <path_to_output_file>"
  echo "This script also depends on existence of 2 environment variables - 'IC_USERNAME' and 'IC_API_KEY' which should contain the Instaclustr username and Provisioning API Key respectively."
  echo "Example - $0 instaclustr/terraform-import.zip"
  echo
  exit 0
fi

INSTACLUSTR_API_URL="$IC_API_URL"

if [ -z "$INSTACLUSTR_API_URL" ]
then
  INSTACLUSTR_API_URL="https://api.instaclustr.com"
fi

curl $INSTACLUSTR_API_URL/cluster-management/v2/operations/generate-terraform-code/v2 -u "$INSTACLUSTR_USERNAME:$INSTACLUSTR_API_KEY" --output "$DEST_FILE_NAME" --fail

echo "Terraform files downloaded to '$DEST_FILE_NAME'"
