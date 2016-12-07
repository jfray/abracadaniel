# o2m8 (automate)
## Mashup of Wunderlist + Magic == Rainbows

### **_Have My People Call My People_**
![](https://camo.githubusercontent.com/7c2d252b7fd984ec30d7446eccb604a297529cef/687474703a2f2f692e696d6775722e636f6d2f347452346145762e706e67)

#### General Idea
* User adds Abracadaniel to their Wunderlist user list
* User assigns a task to Abracadaniel, which sends an email to a specially formatted email address that encodes the info that it needs.
* Abracadaniel then sends a text to Magic, using an internal-use SMS code so that it can capture the response.
* The response is posted to a Slack channel or some other chat service so that the user can interact with the Magic representative in real time.
* The user is sent a notification pointing them to the slack/chat channel
* The channel has bots that handle things like checking calendar automatically, verifying funds for purchase, etc.
* Instead of using a stateful chat, maybe just use a group SMS?

#### Possible State Diagram ####
![](https://cloud.githubusercontent.com/assets/192726/9698413/e2e1a808-5369-11e5-9c62-4e9c8985a6d0.png)
