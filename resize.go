package main

import (
    "fmt"
    "github.com/nfnt/resize"
    "image/jpeg"
    "image/png"
    "image/gif"
    "image"
    "os"
    "time"
    "path/filepath"
    "strings"
)

var rawDir string = "/Users/nci/Downloads/media/"
var resizedDir string = "/Users/nci/Desktop/media/"
var resizedCount = 0
var resizedErrors = 0

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func resizeImage(path string, f os.FileInfo, err error) error {
    var outPath string = strings.Replace(path, rawDir, resizedDir, -1)
    if f.IsDir() || Exists(outPath) {
        // fmt.Println("Skip: ", path)
        return nil
    }
    var extension = strings.ToLower(filepath.Ext(path))
    if extension != ".jpeg" && extension != ".jpg" && extension != ".png" && extension != ".gif" {
        fmt.Println("Extension not support: ", extension, ", path: ", path)
        resizedErrors++
        return nil
    }
    // fmt.Println("Resizing: ", f.Name())
    if resizedCount % 100 == 0 {
        fmt.Println("Resizing on process! Resized: ", resizedCount, ". Errors: ", resizedErrors)
    }

    fin, err := os.Open(path)
    if err != nil {
        fin.Close()
        resizedErrors++
        fmt.Println("Open error: ", path)
        return nil
    }
    defer fin.Close()

    img, format, err := image.Decode(fin)
    if err != nil {
        resizedErrors++
        fmt.Println("Decode error: ", path)
        return nil
    }
    fin.Close()

    var maxWidth uint = 99999
    var maxHeight uint = 1024
    origBounds := img.Bounds()
    origWidth := uint(origBounds.Dx())
    origHeight := uint(origBounds.Dy())
    newWidth, newHeight := origWidth, origHeight

    // Preserve aspect ratio
    if origWidth > maxWidth {
        newHeight = uint(origHeight * maxWidth / origWidth)
        if newHeight < 1 {
            newHeight = 1
        }
        newWidth = maxWidth
    }

    if newHeight > maxHeight {
        newWidth = uint(newWidth * maxHeight / newHeight)
        if newWidth < 1 {
            newWidth = 1
        }
        newHeight = maxHeight
    }

    m := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

    outDir, err := filepath.Abs(filepath.Dir(outPath))
    if err != nil {
        resizedErrors++
        fmt.Println("Creation error: ", outDir)
        return nil
    }
    os.MkdirAll(outDir, os.ModePerm)

    fout, err := os.Create(outPath)
    if err != nil {
        fout.Close()
        resizedErrors++
        fmt.Println("Creation error: ", outPath)
        return nil
    }
    defer fout.Close()

    // write new image to file
    switch format {
        case "jpg":
            jpeg.Encode(fout, m, &jpeg.Options{50})
        case "jpeg":
            jpeg.Encode(fout, m, &jpeg.Options{50})
        case "png":
            png.Encode(fout, m)
        case "gif":
            gif.Encode(fout, m, nil)
        default:
            fout.Close()
            resizedErrors++
            fmt.Println("Creation error: ", outPath)
            return nil
    }
    fout.Close()
    resizedCount++
    return nil
}


func main() {

    // FOR LOGGING
    var startTs int64 = time.Now().UnixNano() / int64(time.Millisecond)
    fmt.Println("Image resize start at: ", startTs)

    err := filepath.Walk(rawDir, resizeImage)
    fmt.Printf("END %v\n", err)

    var endTs int64 = time.Now().UnixNano() / int64(time.Millisecond)
    fmt.Println("End at: ", endTs, ". Total time: ", endTs - startTs)
    fmt.Println("Resized files: ", resizedCount, "|| Errors: ", resizedErrors)
}
