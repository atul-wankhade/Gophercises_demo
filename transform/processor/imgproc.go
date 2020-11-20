package processor

import (
	"fmt"
	"os/exec"
	"strings"
)

//Mode is mode of image transformation
type Mode int

const (
	//ModeCombo is a primitive mode
	ModeCombo Mode = iota

	//ModeTriangle is a primitive mode
	ModeTriangle

	//ModeRectangle is a primitive mode
	ModeRectangle

	//ModeEllipse is a primitive mode
	ModeEllipse

	//ModeCircle is a primitive mode
	ModeCircle

	//ModeRotatedrect is a primitive mode
	ModeRotatedrect

	//ModeBeziers is a primitive mode
	ModeBeziers

	//ModeRotatedellipse is a primitive mode
	ModeRotatedellipse

	//ModePolygon is a primitive mode
	ModePolygon
)

//WithMode returs the function, the underlying function return the diffrent mode options
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

//Transform accepts an image and transform it using defined mode and shapes
// func Transform(image io.Reader, ext string, numShapes int, mode Mode) (io.Reader, error) {
// 	//Create the file to store input io reader
// 	in, err := MyTempFile("in_", ext, "")
// 	if err == nil {
// 		//Create the file to store transformed image
// 		defer os.Remove(in.Name())
// 		//Creating temp file to store out image
// 		out, err := MyTempFile("out_", ext, "")
// 		defer os.Remove(out.Name())
// 		if err == nil {
// 			//Read image into file
// 			_, err = io.Copy(in, image)
// 			if err != nil {
// 				return nil, errors.New("primitive: Error while reading image into file")
// 			}

// 			//Transform image using primitive api
// 			stdOutput, err := Primitive(in.Name(), out.Name(), numShapes, mode)
// 			if err != nil {
// 				return nil, errors.New("primitive: Error in transform image")
// 			}
// 			_ = stdOutput
// 			b := bytes.NewBuffer(nil)
// 			//Coping the content of image output buffer
// 			_, err = io.Copy(b, out)
// 			if err == nil {
// 				return b, nil
// 			}
// 		}
// 	}
// 	return nil, errors.New("primitive: Error in creating input temp file")

// }

//Primitive takes the input, output file , number of shapes, and mode and returns the output and error
func Primitive(inputFile, outputFile string, numShapes int, mode Mode) (out string, err error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(argStr)...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

//MyTempFile takes prefix, extension and directory string and returns a newly created temporary file
// func MyTempFile(prefix, ext, dir string) (*os.File, error) {
// 	in, err := ioutil.TempFile(dir, prefix)
// 	if err != nil {
// 		return nil, errors.New("error while creating temparary file")
// 	}
// 	defer os.Remove(in.Name())
// 	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
// }
