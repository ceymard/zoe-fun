package main

import (
	"log"
	"os"

	zoe "github.com/ceymard/zoe/compiler"
)

func main() {

	for _, fname := range os.Args[1:] {

		log.Print(`Handling '`, fname, `'`)
		ctx, err := zoe.NewZoeContext(fname)
		if err != nil {
			log.Print(err)
			continue
		}

		// for iter := lx.FirstToken; iter != nil; iter = iter.Next {
		// 	if iter.Kind == zoe.TK_WHITESPACE {
		// 		continue
		// 	}
		// 	pp.Println(iter.KindStr(), iter.Value(b), iter.Line, iter.Column)
		// }

		res := ctx.ParseFile()
		// w := &zoe.ZoeWriter{Writer: os.Stdout}
		// res.WriteTree(w, "")
		_, _ = os.Stdout.WriteString(res.Debug() + "\n")

		for _, err := range ctx.Errors {
			err.Print(os.Stderr)
		}
	}
}
