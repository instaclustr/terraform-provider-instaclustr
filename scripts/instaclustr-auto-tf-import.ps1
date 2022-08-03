$ErrorActionPreference = "Stop"

$SCRIPT_NAME = "instaclustr-auto-import.ps1"

if ($args.Count -lt 3) {
    Write-Output ""
    Write-Output "Usage: $SCRIPT_NAME <instaclustr_username> <instaclustr_provisioning_api_key> <destination_directory> [auto-approve]"
    Write-Output "'auto-approve' is an optional argument which will skip the confirmation on emptying the destination directory."
    Write-Output "Example - $SCRIPT_NAME johndoe 0a1b2c3daabbccdd00112233e4f5g6h7 instaclustr"
    Write-Output ""
    exit
}

$ZIP_FILE_NAME="instaclustr-terraform-import.zip"

$INSTACLUSTR_USERNAME = $args[0]
$INSTACLUSTR_API_KEY = $args[1]
$DEST_FOLDER_NAME = $args[2]

if ($args[3] -ne "auto-approve") {
    $choice = Read-Host "This script will delete the contents of the folder '$DEST_FOLDER_NAME', do you wish to proceed (y/n)? "
    if ($choice -notmatch '^[Yy]$') {
        Write-Host "Execution cancelled."
        exit
    }
}

$basicAuthHeaderValue = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("${INSTACLUSTR_USERNAME}:$INSTACLUSTR_API_KEY"))
Invoke-WebRequest https://api.instaclustr.com/cluster-management/v2/operations/terraform-import -Headers @{Authorization="Basic $basicAuthHeaderValue"} -OutFile "$ZIP_FILE_NAME"


cmd /c rmdir /s /q "$DEST_FOLDER_NAME" #need to do this instead of Remove-Item or other methods to deal with symlinks to terraform provider

Expand-Archive -Path "$ZIP_FILE_NAME" -DestinationPath "$DEST_FOLDER_NAME"

Remove-Item "$ZIP_FILE_NAME"

Push-Location "$DEST_FOLDER_NAME"

terraform init

./import-all.ps1 $INSTACLUSTR_USERNAME $INSTACLUSTR_API_KEY

Pop-Location

Write-Host "Script execution completed."
