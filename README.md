# About

Rewrite of [Archon](https://github.com/adefreitas/archon) in Go. Composes frames from a group of given images, and generates a video out of them

## Background
Originally wrote Archon in Typescript for covenience being familiar with TypeScript, but decided to give Go a first try and rewrote the project in Go while picking it up along the way. I did not document the performance improvements but was able to generate more videos per hour than with the TypeScript implementation

## External dependencies

### ffmpeg

- MacOS:
  Run `brew install ffmpeg`
- Ubuntu
  Run `sudo apt-get install ffmpeg`
- Windows
  Run `?`

## Run

```
go run .
```

Remember to set your path like

```
export GOPATH=/home/${USER}/GO
```

## Disclaimer

Not for commercial use of any kind. Reach out for enquiries about licensing.

## Name inspiration

Archons from Star Craft! Merging two into one for extra power

<img width="398" alt="image" src="https://static.wikia.nocookie.net/starcraft/images/a/a1/Archon_SCR_HeadAnim.gif/revision/latest/scale-to-width-down/224?cb=20170728135257">
