package main

import (
	"fmt"
	"flag"
	"github.com/pdtrang/kstar"
	"github.com/vtphan/fmi"
)

var idx fmi.Index
var count_10 = 0
var count_20 = 0
var count_40 = 0

func AppendIfMissing(slice []int, i int) []int {
   for _, ele := range slice {
      if ele == i {
         return slice
      }
   }

   return append(slice, i)
}

func update_count(k int){

	if (k == 10){
		count_10 = count_10 + 1
	}else{
		if (k == 20){
			count_20 = count_20 + 1
		}else{
			if (k == 40){
				count_40 = count_40 + 1
			}
		}
	}
}

func count(sequence []byte, k int, pos []int, n int) ([]int) {

	var p []int
	
	L := len(sequence)
	if (n == 0){
		for i := 0; i < L; i++ {
			if (i+k <= L){
				var a = idx.Search(sequence[i:i+k])	
				
				//occ = 1
				if (len(a)==1){
					a[0] = a[0]+k-1
					update_count(k)
				}else{
					//occ > 1
					for j := 0; j < len(a); j++ {
						a[j] = a[j]+k-1
						p = AppendIfMissing(p, a[j])
					}
								
				}
			}
		}
	}else{
		for i := 0; i < len(pos); i++ {
			if (pos[i]+k <= L){
				var a = idx.Search(sequence[pos[i]:pos[i]+k])	
				
				//occ = 1
				if (len(a)==1){
					a[0] = a[0]+k-1
					update_count(k)					
				}else{
					//occ > 1
					for j := 0; j < len(a); j++ {
						a[j] = a[j]+k-1
						p = AppendIfMissing(p, a[j])
					}
								
				}
			}
		}
	}

	return p
}

func count_occ(sequence []byte, k int) {
	var p []int
	
	fmt.Println("Begin counting...")

	fmt.Println("10-mer")
	p = count(sequence, k, p, 0)
	
	fmt.Println("20-mer")
	p = count(sequence, 2*k, p, 1)
	
	fmt.Println("40-mer")
	p = count(sequence, 4*k, p, 1)

	fmt.Println("Finish counting")

}

func main(){
	var genome_file = flag.String("g", "", "reference genome file")
	var index_file = flag.String("i", "", "index file of genome")
	
	flag.Parse()

	var k = 10

	fmt.Println("Read FASTA")
  	sequence := kstar.ReadFASTA(*genome_file)

  	//fmt.Println(string(sequence))
  	fmt.Println(len(sequence))
    
	idx = *fmi.New(*genome_file)
	idx.Save(*index_file)

	fmt.Println("Finish indexing multigenome...")

	count_occ(sequence, k)

	fmt.Println("\nNumber of occ = 1 of ")
	fmt.Println("10-mer ", count_10, "\n20-mer ", count_20, "\n40-mer ", count_40)
}
