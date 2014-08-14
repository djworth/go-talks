func updateComment(c web.C, w http.ResponseWriter, r *http.Request) {
    // Don't do it like this but you get the point
    comment, _ := db.GetComment(strconv.Itoa(c.URLParams["id"]))
    db.UpdateComment(comment, r)
}

func main() {
    re := regexp.MustCompile("^/comment/(?P<id>\d+)$")
    goji.Put(re, updateComment)
    goji.Serve()
}
