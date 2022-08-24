<p align="center">
	<a href="https://discord.gg/TBR9bRjd6Z">
		<img src="https://discordapp.com/api/guilds/861917584437805127/widget.png?style=banner2" alt="Discord Banner"/>
	</a>
</p>

---

# M3O CLI

The command line for M3O

## Install

Quick install

```sh
curl -fssl https://install.m3o.com/cli | /bin/bash
```

From source

```
go get github.com/m3o/m3o/cmd/m3o
```

Otherwise download the latest [release](https://github.com/m3o/m3o/releases/latest) binary.

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


