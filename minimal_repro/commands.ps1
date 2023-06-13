cd .\minimal_repro\
docker stop hackathon-square-godebug-issue-0-container
docker rm hackathon-square-godebug-issue-0-container
docker build --tag windmillcode/hackathon-square-godebug-issue-0:latest .
docker run -it --name hackathon-square-godebug-issue-0-container  --security-opt="seccomp=unconfined" --cap-add=SYS_PTRACE -p:5001:5000 -p:2345:2345 windmillcode/hackathon-square-godebug-issue-0:latest 
