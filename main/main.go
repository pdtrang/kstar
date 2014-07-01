package main

import (
    "fmt"
    "flag"
    //"readfile"
    "github.com/pdtrang/kstar"
)

func main () {
    var genome_file = flag.String("g", "", "reference genome file")
    var snp_file = flag.String("s", "", "snp profile file")
    flag.Parse()

    sequence := readfile.ReadFASTA(*genome_file)   
    dbsnp := readfile.ReadVCF(*snp_file)
    
    fmt.Println(string(sequence))
    fmt.Println(len(sequence))
    
    var i kstar.Index
   	i = *kstar.NewIndex([]byte(sequence))

   	for j := range dbsnp {
        snp := dbsnp[j].GetString()
    	i.AddSNP([]byte(sequence), snp[1:], j)	
   	}
   	
    fmt.Println(i)
}
