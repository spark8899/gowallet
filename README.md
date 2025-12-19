# gowallet
golang wallet tools

# build
```
go build -o gowallet main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gowallet_linux main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gowallet_win main.go
```

# running

## generate private
create private key
```
./gowallet genPrivateKey
0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947:0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
./gowallet genPrivateKey -n 10 # Create 10 private keys
xxxxx
..
```

get public key from private key
```
./gowallet getPublicKey -k 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
0x046005c86a6718f66221713a77073c41291cc3abbfcd03aa4955e9b2b50dbf7f9b6672dad0d46ade61e382f79888a73ea7899d9419becf1d6c9ec2087c1188fa18
```

get address from private key
```
./gowallet getAddress -k 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
```

## generate mnemonic
create mnemonic
```
./gowallet genMnemonic
tag volcano eight thank tide danger coast health above argue embrace heavy
./gowallet genMnemonic -s 24
xxx xx xxx ..
```

get seed from mnemonic
```
./gowallet getSeed -m "tag volcano eight thank tide danger coast health above argue embrace heavy"
efea201152e37883bdabf10b28fdac9c146f80d2e161a544a7079d2ecc4e65948a0d74e47e924f26bf35aaee72b24eb210386bcb1deda70ded202a2b7d1a8c2e
```

get path from mnemonic
```
./gowallet getPath -m "tag volcano eight thank tide danger coast health above argue embrace heavy" -p "m/44'/60'/0'/0/0"
0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947:0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
```

get path from seed
```
./gowallet getPath -s "efea201152e37883bdabf10b28fdac9c146f80d2e161a544a7079d2ecc4e65948a0d74e47e924f26bf35aaee72b24eb210386bcb1deda70ded202a2b7d1a8c2e" -p "m/44'/60'/0'/0/0"
0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947:0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
```

# update pkg

```
curl -X POST https://proxy.golang.org/github.com/spark8899/gowallet/@v/v1.0.0.info
curl -X POST https://pkg.go.dev/fetch/github.com/spark8899/gowallet@v1.0.0
```
