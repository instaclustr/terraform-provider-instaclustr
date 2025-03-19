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

$apiUrl = "https://api.instaclustr.com/cluster-management/v2/operations/generate-terraform-code/v2"
$basicAuthHeaderValue = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("${INSTACLUSTR_USERNAME}:$INSTACLUSTR_API_KEY"))

try {
    $handler = New-Object System.Net.Http.HttpClientHandler
    $handler.AutomaticDecompression = [System.Net.DecompressionMethods]::None

    $client = New-Object System.Net.Http.HttpClient($handler)

    $client.DefaultRequestHeaders.Authorization = 
        [System.Net.Http.Headers.AuthenticationHeaderValue]::new("Basic", $basicAuthHeaderValue)

    $response = $client.GetAsync($apiUrl).Result

    if ($response.IsSuccessStatusCode) {
        $zipBytes = $response.Content.ReadAsByteArrayAsync().Result
        [System.IO.File]::WriteAllBytes($ZIP_FILE_NAME, $zipBytes)
    }
    else {
        $jsonString = $response.Content.ReadAsStringAsync().Result
        Write-Host ""
        Write-Host "Error occurred while calling the endpoint:" -ForegroundColor Red
        Write-Host "Server response (JSON):"
        Write-Host $jsonString
        Write-Host ""
        Write-Host "For more information on how to resolve this issue, try generating the Terraform configuration from the Instaclustr Console under Settings > Cluster Resources > Terraform > Download."
        Write-Host ""
        exit 1
    }
}
catch {
    Write-Host "Exception: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}
finally {
    # Dispose of the client and handler if they were created.
    if ($client) { $client.Dispose() }
    if ($handler) { $handler.Dispose() }
}

cmd /c rmdir /s /q "$DEST_FOLDER_NAME" #need to do this instead of Remove-Item or other methods to deal with symlinks to terraform provider

Expand-Archive -Path "$ZIP_FILE_NAME" -DestinationPath "$DEST_FOLDER_NAME"

Remove-Item "$ZIP_FILE_NAME"

Push-Location "$DEST_FOLDER_NAME"

terraform init

./import-all.ps1

Pop-Location

Write-Host "Script execution completed."
