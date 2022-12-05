# DEX-TRADE Trading API for Golang.

## Installation

``` bash
go get github.com/daefrom/go-dextrade
```

## Usage

### Basic requests

``` go
package main

import (
	"fmt"
	"github.com/daefrom/go-dextrade/v1"
)

client := dextrade.NewClient()
	pairs := []string{"BTCUSDT", "DAEUSDT"}

	trades, err := client.Trades.Get(pairs)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(trades)
	}
```

### Authentication

``` go
client := dextrade.NewClient().Auth(key, secret)

```

### Order create

``` go
client := dextrade.NewClient().Auth(key, secret)

order, err := client.Order.Create("DAEUSDT", 3, 0.0043, "sell", "0")

if err == nil {
	fmt.Println(order)
}
```

## Testing

All integration tests are stored in `tests/integration` directory.

Run tests using:
``` bash
export DEXTRADE_API_KEY="api-key"
export DEXTRADE_API_SECRET="api-secret"
go test -v ./tests/integration
```

## Contributing

1. Fork it (https://github.com/daefrom/go-dextrade/fork)
2. Create your feature branch (`git checkout -b my-new-feature)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
