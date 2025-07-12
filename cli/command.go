package cli

var Command struct {
	List struct {
		Recursive bool `help:"Recursively search all sub directories."`

		Path string `arg:"" name:"path" help:"path of directory to list." type:"existingdirectory"`
	} `cmd:"" help:"lists all text files with no of lines."`
}