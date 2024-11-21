# WebGL with Go

Experiment with Go (TinyGo) and WebGL.

## TinyGo
Most of my work is done using the official [Go](https://go.dev/) compiler, but
for WebAssembly I've been experimenting with [TinyGo](https://tinygo.org/) because
it gives me much more compact binaries to deploy. Check it out, it's an interesting
project!

## Zed
I'm exploring using [Zed](https://zed.dev/) as my primary editor and there's a `zed.sh`
file that starts Zed with the correct WASM environment variables so that the editor
has context about the code.  You are of course not forced to use Zed, the project
does not depend on it in any way.

## Task
To save myself endless typing, [Task](https://taskfile.dev/) is used to do things
like manage the build process etc.  There's a `Taskfile.yml` file that contains
all the build, run, clean and lint instructions for the project.

## Technologies used
- Go <https://go.dev/>
- TinyGo <https://tinygo.org/>
- Task <https://taskfile.dev/>
