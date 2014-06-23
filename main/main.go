package main

import (
   "fmt"
   "kstar"
)

func main() {
   sequence := "CAAACAAAGCAATTGCAT"
   var i kstar.Index
   i = *kstar.NewIndex([]byte(sequence))
      
   snp:= []string{"G", "GCT"}
   i.AddSNP([]byte(sequence), snp, 4)

   fmt.Println(i)
}
