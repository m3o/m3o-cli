# M3O CLI

The command line for M3O

## Install

Building from source

```
go get github.com/m3o/m3o-cli/cmd/m3o
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

### List services

```
m3o explore list
```

### Search services

```
m3o explore search --query=db
```

### Call service

```
$ m3o call helloworld Call '{"name": "Alice"}'
{
 "message": "Hello Alice"
}
```

### Stream response

```
m3o stream notes Subscribe '{}'
```


