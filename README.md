**Notes:**
- Should follow https://github.com/golang-standards/project-layout
- This is a learning experience too, so some code, more specifically docker pkg is for fun and should not be used seriously. Only wanted to write parts of it for learning more about the docker API

**General structure**
- commons/ This is where common shared things that can be used for all services
- docker/ A docker client for getting the information about services and containers
- faas-gateway/ The L7 reverse proxy
- multiply/ The dummy service

**Overall architecture**

Task was to create a L7 reverse proxy and a dummy service. There is not much to say, but I follow the recommendations given. We let  each faas gateway container poll the docker API every X second. 

If we let the number of manager nodes be few and we have a lot of worker nodes, scaling the gateway will be limited since it is dependent on having access to the docker API, which only a manager node has acccess too (if using the unix socket approach). 

Requirements on load balancing is automatically taken care of by using the overlay network in docker swarm that automatically balances between nodes in a service. 

**TODOS:**
- Investigate net.dial and if we need to close unix socket after in the docker pkg
- Copy paste code in docker/pkg default_impl, check how to make it better
- Fix the dockerfiles so that common libraries are fetched from local and not from github by govendor

**Testing the code**

docker-compose build
docker stack deploy --compose-file docker-compose.yml faas
docker service scale X=10 // Scale it to 10 instances

Visit http://127.0.0.1:8081/multiply/api/v1/add?x=10&y=25 for example

**Testing the load balancing**

Paste this in the browser for example the inspect tool's console

    var buckets = {}; 
    var max_calls = 5000; 
    var calls = 0; 
    var intervaller = setInterval(() => { 
    calls++; 
    if (calls % (max_calls / 10) === 0) {
        console.log(calls + "/" + max_calls)
    }
    if (calls > max_calls) {
        console.log(buckets);clearInterval(intervaller)
    } 
    var data = fetch("http://127.0.0.1:8081/multiply/api/v1/add?x=10&y=25")
              .then((resp) => { return resp.json();})
              .then((json) => {
                var serviceID = json.Data.ServiceID; 
                if (buckets[serviceID] == null) {
                  buckets[serviceID] = 0
                }
                buckets[serviceID] = buckets[serviceID] + 1;
              });
    }, 0);
    
It will print out what container the requests went too. Depending on how many parallel requests can be made and the internals of docker swarm routing the results can vary and even be not fair.