$ErrorActionPreference = "Stop"

$SCRIPT_NAME = "./instaclustr-download-tf-import.ps1"

$INSTACLUSTR_USERNAME = [System.Environment]::GetEnvironmentVariable("INSTACLUSTR_TF_IMPORT_USERNAME")
$INSTACLUSTR_API_KEY = [System.Environment]::GetEnvironmentVariable("INSTACLUSTR_TF_IMPORT_API_KEY")
$DEST_FILE_NAME = $args[0]

if ($DEST_FILE_NAME -eq $null -Or $INSTACLUSTR_USERNAME -eq $null -Or $INSTACLUSTR_API_KEY -eq $null) {
    Write-Output ""

    if ($INSTACLUSTR_USERNAME -eq $null) {
        Write-Output "Missing required environment variable 'INSTACLUSTR_TF_IMPORT_USERNAME'"
        Write-Output ""
    }

    if ($INSTACLUSTR_API_KEY -eq $null) {
        Write-Output "Missing required environment variable 'INSTACLUSTR_TF_IMPORT_API_KEY'"
        Write-Output ""
    }

    Write-Output "Usage: $SCRIPT_NAME <path_to_output_file>"
    Write-Output "This script also depends on existence of 2 environment variables - 'INSTACLUSTR_TF_IMPORT_USERNAME' and 'INSTACLUSTR_TF_IMPORT_API_KEY' which should contain the Instaclustr username and Provisioning API Key respectively."
    Write-Output "Example - $SCRIPT_NAME instaclustr\terraform-import.zip"
    Write-Output ""
    exit
}

$basicAuthHeaderValue = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("${INSTACLUSTR_USERNAME}:$INSTACLUSTR_API_KEY"))
Invoke-WebRequest https://api.instaclustr.com/cluster-management/v2/operations/terraform-import -Headers @{Authorization="Basic $basicAuthHeaderValue"} -OutFile "$DEST_FILE_NAME"

Write-Host "Terraform files downloaded to '$DEST_FILE_NAME'"
