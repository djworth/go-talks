type Pattern interface {
    // In practice, most real-world routes have a string prefix that can be
    // used to quickly determine if a pattern is an eligible match. The
    // router uses the result of this function to optimize away calls to the
    // full Match function, which is likely much more expensive to compute.
    // If your Pattern does not support prefixes, this function should
    // return the empty string.
    Prefix() string
    // Returns true if the request satisfies the pattern. This function is
    // free to examine both the request and the context to make this
    // decision. Match should not modify either argument, and since it will
    // potentially be called several times over the course of matching a
    // request, it should be reasonably efficient.
    Match(r *http.Request, c *C) bool
    // Run the pattern on the request and context, modifying the context as
    // necessary to bind URL parameters or other parsed state.
    Run(r *http.Request, c *C)
}