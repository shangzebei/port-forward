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
                title: '限制速度 | 使用流量',
                formatter: function (value, row, index) {
                    if (row.LimitSpeed.indexOf('0') === 0) {
                        return "-- | " + row.UseBytes
                    }
                    return row.LimitSpeed + " | " + row.UseBytes;
                }
            }, {
                field: 'ops',
                title: '操作',
                formatter: operateFormatter
            }],
            data: data
        });
    });
    // initWebSocket();
})

function viewFormatter(value, row, index) {
    return [
        '<botton type="button" class="btn btn-sm btn-default cs-btn" onclick="info(\'' + row.Src + '\')"><span class="glyphicon glyphicon-stats" aria-hidden="true"/> info</botton>',
        '<botton type="button" class="btn btn-sm btn-default cs-btn" onclick="limitSpeed(\'' + row.Src + '\')"><span class="glyphicon glyphicon-tint" aria-hidden="true"/> limit</botton>'
    ].join('');
}

function operateFormatter(value, row, index) {
    return [
        // '<button type="button" class="btn btn-sm btn-default cs-btn" onclick="change(' + index + ')"><span class="glyphicon glyphicon-play" aria-hidden="true"/> start</button>',
        // '<button type="button" class="btn btn-sm btn-default cs-btn" onclick="del(\'' + row.title + '\')"><span class="glyphicon glyphicon-pause" aria-hidden="true"/> pause</button>',
        '<button type="button" class="btn btn-sm btn-default cs-btn" onclick="del(\'' + row.Src + '\')"><span class="glyphicon glyphicon-trash" aria-hidden="true"/> delete</button>'//
    ].join('');
}


function add() {
    $(".add_dia").modal("show");
}

function del(i) {
    BootstrapDialog.confirm('确认要删除端口 ' + i + " 这条映射？", function (result) {
        if (result) {
            var data = new FormData();
            data.append("port", i)
            http.postAjax_clean("/v1/stopPort", data, function (resdate) {
                if (resdate.state === "ok") {
                    window.location.reload();
                }
            })
        }
    });


}

function startPort() {
    var localPort = $("#localPort").val();
    var forward = $("#forward").val();
    var from = new FormData();
    from.append("src", localPort);
    from.append("dst", forward);
    http.postAjax_clean("/v1/startPort", from, function (data) {
        if (data.state === 'ok') {
            window.location.reload();
        }
    })
}


function info(i) {
    self.location = 'static/info.html#' + i;
}

function limitSpeed(port) {

    BootstrapDialog.show({
        title: '限制速度',
        message: $('<div></div>').load("static/speed.html"),
        buttons: [{
            label: '确认',
            action: function (dialogRef) {
                var type = $("#type").val();
                var speed = $("#speed").val();
                if (speed === '') {
                    speed = 0;
                }
                var fromDate = new FormData();
                fromDate.append("port", port);
                fromDate.append("speed", speed + type);
                http.postAjax_clean("/v1/setSpeed", fromDate, function (resdate) {
                    if (resdate.state === "ok") {
                        window.location.reload();
                        dialogRef.close();
                    }
                })
            }
        }]
    });
}