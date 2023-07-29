# GoCiCd


### Add To .GitIgnore

```sh
echo '.gocid' >> .gitignore
```

### Configure Settings File In Initial Project

```sh
echo '{
    "userRegisterContainer": "<userDockerHub>",
    "appName": "<TagAppName>",
    "version": "latest",
    "containerName": "<containerName>",
    "shTest": "",
    "hcl":"./scripts/deploy.hcl"
    }' > .gocid.json
```

### Download Gocid In Inicial Project

*Download Executable arch linux-amd64*
```sh
wget -o ./.gocid  https://github.com/IsaacDSC/GoCiCd/blob/v0.1/dist/linux-amd64
```

*Download Executable arch linux-arm64*
```sh
wget -o ./.gocid  https://github.com/IsaacDSC/GoCiCd/blob/v0.1/dist/linux-arm64
```

*Download Executable arch windows-386*
```sh
wget -o ./.gocid  https://github.com/IsaacDSC/GoCiCd/blob/v0.1/dist/windows-386
```

*Download Executable arch windows-amd64*
```sh
wget -o ./.gocid  https://github.com/IsaacDSC/GoCiCd/blob/v0.1/dist/windows-amd64
```


### Execute Deploy

```sh
./gocid
```