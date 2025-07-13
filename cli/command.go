package cli

type Command struct {
	List struct {
		Recursive bool `help:"Recursively search all sub directories."`
		WordCount bool `help:"To get 10 most frequent words."`

		Path string `arg:"" optional:"" name:"path" help:"path of directory to list." type:"path"`
	} `cmd:"" help:"lists all text files with no of lines."`
}