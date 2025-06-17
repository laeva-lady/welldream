# Welldream

> [!WARNING]
> This project only supports Hyprland

Log and display your app usage

![Example usage](imgs/example_usage.png)

## Usage
- `welldream -d` start the daemon
- `welldream` show the usage of the current day

`--debug` can be used to show debug info

## Cache
It keeps cache here `~/.cache/welldream`

## Build
### First option
Compile the project:
```bash
go build
```
you can then move the `welldream` binary wherever you want

### Second option
Use the Makefile:
```
make
```
the binary will be put in `$HOME/.local/bin/`



## todo
 - [ ] add weekly report
 - [ ] add monthy report
 - [x] add sorting
 - [ ] change sorting to sort by most used app instead of alphabetically
