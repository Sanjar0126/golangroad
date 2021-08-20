package main

import (
	"archive/zip"
	"fmt"
	"github.com/jonas-p/go-shp"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

func main() {
	//read, err := zip.OpenReader("Toshkent_vino.zip")
	//if err != nil {
	//	msg := "Failed to open: %s"
	//	log.Fatalf(msg, err)
	//}
	//
	//defer read.Close()
	// for _, file := range read.File {
	// if err := listFiles(file); err != nil {
	// log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
	// }
	// }
	//shapeFileRead()

	g := geojson.NewPointGeometry([]float64{1, 2})
	rawJSON, err := g.MarshalJSON()

	fc := geojson.NewFeatureCollection()
	fc.AddFeature(geojson.NewPointFeature([]float64{1, 2}))
	rawJSON, err := fc.MarshalJSON()
}

func shapeFileRead() {
	// open a shapefile for reading
	shape, err := shp.OpenZip("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()
		// print feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

		// print attributes
		for _, f := range fields {
			fmt.Println(f)
			val := shape.Attribute(n)
			fmt.Printf("\t%v : %v\n", f, val)
		}
		fmt.Println()
	}
}

func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()
	var extension = filepath.Ext(file.Name)
	fmt.Println(extension)
	fmt.Fprintf(os.Stdout, "%s:", file.Name)
	if extension == ".shp" {
		// shapeFileRead(file)
		fmt.Println("------")
	}
	if err != nil {
		msg := "Failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	return nil
}
