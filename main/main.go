package main

import (
   "fmt"
   "flag"
   "kstar"
   "readfile"
)

func main() {
	var genome_file = flag.String("g", "", "reference genome file")
    flag.Parse()

    sequence := string(readfile.ReadFASTA(*genome_file))
    
    //fmt.Println(sequence)
    //fmt.Println(len(sequence))
   
	var i kstar.Index
   	i = *kstar.NewIndex([]byte(sequence))
      
   	snp:= []string{"G", "GCT"}
   	i.AddSNP([]byte(sequence), snp, 4)

   	fmt.Println(i)
}
