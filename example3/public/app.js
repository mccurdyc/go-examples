var ws = new WebSocket('ws://' + window.location.host + '/chat');
var username = String(Math.floor(Math.random() * 1000000000000));

var messageHistory = document.getElementById('history');
var send = document.getElementById("send");
var message = document.getElementById("message");

document.getElementById("username").innerHTML = username;

ws.addEventListener('open', function(e) {
  console.log("connected");
});

ws.addEventListener('error', function(e) {
  console.log("error");
});

ws.addEventListener('message', function(e) {
  var msg = JSON.parse(e.data);

  messageHistory.innerHTML += "<p>" + msg.username + " says: " + msg.message + "</p>";
});

send.addEventListener("click", function(e) {
  ws.send(
    JSON.stringify({
      username: username,
      message: message.value
    })
  );
});

