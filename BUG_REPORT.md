#


## Links to issue
*[Github](https://github.com/WindMillCode/overide-hashicorp-root-token-invalid-value)

## Addtl links
Current Session

   https://app.slack.com/client/T029RQSE6/C2B4L99RS/thread/C029RQSEE-1686429799.955169
   https://www.facebook.com/search/groups/?q=golang&sde=AbotKqEx6VFO7Xg9wiFDgE3hDSnw1teizHLB49vZpXF3bLbxyZ0OOuvrfrGEGCpzofDakrn8opGG9ML5UM4eRoBs
   https://www.facebook.com/groups/265634821388311/
   https://www.facebook.com/groups/1693451300902571/
   https://www.facebook.com/groups/1622817837978047/
   https://www.facebook.com/groups/golanggonuts/
   https://www.facebook.com/groups/learnGolang/
   https://www.facebook.com/groups/422428585164623/
   https://www.facebook.com/groups/1908750736042854/
   https://www.facebook.com/groups/mygolang/
   https://www.facebook.com/groups/586803232936791/
   https://www.facebook.com/groups/139882946682393/
   https://www.reddit.com/r/golang/submit
   https://stackoverflow.com/questions/76466880/dlv-debug-issue-to-go-application-running-in-docker-container
   https://stackoverflow.com/questions/74740177/dlv-debug-error-could-not-launch-process-fork-exec-demo-function-not-impleme

   chrome://newtab/
### Problem
Greetings,

I am currently facing an issue related to setting the root token ID to a previous value in the HashiCorp Vault Go code. I have tried debugging the application using dlv debug within a Docker container. However, when attempting to run the application on my Windows machine, it hangs. Although the application eventually runs on the Docker machine, I encountered an issue with the debugger not functioning properly when setting breakpoints.

To provide a better understanding of the problem, I have attached a minimal reproduction of the issue on GitHub. You can find it at the following link: GitHub - WindMillCode/overide-hashicorp-root-token-invalid-value.

I kindly request the to help me resolve this problem as quickly as possible

Thank you in advance.

## Current Behavior
When I run the application via start debugging and I set breakpoints the debugging skips over the breakpoints not allowing me to explore the code. My use case I am working with the [hashicorp vault code](https://github.com/hashicorp/vault) trying to put break points ultimately to modify the code to accept any dev token I give it (right now its rejecting using its own tokens). I am using this course for reference https://app.pluralsight.com/library/courses/go-debugging/table-of-contents.

* This is the breakpoint
[FILE](minimal_repro\main.go)
![1686669469254](image/BUG_REPORT/1686669469254.png)

* We can see after I start debugging the breakpoint is skipped
![1686669583253](image/BUG_REPORT/1686669583253.png)
![1686669596457](image/BUG_REPORT/1686669596457.png)

* and I can access the application as its has gotten to http.ListenAndServe
![1686669631486](image/BUG_REPORT/1686669631486.png)

* according to this [link](https://github.com/go-delve/delve/blob/master/Documentation/faq.md#substpath) I can the content of main.go
![1686680350276](image/BUG_REPORT/1686680350276.png)

## Expected Behavior
stop at the breakpoint and I should get connection refesued if I try to open the application in the browser


## Possible Solution
* does the go in the docker and the go in my local have to be the same verison I feel like it defeats the point
* I am getting in go.mod
![1686669695784](image/BUG_REPORT/1686669695784.png)


## Steps to Reproduce

### Pre reqs
install docker and  go1.20 on the computer and

1. Go to project root in terminal
2. Copy paste and run commands in terminal minimal_repro\commands.txt
3. set a breakpoint on  line 15 in minimal_repro\main.go
4. Click [Top left of vscode > Run > Start debugging]
5. access localhost:5001 in the browser if you access it this means the debugger did not work

## Environment
* Local machine
go version go1.20.5 windows/amd64
OS Name:                   Microsoft Windows 11 Pro
OS Version:                10.0.22621 N/A Build 22621
System Manufacturer:       LENOVO
System Model:              82SX
System Type:               x64-based PC

* Docker container
go version go1.20.5 windows/amd64
Linux 23fb080a459e 5.15.90.1-microsoft-standard-WSL2 #1 SMP Fri Jan 27 02:56:13 UTC 2023 x86_64 GNU/Linux

## Files


## Detailed Description

## Possible Implementation


## Additional Issues
