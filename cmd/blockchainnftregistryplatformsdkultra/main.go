// cmd/blockchainnftregistryplatformsdkultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistryplatformsdkultra/internal/blockchainnftregistryplatformsdkultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistryplatformsdkultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
