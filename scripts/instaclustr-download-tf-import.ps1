$ErrorActionPreference = "Stop"

$SCRIPT_NAME = "instaclustr-download-tf-import.ps1"

if ($args.Count -lt 3) {
    Write-Output ""
    Write-Output "Usage: $SCRIPT_NAME <instaclustr_username> <instaclustr_provisioning_api_key> <path_to_output_file>"
    Write-Output "Example - $SCRIPT_NAME johndoe 0a1b2c3daabbccdd00112233e4f5g6h7 instaclustr\terraform-import.zip"
    Write-Output ""
    exit
}

$ZIP_FILE_NAME="instaclustr-terraform-import.zip"

$INSTACLUSTR_USERNAME = $args[0]
$INSTACLUSTR_API_KEY = $args[1]
$DEST_FILE_NAME = $args[2]

$basicAuthHeaderValue = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("${INSTACLUSTR_USERNAME}:$INSTACLUSTR_API_KEY"))
Invoke-WebRequest https://api.instaclustr.com/cluster-management/v2/operations/terraform-import -Headers @{Authorization="Basic $basicAuthHeaderValue"} -OutFile "$DEST_FILE_NAME"

Write-Host "Terraform files downloaded to '$DEST_FILE_NAME'"
