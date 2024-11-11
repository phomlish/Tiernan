
# Your architectural decisions

I kept it simple. Normally I would create packages to keep things organized. IE: http, commands

# How you would scale this for a large organization

As application doesn't write or read anything to a database there would be no thread issues so a simple deployment of multiple instances and a load balancer would be appropriate. Did I mentionthat I'm a big fan of kubernetes.

# Security considerations

Implement TLS endpoint and disbale non TLS.
Require authentication for the API calls.
Implement anti-DDoS hardware/software

# Performance optimizations

Spawn the commands into threads. Monitor the threads and return the results when they all finish.

# How you'd handle offline functionality

I'm not sure how this applies to the application, each endpoint already has an error condition that would indiacate 'offline'. IE: Ping test failes when network is down.
For the backend I would save the execution results for each run into a history (database, event queue, etc).
