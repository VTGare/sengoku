# sengoku
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FVTGare%2Fsengoku.svg?type=small)](https://app.fossa.com/projects/git%2Bgithub.com%2FVTGare%2Fsengoku?ref=badge_small) [![PkgGoDev](https://pkg.go.dev/badge/github.com/VTGare/sengoku)](https://pkg.go.dev/github.com/VTGare/sengoku)

A simple SauceNAO Go wrapper. <br />
Inspired by **[Sagiri](https://github.com/ClarityCafe/Sagiri)**, a SauceNAO wrapper for NodeJS.

## Installation
Add ``sengoku`` to your ``go.mod`` file <br />
```
module github.com/x/y

go 1.15

require(
    github.com/VTGare/sengoku latest
)
```
or ``go get -u github.com/VTGare/sengoku``

## Examples
```go
func main() {
    sen := sengoku.NewSengoku(os.Getenv("SAUCENAO_APITOKEN"))
    uri := "https://imgur.com/someimage.png"
    
    //Search using an image URL with a default query.
    sauce := sen.Search(uri)

    //Search with a different query.
    sauce := sen.SearchWithConfig(uri, &sengoku.Config{DB: sengoku.Pixiv, Results: 1})
}
```
