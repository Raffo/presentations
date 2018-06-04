srv := &http.Server{
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
    TLSConfig:    tlsConfig,
    Handler:      serveMux,
}
log.Println(srv.ListenAndServeTLS("", ""))