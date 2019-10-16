param (
    [Parameter(Mandatory = $false)][ValidateSet("true", "false")][string]$buildFrontend = "false"
)

# if($arg)
IF ($buildFrontend -eq $true) {
    Write-Output "Building frontend..."
    Set-Location frontend
    npm run build -quiet
    IF ($LASTEXITCODE -eq 0) {
        Write-Output "Packing frontend..."
        Write-Output $(Get-Location)
        go-bindata-assetfs.exe -pkg=frontend -nocompress=false build/...
    }
    else {
        Write-Output "Error on build frontend"
        $LASTEXITCODE = 1
    }
    Set-Location ../
}

if (($buildFrontend -ne $true) -or ($LASTEXITCODE -eq 0)) {
    go install -v; if ($LASTEXITCODE -eq 0) { go-reactjs.exe }
}