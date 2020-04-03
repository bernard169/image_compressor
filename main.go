// compressor algorithm :
//		first : list all bytes of the image
//		second : for each different byte, find how many occurence it has
// 		third : sort the list in increasing occurences value order
//      fourth : two by two, set the different bytes as leaves of a tree
// 				structure, starting by those with most occurences
//		fifth : attribute values to the branches directly above the leaves
// 				(bytes) : the leaf

package main

func main() {
	/*app := &cli.App{
			Name: "greet",
			Usage: "say a greeting",
			Action: func(c *cli.Context) error {
				fmt.Println("Greetings")
				return nil
		},
		}

	app.Run(os.Args) */
}
