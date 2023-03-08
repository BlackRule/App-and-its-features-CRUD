REM ssh -t %server_user%@%server_addr% "cd %deployment_path% && su root -c 'docker kill $(docker ps -q)'"
REM ssh -t %server_user%@%server_addr% "cd %deployment_path% && su root -c 'docker-compose stop'"
ssh -t %server_user%@%server_addr% "cd %deployment_path% && su root -c 'docker update --restart=unless-stopped $(docker container ls -q)'"
ssh -t %server_user%@%server_addr% "cd %deployment_path% && su root -c 'docker container stop $(docker container ls -q)'"