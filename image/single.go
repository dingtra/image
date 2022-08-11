package image


import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)


func UploadSingleImage(r *http.Request) string {

	collectimages := []string{}

	r.ParseMultipartForm( 10 << 20 )

	file, handler, err := r.FormFile("file")

    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
    }

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    defer file.Close()

	tempfile, err := ioutil.TempFile("./jslmnjsyofubfgd/images/profile", "dingtra-*.jpg")

    if err != nil {
        fmt.Println(err)
    }

	ma := strings.Split(tempfile.Name(), "/")


	collectimages = append(collectimages, ma[len(ma)-1])

    defer tempfile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempfile.Write(fileBytes)

	return strings.Join(collectimages, "")
}
