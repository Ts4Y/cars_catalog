ROOT=../../
source variables.sh

cd $ROOT/docker
docker compose build
docker compose  up  --force-recreate --remove-orphans