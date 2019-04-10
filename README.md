# GoSearch
Search text through files in given folder
* cmd
* gui - build with gotron - https://github.com/Equanox/gotron

## CMD

<img src="https://raw.githubusercontent.com/FilipAnteKovacic/gosearch/master/docs/cmd-search_res.png" alt="cmd" />

### Install

```
clone project
```

```
go build -o ./gosearch
```

```
mv /usr/local/bin gosearch
```

### Usage

Flags
* -f="{folder}" - search specific folder - default current folder
* -l="txt"      - generate log file with results

Simple search in current folder
```
gosearch {text}
```

Search given folder and log in txt file
```
gosearch -f="{folder}" -l="txt" {text}
```

## GUI-GOTRON

<img src="https://raw.githubusercontent.com/FilipAnteKovacic/gosearch/master/docs/gui-search_res.png" alt="cmd" />

### Install
Install gotron-builder 
* https://github.com/Equanox/gotron

```
clone project
```

```
go build -o ./gosearch
```

```
./gosearch
```

```
gotron-builder --win     // windows (wine required)
gotron-builder --macos   // macos
gotron-builder --linux   // linux
```

# Tasks
- [ ] Search through other formats ( pdf, excel, word )
- [ ] Search file names & folders
- [ ] Log files in other formats 
- [ ] Casesensitive search



