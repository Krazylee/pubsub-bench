var redis = require("redis"),
	pub = redis.createClient(),
	client = redis.createClient();

var uid = process.argv[2]
if(uid == "undefined") {
	console.log("usage: node simple.js uid{1,2,3...}")
	uid = 0;
}

console.log("\n\nuser " + uid + " Published to his channel:");
pub.publish("channel" + uid, "I am node.js");

client.on("message", function (channel, message) {
	console.log("user " + uid + " received from: " + channel + " message: " + message);
});

for(var i =0 ; i <= 300; i++){
	if( i != uid ){
		client.subscribe("channel" + i);
	}
}
