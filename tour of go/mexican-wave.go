package main

import (
    "fmt"
	"strings"
)

func wave(words string) (s []string) {
	// Your code here and happy coding!
  
	s = []string{};

	if words == "" {
		return s;
	}

	for i := 0; i < len(words); i++ {

		if (string (words[i]) == " ") {
			continue;
		}; 

		c := strings.ToUpper(string (words[i]))
		pos := i;

		left := words[:pos]; 
		right := words[pos + 1:]; 


		w := left + c + right;

		s = append(s, w); 


	}

	if len(s) == 0 {
		return nil;
	}

	return s;

}


func main()  {
	
	fmt.Println(wave(" xd dthrjr ")); 
}