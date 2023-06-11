# nomad job run ./scripts/deployment/app.hcl
# nomad job run -var image=rodrigoresende/oli-bff-sdk:1.0 ./scripts/deployment/app.hcl

nomad job run -var image=$1 ./lib/cli/scripts/deployment/app.hcl

