#!/usr/bin/env bash

set -e

INSTACLUSTR_USERNAME="$IC_USERNAME"
INSTACLUSTR_API_KEY="$IC_API_KEY"
DEST_FOLDER_NAME="$1"

if [ -z "$DEST_FOLDER_NAME" ] || [ -z "$INSTACLUSTR_USERNAME" ] || [ -z "$INSTACLUSTR_API_KEY" ]
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

  echo "Usage: $0 <destination_directory> [auto-approve]"
  echo "'auto-approve' is an optional argument which will skip the confirmation on emptying the destination directory."
  echo "This script also depends on existence of 2 environment variables - 'IC_USERNAME' and 'IC_API_KEY' which should contain the Instaclustr username and Provisioning API Key respectively."
  echo "Example - $0 instaclustr"
  echo
  exit 0
fi

INSTACLUSTR_API_URL="$IC_API_URL"

if [ -z "$INSTACLUSTR_API_URL" ]
then
  INSTACLUSTR_API_URL="https://api.instaclustr.com"
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

curl $INSTACLUSTR_API_URL/cluster-management/v2/operations/terraform-import -u "$INSTACLUSTR_USERNAME:$INSTACLUSTR_API_KEY" --output "$ZIP_FILE_NAME" --fail

rm -rf "$DEST_FOLDER_NAME"

mkdir -p "$DEST_FOLDER_NAME"

tar xvf "$ZIP_FILE_NAME" -C "$DEST_FOLDER_NAME"

rm "$ZIP_FILE_NAME"

cd "$DEST_FOLDER_NAME"

terraform init

sh import-all.sh

echo "Script execution completed."
