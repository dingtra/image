package image


import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
)



func UploadImage(r *http.Request) {

	r.ParseMultipartForm ( 20 << 30 )

	raw_file := r.MultipartForm

	get := raw_file.File["files"]

	fmt.Println(len(get))

	for _, files := range get {
		vr, _ := files.Open()
		defer vr.Close()
		

		tempfile, err := ioutil.TempFile("../jslmnjsyofubfgd/images/post", "dingtra-*.jpg")

		if err != nil{
			fmt.Println(err)		
		}

		defer tempfile.Close()

		fmt.Println(strings.Split(tempfile.Name(), "/"))

		filebytes, err := ioutil.ReadAll(vr)

		if err != nil{
			fmt.Println(err)
		}
				
		tempfile.Write(filebytes)
			
				


	}
}