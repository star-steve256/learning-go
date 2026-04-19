package importexport

import "fmt"

// Capitalized - can be imported by other files
// non-capitalized - local to only

// local - accessible only in this file
var localVersion = "l0.1.0"

// exported - can be imported by importexport.Version
var ExportedVersion = "E0.1.0"

/* however, exported funcs that contain
   local var can modify the local var behaviours
*/

// in this case this acts like "getter, setter" in Java
// does not expose the var directly, only with funcs
func GetLocalVersion() string {
	return localVersion
}

func SetLocalVersion(input string) {
	localVersion = input
}

func PrintVersions() {
	fmt.Println("============================")
	fmt.Println("Exported version:", ExportedVersion)
	fmt.Println("Local version:", localVersion)
	fmt.Println("============================")
}
