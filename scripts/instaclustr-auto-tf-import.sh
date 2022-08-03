#!/usr/bin/env bash

set -e

if [ "$#" -lt 3 ]
then
  echo
  echo "Usage: $0 <instaclustr_username> <instaclustr_provisioning_api_key> <destination_directory> [auto-approve]"
  echo "'auto-approve' is an optional argument which will skip the confirmation on emptying the destination directory."
  echo "Example - $0 johndoe 0a1b2c3daabbccdd00112233e4f5g6h7 instaclustr"
  echo
  exit 0
fi


ZIP_FILE_NAME="instaclustr-terraform-import.zip"

INSTACLUSTR_USERNAME="$1"
INSTACLUSTR_API_KEY="$2"
DEST_FOLDER_NAME="$3"

if ! [[ "$4" = "auto-approve" ]]
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

sh import-all.sh $INSTACLUSTR_USERNAME $INSTACLUSTR_API_KEY

echo "Script execution completed."
