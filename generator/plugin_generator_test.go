package generator

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/bilginyuksel/cordova-plugin-helper/reader"
)

func TestCreatePluginWithGivenName_ExpectPluginFilesCreated(t *testing.T) {
	CreateBasePlugin(".", "group", "project")

	linuxFileList := []string{"cordova-plugin-group-project/README.md", "cordova-plugin-group-project/tsconfig.json",
		"cordova-plugin-group-project/scripts/util.ts", "cordova-plugin-group-project/scripts/project.ts",
		"cordova-plugin-group-project/src/main/java/com/group/cordova/project/project.java",
		"cordova-plugin-group-project/types/util.d.ts", "cordova-plugin-group-project/www/project.js",
		"cordova-plugin-group-project/types/project.d.ts", "cordova-plugin-group-project/www/util.js",
		"cordova-plugin-group-project/www/project.js.map", "cordova-plugin-group-project/www/util.js.map",
		"cordova-plugin-group-project/package.json", "cordova-plugin-group-project/plugin.xml"}

	winFileList := []string{"cordova-plugin-group-project\\README.md", "cordova-plugin-group-project\\tsconfig.json",
		"cordova-plugin-group-project\\scripts\\util.ts", "cordova-plugin-group-project\\scripts\\project.ts",
		"cordova-plugin-group-project\\src\\main\\java\\com\\group\\cordova\\project\\project.java",
		"cordova-plugin-group-project\\types\\util.d.ts", "cordova-plugin-group-project\\www\\project.js",
		"cordova-plugin-group-project\\types\\project.d.ts", "cordova-plugin-group-project\\www\\util.js",
		"cordova-plugin-group-project\\www\\project.js.map", "cordova-plugin-group-project\\www\\util.js.map",
		"cordova-plugin-group-project\\package.json", "cordova-plugin-group-project\\plugin.xml"}

	// fmt.Println(linuxFileList)
	// fmt.Println(winFileList)

	files, _ := reader.FilePathWalkDir("cordova-plugin-group-project")

	checkFiles := func(given []string, want []string, t *testing.T) {
		fmap := make(map[string]string)
		for _, file := range given {
			fmap[file] = file
		}

		fmt.Println(fmap)
		for _, file := range want {
			_, ok := fmap[file]
			if !ok {
				t.Errorf("want: %s, given: %s", fmap[file], file)
			}
		}
	}

	if len(files) != len(winFileList) {
		t.Error()
	}

	if runtime.GOOS == "windows" {
		checkFiles(files, winFileList, t)
	} else {
		checkFiles(files, linuxFileList, t)
	}

	os.RemoveAll("cordova-plugin-group-project")
}