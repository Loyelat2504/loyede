package main
func RepeatAlpha (s string) string{
	result := ""
	for_, char := range s {
		if char >= 'a' && char <= 'z' {
			repeatcount := int(char - 'a' + 1)
		}
		for i := 0 ; i < repeatcount ; i++ {
			result += string (char)
		}
	} else if char >= 'A' && char <= 'Z' { 
		repeatcount := int(char - 'A' + 1)	                           

	}  
	    for i := 0 ; i < repeatcount ; i++ {
			result += string (char)
		}

	}
