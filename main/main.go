package main

import (
"flag"
"github.com/pdtrang/kstar"
"os"
"fmt"
"bufio"
"time"
"math"
"runtime"
"log"
)

var min = 5
var max = 15

func main () {
    var genome_file = flag.String("g", "", "reference genome file")
    //var snp_file = flag.String("s", "", "snp profile file")
    var output_index = flag.String("i", "", "output index file")
    var output2 = flag.String("o", "", "output index 2")
    var logfile = flag.String("l", "", "logfile")
    flag.Parse()

    file, err := os.Create(*logfile)
    if err != nil {
            fmt.Println("error created file")
    }
    defer file.Close()

    wr := bufio.NewWriter(file)


    begin := time.Now()

    memstats := new(runtime.MemStats)
    runtime.ReadMemStats(memstats)
    log.Printf("begin main: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    fmt.Fprintln(wr, "begin main: memstats: Alloc ", memstats.Alloc, "\tTotal Alloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys", memstats.HeapSys)
    sequence := kstar.ReadFASTA(*genome_file)
    prefix, err := kstar.ReadLines(*genome_file)
    if err != nil {
            fmt.Println(prefix)
    }

    K := int(math.Log(float64(len(sequence)))/math.Log(4.0))
    //fmt.Println("K = ", math.Log(float64(len(sequence)))/math.Log(4.0) )                                                                                                                                                               
    //fmt.Println("K = ", int(math.Log(float64(len(sequence)))/math.Log(4.0)))                                                                                                                                                           

    if(K < min){
            K = min
    } else{
            if (K > max){
                    K = max
            }
    }

    //dbsnp := kstar.ReadVCF(*snp_file)
    
    fmt.Println("K = ", K)
    fmt.Println("Indexing.......")
    fmt.Fprintln(wr, "K = ", K)

    var i kstar.Index
    runtime.ReadMemStats(memstats)
    fmt.Fprintln(wr, "\nIndexing")
    log.Printf("begin Indexing: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    fmt.Fprintln(wr, "begin Indexing: memstats: Alloc ", memstats.Alloc, "\tTotal Alloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys", memstats.HeapSys)
    si := time.Now()
    i = *kstar.NewIndex([]byte(sequence), K)
    //fmt.Println(len(i.GetQgram()))                                                                                                                                                                                                     
    ei := time.Since(si)
    fmt.Println("Finished Indexing......")
    runtime.ReadMemStats(memstats)
    log.Printf("end Indexing: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    fmt.Fprintln(wr, "end Indexing: memstats: Alloc ", memstats.Alloc, "\tTotal Alloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys", memstats.HeapSys)
    fmt.Fprintln(wr, "Indexing running time ", ei)

    i_qgram := i.GetQgram()

   	/*for j := range dbsnp {
        snp := dbsnp[j].GetString()
    	i.AddSNP([]byte(sequence), snp[1:], j)	
   	}*/
   	
    fmt.Println("\nSave Index....")
    fmt.Fprintln(wr, "\nSave Index....")
    runtime.ReadMemStats(memstats)
    log.Printf("begin Save Index: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    fmt.Fprintln(wr, "begin Save Index: memstats: Alloc ", memstats.Alloc, "\tTotal Alloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys", memstats.HeapSys)
    sys := memstats.Sys
    ts1 := time.Now()
    kstar.SaveIndex(i_qgram, prefix, K, *output_index)
    te1 := time.Since(ts1)
    runtime.ReadMemStats(memstats)
    log.Printf("end Save Index: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    log.Printf("sys = %d", memstats.Sys - sys)
    fmt.Fprintln(wr, "end Save Index: memstats: Alloc ", memstats.Alloc, "\tTotal Alloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys", memstats.HeapSys)
    fmt.Fprintln(wr, "Save Index running time ", te1)

    fmt.Println("\nSave Index 2....")
    fmt.Fprintln(wr, "\nSave Index 2....")
    runtime.ReadMemStats(memstats)
    log.Printf("begin Save Index 2: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    fmt.Fprintln(wr, "begin Save Index 2: memstats: Alloc ", memstats.Alloc, "\tTotal Alloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys", memstats.HeapSys)
    sys = memstats.Sys
    ts2 := time.Now()
    kstar.SaveIndex2(i_qgram, prefix, K, *output2)
    te2 := time.Since(ts2)
    runtime.ReadMemStats(memstats)
    log.Printf("end Save Index 2: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    log.Printf("sys 2 = %d", memstats.Sys - sys)
    fmt.Fprintln(wr, "end Save Index 2: memstats: Alloc ", memstats.Alloc, "\tTotal Alloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys", memstats.HeapSys)
    fmt.Fprintln(wr, "Save Index2 running time ", te2)

    elapsed := time.Since(begin)
    fmt.Fprintln(wr, "\nTotal running time ", elapsed)
    fmt.Println("Total running time  ", elapsed)

    runtime.ReadMemStats(memstats)
    log.Printf("end main: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc,   memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
    fmt.Fprintln(wr, "end main: memstats: Alloc ", memstats.Alloc, "\tTotalAlloc ", memstats.TotalAlloc, "\tSys ", memstats.Sys, "\tHeapAlloc ", memstats.HeapAlloc, "\tHeapSys ", memstats.HeapSys)
    wr.Flush()
    
    query := "CAATCTGCA"
    //fmt.Println(query)
      
    pos := kstar.FindQuery(query, i_qgram, K)

    fmt.Println("Position of Read: ", pos)
}
