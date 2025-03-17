$ErrorActionPreference = "Stop"

$SCRIPT_NAME = "./instaclustr-download-tf-import.ps1"

$INSTACLUSTR_USERNAME = [System.Environment]::GetEnvironmentVariable("IC_USERNAME")
$INSTACLUSTR_API_KEY = [System.Environment]::GetEnvironmentVariable("IC_API_KEY")
$DEST_FILE_NAME = $args[0]

if ($DEST_FILE_NAME -eq $null -Or $INSTACLUSTR_USERNAME -eq $null -Or $INSTACLUSTR_API_KEY -eq $null) {
    Write-Output ""

    if ($INSTACLUSTR_USERNAME -eq $null) {
        Write-Output "Missing required environment variable 'IC_USERNAME'"
        Write-Output ""
    }

    if ($INSTACLUSTR_API_KEY -eq $null) {
        Write-Output "Missing required environment variable 'IC_API_KEY'"
        Write-Output ""
    }

    Write-Output "Usage: $SCRIPT_NAME <path_to_output_file>"
    Write-Output "This script also depends on existence of 2 environment variables - 'IC_USERNAME' and 'IC_API_KEY' which should contain the Instaclustr username and Provisioning API Key respectively."
    Write-Output "Example - $SCRIPT_NAME instaclustr\terraform-import.zip"
    Write-Output ""
    exit
}

$basicAuthHeaderValue = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("${INSTACLUSTR_USERNAME}:$INSTACLUSTR_API_KEY"))
$apiUrl = "https://api.instaclustr.com/cluster-management/v2/operations/generate-terraform-code/v2"

try {
    Invoke-WebRequest $apiUrl `
        -Headers @{ Authorization = "Basic $basicAuthHeaderValue" } `
        -OutFile "$DEST_FILE_NAME" `
        -ErrorAction Stop

    Write-Host "Terraform files downloaded to '$DEST_FILE_NAME'"
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
    }
    else {
        Write-Host $_.Exception.Message -ForegroundColor Red
    }
    exit 1
}

