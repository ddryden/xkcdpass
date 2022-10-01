# xkcdpass
![correct horse battery staple](https://imgs.xkcd.com/comics/password_strength.png)

https://xkcd.com/936/
Generates password from local dictionary on Mac and Linux

## Install
`go install github.com/ddryden/xkcdpass@latest`

## Usage of `xkcdpass`:
*  -dict string
    Path to a custom dictionary (default "/usr/share/dict/words")
*  -max int
    maximum word length (default 9)
*  -min int
    minimum word length (default 5)
*  -ws string
    String to separate words (default " ")
