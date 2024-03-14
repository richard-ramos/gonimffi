proc NimMain() {.importc.}

proc hello() {.dynlib, exportc, cdecl.} =
    NimMain()
    echo "hello world"