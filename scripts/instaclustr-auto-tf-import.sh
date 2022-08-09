#!/usr/bin/env bash

set -e

INSTACLUSTR_USERNAME="$INSTACLUSTR_TF_IMPORT_USERNAME"
INSTACLUSTR_API_KEY="$INSTACLUSTR_TF_IMPORT_API_KEY"
DEST_FOLDER_NAME="$1"

if [ -z "$DEST_FOLDER_NAME" ] || [ -z "$INSTACLUSTR_USERNAME" ] || [ -z "$INSTACLUSTR_API_KEY" ]
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

  echo "Usage: $0 <destination_directory> [auto-approve]"
  echo "'auto-approve' is an optional argument which will skip the confirmation on emptying the destination directory."
  echo "This script also depends on existence of 2 environment variables - 'INSTACLUSTR_TF_IMPORT_USERNAME' and 'INSTACLUSTR_TF_IMPORT_API_KEY' which should contain the Instaclustr username and Provisioning API Key respectively."
  echo "Example - $0 instaclustr"
  echo
  exit 0
fi


ZIP_FILE_NAME="instaclustr-terraform-import.zip"

if ! [[ "$2" = "auto-approve" ]]
then
  read -p "This script will delete the contents of the folder '$DEST_FOLDER_NAME', do you wish to proceed (y/n)? " choice
  echo

  if ! [[ $choice =~ ^[Yy]$ ]]
  then
    echo "Execution cancelled."
    exit 0
  fi
fi

curl https://api.instaclustr.com/cluster-management/v2/operations/terraform-import -u "$INSTACLUSTR_USERNAME:$INSTACLUSTR_API_KEY" --output "$ZIP_FILE_NAME" --fail

rm -rf "$DEST_FOLDER_NAME"

mkdir -p "$DEST_FOLDER_NAME"

tar xvf "$ZIP_FILE_NAME" -C "$DEST_FOLDER_NAME"

rm "$ZIP_FILE_NAME"

cd "$DEST_FOLDER_NAME"

terraform init

sh import-all.sh

echo "Script execution completed."
