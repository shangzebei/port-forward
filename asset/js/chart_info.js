var websocket = null;
var localurl = document.location.href.split("/")[2] + "/v1/info";
initWebSocket()

function initWebSocket() {

    if ('WebSocket' in window) {
        websocket = new WebSocket("ws://" + localurl);
    }
    else if ('MozWebSocket' in window) {
        websocket = new MozWebSocket("ws://" + localurl);
    }
    else {
        websocket = new SockJS("ws://" + localurl);
    }
    websocket.onopen = onOpen;
    websocket.onmessage = onMessage;
    // websocket.onerror = onError;
    websocket.onclose = onClose;
}

function onOpen(evt) {
    websocket.send("6677")
}

function onClose(evt) {
    debugger
}

function onMessage(evt) {
    var sum = JSON.parse(evt.data)
    cpuData.push({value: [new Date(), sum.speed]})
    myChart.setOption({
        series: [{
            data: cpuData,
        }]
    })

}

function formatBytes(value) {
    var bytes = parseFloat(value);
    if (bytes > 1048576 * 1024) {
        return (bytes / (1048576 * 1024.0)).toFixed(2) + "GB"
    }
    if (bytes > 1048576) {
        return (bytes / 1048576.0).toFixed(2) + "MB"
    }
    if (bytes > 1024) {
        return (bytes / 1024.0).toFixed(2) + "KB"
    }
    return bytes
}