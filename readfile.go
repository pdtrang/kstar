File Edit Options Buffers Tools Help                                                                                                                                                                                                         
package kstar

import (
     "os"
    "fmt"
    "bufio"
    "bytes"
    //"log"                                                                                                                                                                                                                                  
    "strings"
    "strconv"
    "sort"
//      "runtime"                                                                                                                                                                                                                            
//      "log"                                                                                                                                                                                                                                
)

type SNP struct{
    profile []string
}

func (f SNP) GetString() []string {
    return f.profile
}

func SaveIndex2(i_qgram [][]int, prefix []string, K int, path string){

//      memstats := new(runtime.MemStats)                                                                                                                                                                                                    
//      runtime.ReadMemStats(memstats)                                                                                                                                                                                                       
//      log.Printf("Save Index 2: Alloc %d\tTotalAlloc %d\tSys %d\tHeapAlloc %d\tHeapSys %d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)                                                       
        file, err := os.Create(path)
        if err != nil {
                fmt.Println("error created file")
        }
        defer file.Close()

        w := bufio.NewWriter(file)

        fmt.Fprintln(w, prefix[0], ", K =", K)
        for j := 0; j < len(i_qgram); j++ {
                if(len(i_qgram[j])>0){
                        fmt.Fprint(w, j, "\t")
                        for k := 0; k < len(i_qgram[j]); k++ {
                                fmt.Fprint(w, i_qgram[j][k], " ")
                        }
                        fmt.Fprint(w, "\n")
                }

        }

//      runtime.ReadMemStats(memstats)                                                                                                                                                                                                       
//      log.Printf("Save Index 2: Alloc %d\tTotalAlloc %d\tSys %d\tHeapAlloc %d\tHeapSys %d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)                                                       

        w.Flush()

}

func SaveIndex(i_qgram [][]int, prefix []string, K int, path string){
//      memstats := new(runtime.MemStats)                                                                                                                                                                                                    
//      runtime.ReadMemStats(memstats)                                                                                                                                                                                                       
//      log.Printf("Save Index 1: Alloc %d\tTotalAlloc %d\tSys %d\tHeapAlloc %d\tHeapSys %d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)                                                       

-UU-:----F1  readfile.go    Top L1    Git:!  (Go)--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Loading vc-git...done
