# Templater
CLI for filling template files using go-templates syntax.

## Usage

You can find an example in the examples catalog.

```shell
Usage of templater:
  -input string
    	Input file
  -output string
    	Output file
  -var value
    	Thempate vars. ex: -var title='hello' -var body='my friend'
    	
# Fill template
templater -input server.yaml.tmpl -output server.yaml \
          -var environment="production" \
          -var addr="192.168.0.1" \
          -var port="8080" \
          -var password="MyVeryLongCoolPassword123"

# Encode a var value to base64 string
# just add "b64:" prefix to your value
... -var password="b64:MyVeryLongCoolPassword123"
```
## Install 
```shell
go install github.com/yousysadmin/templater/cmd/templater
```
## Build
```shell
CGO_ENABLED=0 go build -v -ldflags='-s -w -X "main.version=0.0.1"' -trimpath -o dist/templater ./cmd/templater
```


