func hello(c web.C, w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {
        goji.Get("/hello/:name", hello)
        goji.Serve()
}