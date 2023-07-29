# GoCiCd


### Configure Settings File

```sh
echo '{
    "userRegisterContainer": "<userDockerHub>",
    "appName": "<TagAppName>",
    "version": "latest",
    "containerName": "<containerName>",
    "shTest": "",
    "hcl":"./scripts/deploy.hcl"
    }' > .gocid-example.json
```