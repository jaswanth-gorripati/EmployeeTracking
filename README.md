# FilesMetaStore

> Note : Before starting this project make sure below softwares are installed 
- Golang
- Docker
- Docker-compose
- Nodejs v12 or higher

### Pre requirements

- Download the binaries 
```bash
./prereq.sh
```

- Start the Network and deploy the chaincode into the channel

```bash
./startFabric.sh
```

- To run the apis

```bash
cd api
npm i 
node registerOrg1User.js
node server.js
```

Now open http://localhost:3000/api-docs  in your browser for swagger ui


- To stop the network
```bash
./networkDown.sh
```


