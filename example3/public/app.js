new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        username: null, // Our username
        joined: false // True if username has been filled in
    },
     created: function() {
        var self = this;
        // create websocket connection
        this.ws = new WebSocket('ws://' + window.location.host + '/chat');
        // listen for messages coming in - JS onmessage is MessageEvent named "message"
        this.ws.addEventListener('message', function(e) {
            // parse the message
            var msg = JSON.parse(e.data);
            // what is "chip" - (contact chip) https://www.w3schools.com/howto/howto_css_contact_chips.asp?
            self.chatContent += '<div class="chip">'
                + msg.username
                + '</div>'
                + emojione.toImage(msg.message) + '<br/>'; // Parse emojis

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },
    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },
        join: function () {
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.username = $('<p>').html(this.username).text();
            this.joined = true;

            this.chatContent += '<div class="chip">'
                + '<b>' + this.username + '</b>'
                + ' joined'
                + '</div>'
                + '</br>';

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        }
    }
});
