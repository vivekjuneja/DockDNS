A sample Fig Application inspired from the sample at http://www.fig.sh/

**Before Using**

1. Change the "dns" entry in the fig.yml to the appropriate server endpoint of DockDNS

2. Note the reference to the "Redis" service in the app.py as : "redisserver.docker". "redisserver" is the name of the Service declared in the fig.yml. ".docker" is the standard suffix used for all containers. 

3. ".docker" refers to the appropriate DNS zone as maintained by the DockDNS

