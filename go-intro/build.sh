$ GOOS=linux go build hello.go
$ file hello
$ hello: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, with debug_info, not stripped

$ GOOS=darwin go build hello.go
$ file hello
$ hello: Mach-O 64-bit executable x86_64
