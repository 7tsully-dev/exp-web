#shell.nix
{ pkgs ? import <nixpkgs> {}}:
with pkgs; mkShell {
 packages = [
    figlet
    lolcat
    go
    gotools
    gopls
    go-outline
    gopkgs
    gocode-gomod
    bash
 ];
 shellHook = ''
    export GOPATH="/home/$USER/go"
    export GOMODCACHE="$GOPATH/pkg/mod"
    figlet Welcome to Steven\'s go dev env | lolcat
    go mod tidy
 '';
}
