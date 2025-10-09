const ws = new WebSocket("ws://localhost:8081/ws");

ws.addEventListener("message", (event) => {
  const message = JSON.parse(event.data);
  if (message == null || message == undefined) {
    return;
  }

  const target = document.getElementById(message.Id);
  if (target == undefined) {
    return;
  }

  target.outerHTML = message.Value;
});

ws.addEventListener("open", (ev) => {
  const sessionId = document.body.getAttribute("lens-session-id");
  ws.send(sessionId);
});

const buttonClicked = () => {};

const valueChanged = (element) => {
  const id = element.getAttribute("lens-target-id") ?? element.id;
  const value = element.value;
  ws.send(JSON.stringify({ Id: id, Value: value }));
};
