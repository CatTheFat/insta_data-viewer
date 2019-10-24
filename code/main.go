package main

import (
	"fmt"
)

func main() {
	endProg := false

	fmt.Println("Insta DM json parser v1.0")

	openf()

	for !endProg {
		var icmd string

		fmt.Println(`Enter cmd: (help, list, export, fetch, closef, closep)`)
		fmt.Scanln(&icmd)

		switch icmd {
		case "help":
			//---//
			fmt.Println("")
			fmt.Println("Showing help: [cmd] - function")
			fmt.Println("[list] - lists all the participants in the file with their id ([id] - @tag)")
			fmt.Println("[export] - exports all conversations")
			fmt.Println("[fetch] - exports a single convarsation using it's id")
			fmt.Println("[closef] - closes the current file and promts you to open another one")
			fmt.Println("[closep] - closes the current file and terminates the program")
			fmt.Println("")
			//---//
		case "list":
			//---//
			List()
			//---//
		case "export":

			fmt.Println("")
			ExportAll()
			fmt.Println("")
			fmt.Println("Done!!")
			fmt.Println("")

		case "fetch":
			var inID int
			fmt.Scanln(&inID)
			ExportConv(inID, true, "C:\\MaPath\\files\\")

		case "closef":
			//---//
			fmt.Println("")
			openf()
			//---//
		case "closep":
			//---//
			fmt.Println("")
			fmt.Println("Closing program...")
			endProg = true
			//---//
		default:
			//---//
			fmt.Println("")
			fmt.Printf("Unknown cmd %q\n\n", icmd)
			//---//
		}

	}

}
