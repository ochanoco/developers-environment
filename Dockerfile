FROM nixos/nix

# set up argument
ARG HOST_FOLDER="."
ARG NIX_FILE="./default.nix"
ARG BUILD_CMD="go build ."
ARG COMMAND="./serv"

# update
RUN nix-channel --update

# build nix
RUN mkdir /workspace
COPY ${NIX_FILE} /workspace/default.nix 
WORKDIR /workspace
RUN nix-build default.nix

WORKDIR /workspace/serv

# copy project files
# build serv
COPY ${HOST_FOLDER} /workspace/serv
RUN nix-shell ../default.nix --command "${BUILD_CMD}"
CMD nix-shell ../default.nix --command "${COMMAND}"