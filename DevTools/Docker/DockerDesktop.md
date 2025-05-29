# Docker Desktop :
* Docker only runs on linux so they can't run the containers in the Window and MacOs 
* So Docker uses Docker machine it using the virtualbox to create small linux based VM's that only run Docker it was slower than in linux and people using it must have knowledge of virtualbox
* In 2016 Docker creates Dockerdesktop to works with Mac and windows

## Dockerdesktop Improvements :
* use much small and tightly integrated VM's 
* automatically handle volume and network mapping
* really nice GUI
* Vm's integrated with apple native hypervisor virtualkit and hypervisor on windows

## Container :
* container virtualize operating system kernals
* Don't emulate the hardware so no need to boot up
* Take up much less space
* can run only one app at a time usually
* can interact with the host

## Virtual machine :
* Virtual machine virtualize the hardware
* we need to install software and files manually because VM acts as the real hardware
so it can take up a lot of space
* VM use hypervisor to emulate the real hardware
* can run multiple apps
* cannot interact with the hosts