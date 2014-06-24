package readfile

import (
     "os"
    "fmt"
    "bufio"
    "bytes"
)

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
