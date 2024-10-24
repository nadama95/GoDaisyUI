{
  description = "Go Daisy UI Template Library";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = {self, nixpkgs } :
  let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in  {
    devShells.${system}.default = pkgs.mkShell {
      packages = with pkgs; [
        delve
        go
        gopls
        gotests
        go-tools
      ];
    };
  };
}