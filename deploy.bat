call .\stop.bat
scp ".\docker-compose.yml" "%server_user%@%server_addr%:%deployment_path%/docker-compose.yml"
ssh "%server_user%@%server_addr%" "rm -rf '%deployment_path%/dst'"
scp -r ".\dst" "%server_user%@%server_addr%:%deployment_path%"