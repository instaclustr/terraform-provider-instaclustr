$ErrorActionPreference = "Stop"

$SCRIPT_NAME = "./instaclustr-auto-import.ps1"

$INSTACLUSTR_USERNAME = [System.Environment]::GetEnvironmentVariable("IC_USERNAME")
$INSTACLUSTR_API_KEY = [System.Environment]::GetEnvironmentVariable("IC_API_KEY")
$DEST_FOLDER_NAME = $args[0]

if ($DEST_FOLDER_NAME -eq $null -Or $INSTACLUSTR_USERNAME -eq $null -Or $INSTACLUSTR_API_KEY -eq $null) {
    Write-Output ""

    if ($INSTACLUSTR_USERNAME -eq $null) {
        Write-Output "Missing required environment variable 'IC_USERNAME'"
        Write-Output ""
    }

    if ($INSTACLUSTR_API_KEY -eq $null) {
        Write-Output "Missing required environment variable 'IC_API_KEY'"
        Write-Output ""
    }

    Write-Output "Usage: $SCRIPT_NAME <destination_directory> [auto-approve]"
    Write-Output "'auto-approve' is an optional argument which will skip the confirmation on emptying the destination directory."
    Write-Output "This script also depends on existence of 2 environment variables - 'IC_USERNAME' and 'IC_API_KEY' which should contain the Instaclustr username and Provisioning API Key respectively."
    Write-Output "Example - $SCRIPT_NAME instaclustr"
    Write-Output ""
    exit
}

$ZIP_FILE_NAME="instaclustr-terraform-import.zip"

if ($args[1] -ne "auto-approve") {
    $choice = Read-Host "This script will delete the contents of the folder '$DEST_FOLDER_NAME', do you wish to proceed (y/n)? "
    if ($choice -notmatch '^[Yy]$') {
        Write-Host "Execution cancelled."
        exit
    }
}

$basicAuthHeaderValue = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("${INSTACLUSTR_USERNAME}:$INSTACLUSTR_API_KEY"))
$apiUrl = "https://api.instaclustr.com/cluster-management/v2/operations/generate-terraform-code/v2"

try {
    Invoke-WebRequest $apiUrl `
        -Headers @{ Authorization = "Basic $basicAuthHeaderValue" } `
        -OutFile "$ZIP_FILE_NAME" `
        -ErrorAction Stop

    Write-Host "Terraform files downloaded to '$ZIP_FILE_NAME'"
}
catch {
    Write-Host "Error occurred while calling the endpoint:" -ForegroundColor Red
    if ($_.Exception.Response) {
        $stream   = $_.Exception.Response.GetResponseStream()
        $reader   = New-Object System.IO.StreamReader($stream)
        $response = $reader.ReadToEnd()

        Write-Host "Server response:"
        Write-Host $response
        Write-Host ""
        Write-Host "For more information on how to resolve this issue, try to generate the Terraform configuration from the Instaclustr Console under Settings > Cluster Resources > Terraform > Download."
        Remove-Item "$ZIP_FILE_NAME"
    }
    else {
        Write-Host $_.Exception.Message -ForegroundColor Red
    }
    exit 1
}


cmd /c rmdir /s /q "$DEST_FOLDER_NAME" #need to do this instead of Remove-Item or other methods to deal with symlinks to terraform provider

Expand-Archive -Path "$ZIP_FILE_NAME" -DestinationPath "$DEST_FOLDER_NAME"

Remove-Item "$ZIP_FILE_NAME"

Push-Location "$DEST_FOLDER_NAME"

terraform init

./import-all.ps1

Pop-Location

Write-Host "Script execution completed."
