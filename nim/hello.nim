proc NimMain() {.importc.}

proc hello() {.dynlib, exportc, cdecl.} =
    NimMain()
    echo "==================="
    echo "==================="
    echo "==================="
    echo "   hello world"
    echo "==================="
    echo "==================="
    echo "==================="