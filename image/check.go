package image


import (
	"net/http"
	"strings"
	"fmt"
)


func CheckImage(r *http.Request) (bool, bool) {
	r.ParseMultipartForm( 10 << 30)

	check :=  r.MultipartForm

	if check == nil  {
		return false, false
	}else{

		file := r.MultipartForm

		if len(file.File["files"]) > 0 {
			return true, true
		}else{
			return true, false
		}
	}

}



func CheckSingleImage (r *http.Request) (bool) {
	r.ParseMultipartForm(10 << 20)

	_, handle, err := r.FormFile("file")

	var vidcheck bool

	vids := []string{".mp4", ".ogg", ".webm", ".mov", ".rm", ".ram", ".mpeg"}

	if err == nil {

		for _, k := range vids {

			if strings.HasSuffix(handle.Filename, k) == true {
				vidcheck = true
			}
		}

		if vidcheck == true {
			fmt.Println("video not supported")
			return false
		}else{
			return true
		}
	}else{
		return false
	}


}
