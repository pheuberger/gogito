package subcommands

import (
	"fmt"

	"github.com/pheuberger/gogito/internal/paths"
)

func Init(path string) {
	absPath := paths.AbsFrom(path)
	fmt.Printf("Initialized empty Git repository in %s\n", absPath)
	panic("not implemented")
}
