Step 1: download dependency
```shell
go get -u github.com/spf13/cobra/cobra
```

Step 2: initialization
```shell
cobra init cmd-example --pkg-name ./cmd-example
```

Step 3: New command
```shell
cobra add insert
cobra add delete
cobra add list
```

Step 4: Check
```shell
go build -o tt

tt insert 
#:insert called
```
