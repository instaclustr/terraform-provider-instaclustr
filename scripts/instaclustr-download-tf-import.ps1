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
    $handler = New-Object System.Net.Http.HttpClientHandler
    $handler.AutomaticDecompression = [System.Net.DecompressionMethods]::None

    $client = New-Object System.Net.Http.HttpClient($handler)

    $client.DefaultRequestHeaders.Authorization = 
        [System.Net.Http.Headers.AuthenticationHeaderValue]::new("Basic", $basicAuthHeaderValue)

    $response = $client.GetAsync($apiUrl).Result

    if ($response.IsSuccessStatusCode) {
        $zipBytes = $response.Content.ReadAsByteArrayAsync().Result
        [System.IO.File]::WriteAllBytes($DEST_FILE_NAME, $zipBytes)
        Write-Host ""
        Write-Host "Terraform files downloaded to '$DEST_FILE_NAME'"
        Write-Host ""
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
