# chainsaw
> Web Auditor

### Get
- Execute Binary: [Release Page](https://github.com/nic329/chainsaw/releases)
- Source Code `git clone --depth=1 https://github.com/nic329/chainsaw.git`

### Build (Source Code)
- Linux ```CGO_ENABLED=0 go build -o build/chainsaw```
- Windows ```CGO_ENABLED=0 go build -o build/chainsaw.exe```

### Run
- Linux or Mac
```
./chainsaw http://example.com/
./chainsaw -f url.txt
```
- Windows
```
chainsaw.exe http://example.com/
chainsaw.exe -f url.txt
```
