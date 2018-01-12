TIPS:

# Regex templates

```
import (
    "fmt"
    "os"
    "regexp"
)

func main() {
    dst := make([]byte, 0, 1024)
    source := `"matt"`
    re := regexp.MustCompile(`"(?P<name>[a-z]+)"`)
    tpl := "My name is ${name}."

    matches := re.FindStringSubmatchIndex(source)
    out := re.ExpandString(dst, tpl, source, matches)
    fmt.Fprintf(os.Stdout, "%s\n", out)
}
```

# Decorator

func decorator(f func(s string)) func(s string) {

	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func doSomething(s string) {
	fmt.Println(s)
}

func main() {

	decorator(doSomething)("Do something")

}