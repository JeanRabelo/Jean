package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"sort"
)
func substituteString(a string, b string, c string) string{
	re, _:=regexp.Compile(b)
	final := re.ReplaceAllString(a, c)
	return final
}

func findFrequencyOfArray(arr []string) map[string]int{
   frequency := make(map[string]int)
   for _, item := range arr{
	if len(item) > 2{
		if frequency[item] == 0{
		 frequency[item] = 1
		} else {
		 frequency[item]++
		}
	}
   }
   return frequency 
}

func sortByValues(basket map[string]int){
	keys := make([]string, 0, len(basket))

	for key := range basket {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool{
		return basket[keys[i]] > basket[keys[j]]
	})

	for i, item := range keys{
		if i < 10 {
			fmt.Printf("%d - %s ocorre %d vezes.\n", i+1, item, basket[item])
		}
	}
}


func main(){
	data, _ := os.ReadFile("A Riqueza das Nacoes - Adam Smith.txt")
	conteudo_array := strings.Split(string(data), "\n")
	conteudo := strings.Join(conteudo_array, " ") 
	//conteudo := "blablabla, blablabla, . ! sadfsadf???, gota d'água, Abcd"
	conteudo = substituteString(conteudo, "[0-9]", " ")
	conteudo = substituteString(conteudo, "-", "")
	conteudo = substituteString(conteudo, `[^\wçáÁàÀéÉíÍóÓúÚüÜ]`, " ")
	conteudo = strings.ToLower(conteudo)
	word_list := strings.Fields(conteudo)
	frequency := findFrequencyOfArray(word_list)
	sortByValues(frequency)
}
