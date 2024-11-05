1) I found the requirements a bit vague concerning the expected results.
  I used best paractices for returning http errors for bad request type and bad request data
  In other words I only returned CommandResponse if I could properly parse the request.

2) I didn't see a restriction on using packages so I'm using them

3) I got confused about CommandResponse.Success and PingResult.Successful
  TLDR; I added a SystemInfoResult.Successful. I return errors evaluating the request as CommandResponse.Success = false
  Basically 
  If the command was correctly intepreted but the ping failed, I return
     CommandResponse.Success = true and PingResult.Successful = false
  Since the SystemInfoResult doesn't have a Successful I have to put CommandResponse.Success = false 

4) for SystemInfoResult.IPAddress I returned ALL ip addresses for my host

5) My installer , a more elegant solution would be to install as a service

6) in the interest of time, 
  I didn't document how I created the install package, but I did mention some details about the end result in the README.md