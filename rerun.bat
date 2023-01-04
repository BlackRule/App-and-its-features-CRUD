ssh -t %server_user%@%server_addr% "cd %deployment_path% && su root -c 'docker kill $(docker ps -q); docker-compose up'"
