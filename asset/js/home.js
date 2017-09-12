/**
 * Created by shang on 2017/8/2.
 */
$(document).ready(function () {
    http.getAjax_clean("/v1/listAll", function (data) {
        $('#table').bootstrapTable({
            columns: [{
                field: 'Src',
                title: '端口',
            }, {
                field: 'Dst',
                title: '映射',
                formatter: function (value, row, index) {
                    return "<span style='color: #0005ff;background: #ffe96a'>" + row.Src + " --> " + row.Dst + "</span>"
                }
            }, {
                field: 'view',
                title: 'view',
                formatter: viewFormatter
            }, {
                field: 'UseBytes',
                title: '使用流量'
            }, {
                field: 'ops',
                title: '操作',
                formatter: operateFormatter
            }],
            data: data
        });
    });
    initWebSocket();
})

function viewFormatter(value, row, index) {
    return [
        '<botton type="button" class="btn btn-sm btn-default cs-btn" onclick="change(\' + index + \')"><span class="glyphicon glyphicon-stats" aria-hidden="true"/> info</botton>',
        '<botton type="button" class="btn btn-sm btn-default cs-btn" onclick="del(\'' + row.title + '\')"><span class="glyphicon glyphicon-tint" aria-hidden="true"/> limit</botton>'
    ].join('');
}

function operateFormatter(value, row, index) {
    return [
        '<button type="button" class="btn btn-sm btn-default cs-btn" onclick="change(' + index + ')"><span class="glyphicon glyphicon-play" aria-hidden="true"/> start</button>',
        '<button type="button" class="btn btn-sm btn-default cs-btn" onclick="del(\'' + row.title + '\')"><span class="glyphicon glyphicon-pause" aria-hidden="true"/> pause</button>',
        '<button type="button" class="btn btn-sm btn-default cs-btn" onclick="del(\'' + row.title + '\')"><span class="glyphicon glyphicon-trash" aria-hidden="true"/> delete</button>'//
    ].join('');
}


function add() {
    $(".add_dia").modal("show");
}

function saveAdd() {

    var url = $("#title").val();
    var local = $("#local").val();
    var path = $("#path").val();
    var strp = $("#stripPrefix").is(':checked')
    var from = new FormData();
    from.append("title", url);
    from.append("url", local);
    from.append("path", path);
    from.append("stripPrefix", strp);
    http.postAjax_clean("route/add", from, function (data) {
        if (data.state == true) {
            window.location.reload();
        }

    })
}

function del(i) {
    BootstrapDialog.confirm('确认要删除' + i + "这条路由？", function (result) {
        if (result) {
            var data = new FormData();
            data.append("title", i)
            http.postAjax_clean("route/delete", data, function (resdate) {
                if (resdate.state == true) {
                    window.location.reload();
                }
            })
        }
    });


}
function startPort() {
}
function change(i) {
    http.getAjax_clean("route/" + (i), function (data) {
        changDialog(data)
    })

}

function changSave(title) {
    var local = $("#local");
    var stripPrefix = $("#stripPrefix").is(':checked')
    var fromDate = new FormData();
    fromDate.append("title", title);
    fromDate.append("local", local.val());
    fromDate.append("stripPrefix", stripPrefix);
    http.postAjax_clean("route/change", fromDate, function (resdate) {
        if (resdate.state == true) {
            window.location.reload();
        }
    })

}

function changDialog(data) {
    var title = $("#title");
    var local = $("#local");
    var path = $("#path");
    var savebtn = $(".zuul-chang");
    var diaTitle = $("#myModalLabel");
    title.val(data.id);
    title.attr("disabled", true);
    local.val(data.location);
    path.attr("disabled", true);
    savebtn.attr("onclick", "changSave('" + data.id + "')");
    path.val(data.fullPath);
    diaTitle.text("改变路由");
    $(".add_dia").modal("show");
}


var websocket = null;
var localurl = document.location.href.split("/")[2] + "/routes/speed";

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
    websocket.onerror = onError;
    websocket.onclose = onClose;
}

function onOpen(evt) {

}

function onClose(evt) {

}

function onMessage(evt) {
    $(".speed").text(evt.data + " t/s");
}

function onError(evt) {

}