/*
   qgram wildcard index
*/
package kstar

import (
   "fmt"
   "math"
   //"os"
   //"bufio"
)

//var K = 3

type Index struct {
   Q int
   N int
   Qgram [][]int
   SNP []bool
}

func (i Index) GetQgram() [][]int {
    return i.Qgram
}

func FindQuery(query string, i_qgram [][]int, K int) []int{
   q := make([]int, len(query)-K+1)
   
   for j := 0; j < len(query)-K+1; j++ {
      r, a := Acgt(j, K, []byte(query))
      if a {
         q[j] = r
      } 
   }

   //fmt.Println("\n\n", q)

   q_index := make([][]int, len(q))
   for j := 0; j < len(q); j++ {
      //fmt.Println(i_qgram[q[j]]) 
      for k := 0; k < len(i_qgram[q[j]]); k++ {
         tmp := i_qgram[q[j]][k]
         q_index[j] = append(q_index[j], tmp)    
      }   
   }
   //fmt.Println("\n\n", q_index)
    
   for j := 0; j < len(q_index); j++ {
      for k := 0; k < len(q_index[j]); k++ {
         q_index[j][k] = q_index[j][k] - j
      }     
   }

   //fmt.Println("\n\n", q_index)

   /*for j := 0; j < len(q); j++ {
      fmt.Print(i_qgram[q[j]])
   }*/

   pos := q_index[0]
   for j := 1; j < len(q_index); j++ {
      pos = Intersect(pos, q_index[j])
   }

   return pos
}

func Intersect(a []int, b []int) []int{
   //fmt.Println("Intersection")
   i := 0
   j := 0
    
   var c []int

   for (i < len(a) && j < len(b)){

      if (a[i] == b[j]) {
         c = append(c, a[i])
         i = i+1
         j = j+1
      } else{
         if a[i] < b[j] {
            i = i+1
         } else{
            if a[i] > b[j] {
               j = j+1    
            }                
         }
      }
   }

    return c
}

func Acgt(i int, K int, sequence []byte) (int, bool){
   repr := 0
      acgt := true
      for j:=i; j<i+K && acgt; j++ {
         switch sequence[j] {
            case 'A': repr = 4*repr
            case 'C': repr = 4*repr + 1
            case 'G': repr = 4*repr + 2
            case 'T': repr = 4*repr + 3
            default:
               // we skip any qgram that contains a non-standard base, e.g. N
               acgt = false
         }
      }
      
   return repr, acgt
}

func NewIndex(sequence []byte, K int) *Index {
   //memstats := new(runtime.MemStats)
   //runtime.ReadMemStats(memstats)
   //log.Printf("\t BEGIN NewIndex: memstats:\t%d\t%d\t%d\t%d\t%d", memstats.Alloc, memstats.TotalAlloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
   //path := "test_out.index"
   /*file, err := os.Create(path)
   if err != nil {
      fmt.Println("error created file")
   }
   defer file.Close()

   w := bufio.NewWriter(file)*/

   idx := new(Index)
   idx.Q = K
   idx.N = len(sequence)
   idx.SNP = make([]bool, len(sequence))
   idx.Qgram = make([][]int, int(math.Pow(4.0, float64(K))))
   for i:=0; i<idx.N-K+1; i++ {
      r, a := Acgt(i, K, sequence)

      if a {
         idx.Qgram[r] = append(idx.Qgram[r], i)
         //fmt.Printf("%s=%d.  Store %d at location %d\n", string(sequence[i:i+K]), r, i, r)
         //fmt.Printf("%d\t%d\n", i, r)
         //fmt.Fprintln(w, i, r)
      } 
   }

   //w.Flush()

   return idx
}

func (idx Index) AddSNP(sequence []byte, snp []string, pos int, K int) {
   snp_len := make([]int, len(snp))

   for i := 0; i < len(snp); i++ {
      snp_len[i] = len(snp[i])
      //fmt.Println(snp_len[i])
   }

   for m := 0; m<len(snp); m++ {
      s_pos := pos - (K-snp_len[m])
      if s_pos < 0 {
         s_pos = 0
      }
      for i := s_pos; i <= pos; i++ {
         if (i < (len(sequence)-K+1)){
         seq := ""
         fmt.Println("SNP pos = ", pos)
         for j := 0; j <= (K-snp_len[m]); j++ {
            
            if i+j == pos{
               //fmt.Println("i ", i, " + j ", j, " = pos ", pos)
               seq = seq + snp[m]
            }else{
                  seq = seq + string(sequence[i+j])   
                  //fmt.Println("i ", i, " j ", j, " pos ", pos)
            }
           
         }   
         fmt.Println(seq)
      
         r, a := Acgt(0, K, []byte(seq))
         //fmt.Println(r, a)
         /*repr := 0
         acgt := true
         for n:=0; n<K && acgt; n++ {
            switch seq[n] {
               case 'A': repr = 4*repr
               case 'C': repr = 4*repr + 1
               case 'G': repr = 4*repr + 2
               case 'T': repr = 4*repr + 3
               default:
                  // we skip any qgram that contains a non-standard base, e.g. N
                  acgt = false
            }
         }*/
         
         if a {
            idx.Qgram[r] = append(idx.Qgram[r], i)
            fmt.Printf("%s=%d.  Store %d at location %d\n", seq[0:K], r, i, r)
         }
        }
      }
      fmt.Println()
   }
}

