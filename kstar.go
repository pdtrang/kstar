/*
   qgram wildcard index
*/
package kstar

import (
   "fmt"
   "math"
)

var K = 3

type Index struct {
   Q int
   N int
   Qgram [][]int
   SNP []bool
}

func NewIndex(sequence []byte) *Index {
   idx := new(Index)
   idx.Q = K
   idx.N = len(sequence)
   idx.SNP = make([]bool, len(sequence))
   idx.Qgram = make([][]int, int(math.Pow(4.0, float64(K))))
   for i:=0; i<idx.N-K+1; i++ {
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
      if acgt {
         idx.Qgram[repr] = append(idx.Qgram[repr], i)
         fmt.Printf("%s=%d.  Store %d at location %d\n", string(sequence[i:i+K]), repr, i, repr)
      }
   }
   return idx
}

func (idx Index) AddSNP(sequence []byte, snp []string, pos int) {
   fmt.Println("addSNP")

   //fmt.Println(len(snp))
   snp_len := make([]int, len(snp))

   for i := 0; i < len(snp); i++ {
      snp_len[i] = len(snp[i])
      //fmt.Println(snp_len[i])
   }

   for m := 0; m<len(snp); m++ {
      s_pos := pos - (K-snp_len[m])
      for i := s_pos; i <= pos; i++ {
        seq := ""
         for j := 0; j <= (K-snp_len[m]); j++ {
            if i+j == pos{
               seq = seq + snp[m]
            }else{
               seq = seq + string(sequence[i+j])   
            }
           
            
        }
        fmt.Println(seq)

        repr := 0
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
         }
         if acgt {
            idx.Qgram[repr] = append(idx.Qgram[repr], i)
            fmt.Printf("%s=%d.  Store %d at location %d\n", seq[0:K], repr, i, repr)
         }
        
      }
      fmt.Println()
   }
}

