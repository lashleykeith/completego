package main

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(eb englishBot){
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot){
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string{
	// VERy custom logic for generating an english greeting
	return "Hi there!"

}

func (spanishBot) getGreeting() string{
	// VERy custom logic for generating a spanish greeting
	return "Hello!"

}

// an error because you cannot repeat this function