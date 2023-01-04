# M3O CLI

The command line for M3O

## Install

Quick install

```sh
curl -fssl https://install.m3o.com/cli | /bin/bash
```

From source

```
go install github.com/m3o/m3o-cli/cmd/m3o@latest
```

Otherwise download the latest [release](https://github.com/m3o/m3o-cli/releases/latest) binary.

## Usage

Display help

```
m3o -h
```

### API Token

Export your API token as found in the UI

```
export M3O_API_TOKEN=xxxxxx
```

### API Address

Set the API address

```
export M3O_API_ADDRESS=http://localhost:8080
```

### List services

```
m3o explore list
```

### Search services

```
m3o explore search --query=helloworld
```

### Query service

Usage

```
m3o [service] [endpoint] --[param]=value
```

Example
```
m3o helloworld call --name=Alice
```

### Client Call

```
$ m3o client call helloworld Call '{"name": "Alice"}'
{
 "message": "Hello Alice"
}
```

### Client Stream

```
m3o client stream notes Subscribe '{}'
```


