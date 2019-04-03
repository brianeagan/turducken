Turducken

_Decadently form over function._ 

A go "web server" that is compiled into webassmebly, 
loaded directly into the browser, and serves content in templates.

Optionally it can be loaded into a mobile app to achieve complete lolwut status.

 Step 1 - Make the environment
 
 + Make sure you have/use Go > 1.11
 + Copy the proper version of misc/wasm/wasm_exec.js from your GOROOT 
 + Please note, this was developed with Go 1.12.1 and not tested with anything else.
  Go compilation to webassembly is experimental
 + Tested in Chrome, use Chrome
 
 
 Step 2 - build the wasm
 + make build
 + (OSX) open turducken.html
 
 
 Notes
 
 + My hello world wasm file was 1.3mb
 + Libraries don't always understand being compiled into webasm
    + packr v2 uses a director walk library that didn't compile
 + Adding packr with a tiny director => 8.6 mb
 + 
 
 
*Bonus Notes*

Using go modules

1) go mod init
2) go mod vendor
3) go mod tidy <-- when you want a different version/etc 

done

 
 
 