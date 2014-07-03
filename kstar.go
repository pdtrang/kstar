/*
   qgram wildcard index
*/
package kstar

import (
   "fmt"
   "math"
   "os"
   "bufio"
)

var K = 3

type Index struct {
   Q int
   N int
   Qgram [][]int
   SNP []bool
}

func (idx Index) acgt(i int, K int, sequence []byte) (int, bool){
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

func NewIndex(sequence []byte, path string) *Index {

   file, err := os.Create(path)
   if err != nil {
      fmt.Println("error")
   }
   defer file.Close()

   w := bufio.NewWriter(file)

   idx := new(Index)
   idx.Q = K
   idx.N = len(sequence)
   idx.SNP = make([]bool, len(sequence))
   idx.Qgram = make([][]int, int(math.Pow(4.0, float64(K))))
   for i:=0; i<idx.N-K+1; i++ {
      r, a := idx.acgt(i, K, sequence)

      if a {
         idx.Qgram[r] = append(idx.Qgram[r], i)
         //fmt.Printf("%s=%d.  Store %d at location %d\n", string(sequence[i:i+K]), r, i, r)
         fmt.Fprintln(w, r)
      }
   }
   
   w.Flush()
   
   return idx
}

func (idx Index) AddSNP(sequence []byte, snp []string, pos int) {
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
      
         r, a := idx.acgt(0, K, []byte(seq))
         
         if a {
            idx.Qgram[r] = append(idx.Qgram[r], i)
            fmt.Printf("%s=%d.  Store %d at location %d\n", seq[0:K], r, i, r)
         }
        }
      }
      fmt.Println()
   }
}

