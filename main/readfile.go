package readfile

import (
     "os"
    "fmt"
    "bufio"
    "bytes"
    "log"
)

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func ReadVCF(vcf_file string) []string{
    lines, err := readLines(vcf_file)
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }
    //for i, line := range lines {
    //    fmt.Println(i, line)
    //}

    //fmt.Println("read VCF")
    //fmt.Println(lines)
    //for i := 0; i < len(lines); i++ {
    //    fmt.Println(lines[i])
    //}

    
    return lines
}

func ReadFASTA(sequence_file string) []byte {
    f,err := os.Open(sequence_file)
    if err != nil{
        fmt.Printf("%v\n",err)
        os.Exit(1)
    }

    defer f.Close()
    br := bufio.NewReader(f)
    byte_array := bytes.Buffer{}

    //line , err := br.ReadString('\n')
	_ , isPrefix, err := br.ReadLine()
	if err != nil || isPrefix{
		fmt.Printf("%v\n",err)
		os.Exit(1)
	}
    //fmt.Printf("%s",line)

    for {
        line , isPrefix, err := br.ReadLine()
        if err != nil || isPrefix{
            break
        } else {
            byte_array.Write([]byte(line))
        }
    }
    //byte_array.Write([]byte("$"))
    input := []byte(byte_array.String())
    return input
}
