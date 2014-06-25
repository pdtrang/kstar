package main

import (
    "fmt"
    "flag"
    "readfile"
    "strings"
    "strconv"
    "kstar"
)



func main () {
    var genome_file = flag.String("g", "", "reference genome file")
    var snp_file = flag.String("s", "", "snp profile file")
    flag.Parse()

    sequence := readfile.ReadFASTA(*genome_file)   
    snp_f := readfile.ReadVCF(*snp_file)
    
    fmt.Println(string(sequence))
    //fmt.Println(len(sequence))

    /*fmt.Println("print in main")
    fmt.Println(len(snp_f))
    for i := 0; i < len(snp_f); i++ {
        fmt.Println(snp_f[i])
    }*/

    pos := make([]int, len(snp_f))
    snp_a := make([][]string, len(snp_f))
    for i := 0; i < len(snp_f); i++ {
    	a := strings.Split(string(snp_f[i]), "\t")	
    	pos[i],_ = strconv.Atoi(a[0])
    	snp_a[i] = strings.Split(a[1], "/")
    }

    
    //for i := 0; i < len(snp_f); i++ {
    //	fmt.Println("snp pos ", pos[i], "snp ", snp_a[i])	
    //}

    var i kstar.Index
   	i = *kstar.NewIndex([]byte(sequence))

   	for j := 0; j < len(snp_f); j++ {
   		i.AddSNP([]byte(sequence), snp_a[j], pos[j])	
   	}
   	
    fmt.Println(i)
}
