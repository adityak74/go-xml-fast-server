package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    http.HandleFunc("/file-response", fileResponseHandler)
    http.HandleFunc("/optimized-file-response", optimizedFileResponseHandler)

    fmt.Println("Server is running on port 8000...")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Println("Failed to start server:", err)
        os.Exit(1)
    }
}

func fileResponseHandler(w http.ResponseWriter, r *http.Request) {
    filePath := filepath.Join("data", "example.xml")
    http.ServeFile(w, r, filePath)
}

func optimizedFileResponseHandler(w http.ResponseWriter, r *http.Request) {
    filePath := filepath.Join("data", "example.xml")
    file, err := os.Open(filePath)
    if err != nil {
        http.Error(w, "File not found.", http.StatusNotFound)
        return
    }
    defer file.Close()

    fileStat, err := file.Stat()
    if err != nil {
        http.Error(w, "Could not get file info.", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Length", fmt.Sprintf("%d", fileStat.Size()))
    w.Header().Set("Content-Type", "application/xml")
    w.Header().Set("Content-Disposition", `attachment; filename="example.xml"`)
    w.Header().Set("Cache-Control", "no-cache")

    http.ServeContent(w, r, file.Name(), fileStat.ModTime(), file)
}
