{
  description = "NetmanGo";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

    templ.url = "github:a-h/templ/e54517eb7d8a7d9ef67d43177709a0ca26235e43";
    templ.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = inputs@{ ... }:
    let
      system = "x86_64-linux";
      pkgs = inputs.nixpkgs.legacyPackages.${system};
      templ = system: inputs.templ.packages.${system}.templ;
    in
    {

      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          delve
          go
          gopls
          gotests
          go-tools
          (templ system)

          nodejs
          yarn-berry
        ];
      };
    };
}
