$(document).ready(function () {
    var groupColumn = 0;

    var tableDevices = $('#devices table')
        .DataTable({
            pageLength: 50,
            language: {
                url: '/dashboard/datatables/i18n.json'
            },
            ajax: {
                url: '/boggart/manager/?entity=devices',
                dataSrc: 'data'
            },
            order: [[groupColumn, 'asc']],
            columnDefs: [{
                visible: false,
                targets: groupColumn
            }],
            columns: [
                {
                    data: 'type'
                },
                {
                    data: 'tags',
                    render: function (tags) {
                        var content = '';

                        for (var i in tags) {
                            content += '<span class="label label-success">' + tags[i] + '</span> ';
                        }

                        return content;
                    }
                },
                {
                    data: 'status',
                    render: function (status) {
                        switch (status.toLowerCase()) {
                            case 'online':
                                return '<span class="label label-success">' + status + '</span>';

                            case 'offline':
                                return '<span class="label label-danger">' + status + '</span>';

                            case 'removing':
                                return '<span class="label label-info">' + status + '</span>';

                            case 'removed':
                                return '<span class="label label-warning">' + status + '</span>';

                            default:
                                return '<span class="label label-default">' + status + '</span>';
                        }
                    }
                },
                {
                    data: 'tasks',
                    render: function (tasks) {
                        return tasks.length;
                    }
                },
                {
                    data: 'mqtt_publishes',
                    render: function (publishes) {
                        return publishes.length;
                    }
                },
                {
                    data: 'mqtt_subscribers',
                    render: function (subscribers) {
                        return subscribers.length;
                    }
                },
                {
                    data: null,
                    render: function (data, type, row) {
                        var content = '<div class="btn-group" role="group">';

                        if (row.has_widget) {
                            content += '<a href="/boggart/widget/' + row.id + '/" class="btn btn-info btn-icon btn-xs">' +
                                '<i class="glyphicon glyphicon-new-window" title="Open widget"></i>' +
                                '</a>';
                        }

                        return content + '<a href="/boggart/bind/' + row.id + '/" class="btn btn-warning btn-icon btn-xs">' +
                            '<i class="glyphicon glyphicon-edit" title="Edit bind"></i>' +
                            '</a>' +
                            '<button type="button" class="btn btn-primary btn-icon btn-xs" onclick="reloadConfig(\'' + row.id + '\');">' +
                            '<i class="glyphicon glyphicon-upload" title="Reload from config file"></i>' +
                            '</a>' +
                            '<button type="button" class="btn btn-danger btn-icon btn-xs" data-toggle="modal" data-target="#modal" data-modal-title="Confirm unregister device #' + row.id + '" data-modal-callback="bindUnregister(\'' + row.id + '\');">' +
                            '<i class="glyphicon glyphicon-trash" title="Unregister bind"></i>' +
                            '</button>' +
                            '</div>'
                    }
                },
                {
                    data: 'serial_number'
                },
                {
                    data: 'id'
                },
                {
                    data: 'description'
                },
                {
                    data: 'tasks',
                    render: function (tasks) {
                        var content = '';

                        for (var i in tasks) {
                            if (i > 0) {
                                content += '<br />';
                            }

                            content += '<span class="label label-warning">' + tasks[i] + '</span>';
                        }

                        return content;
                    }
                },
                {
                    data: 'mqtt_publishes',
                    render: function (publishes) {
                        var content = '';

                        for (var i in publishes) {
                            if (i > 0) {
                                content += '<br />';
                            }

                            content += '<span class="label label-primary">' + publishes[i] + '</span>';
                        }

                        return content;
                    }
                },
                {
                    data: 'mqtt_subscribers',
                    render: function (subscribers) {
                        var content = '';

                        for (var i in subscribers) {
                            if (i > 0) {
                                content += '<br />';
                            }

                            content += '<span class="label label-info">' + subscribers[i] + '</span>';
                        }

                        return content;
                    }
                },
                {
                    data: 'config',
                    render: function (config) {
                        return '<pre><code class="yaml">' + config + '</code></pre>';
                    }
                }
            ],
            'drawCallback': function (settings) {
                var api = this.api();
                var rows = api.rows({page: 'current'}).nodes();
                var last = null;

                api.column(groupColumn, {page: 'current'}).data().each(function (group, i) {
                    var $aRow = $(rows).eq(i);

                    if (last !== group) {
                        $aRow.before('<tr class="group"></td><td colspan="' + $aRow.children().length + '">' + group + '</td></tr>');
                        last = group;
                    }
                });
            }
        });
    tableDevices.on('click', 'tr.group', function () {
        var currentOrder = tableDevices.order()[0];

        if (currentOrder[0] === groupColumn && currentOrder[1] === 'asc') {
            tableDevices.order([groupColumn, 'desc']).draw();
        }
        else {
            tableDevices.order([groupColumn, 'asc']).draw();
        }
    });

    var tableListeners = $('#listeners table')
        .DataTable({
            language: {
                url: '/dashboard/datatables/i18n.json'
            },
            ajax: {
                url: '/boggart/manager/?entity=listeners',
                dataSrc: 'data'
            },
            columns: [
                {
                    data: 'name'
                },
                {
                    data: 'id'
                },
                {
                    data: 'fires'
                },
                {
                    data: 'fire_first',
                    render: function (date) {
                        if (!date) {
                            return '';
                        }

                        return dateToString(date);
                    }
                },
                {
                    data: 'fire_last',
                    render: function (date) {
                        if (!date) {
                            return '';
                        }

                        return dateToString(date);
                    }
                },
                {
                    data: 'events',
                    render: function (data) {
                        var content = '';

                        for (var eventId in data) {
                            content += '<span class="label label-info">' + data[eventId] + '</span> ';
                        }

                        return content;
                    }
                }
            ]
        });

    window.bindUnregister = function (id) {
        $.ajax({
            type: 'POST',
            url: '/boggart/bind/' + id + '/unregister/',
            success: function (r) {
                if (r.result === 'failed') {
                    new PNotify({
                        title: 'Error',
                        text: r.message,
                        type: 'error',
                        hide: false,
                        styling: 'bootstrap3'
                    });
                    return
                }

                refreshTables();
            }
        });
    };

    window.reloadConfig = function (id) {
        var url = '/boggart/config/reload';

        if (typeof id !== 'undefined') {
            url += '?id=' + id
        }

        $.ajax({
            type: 'POST',
            url: url,
            success: function (r) {
                if (r.result === 'failed') {
                    new PNotify({
                        title: 'Error',
                        text: r.message,
                        type: 'error',
                        hide: false,
                        styling: 'bootstrap3'
                    });
                } else if (r.message !== 'undefined') {
                    new PNotify({
                        title: 'Success',
                        text: r.message,
                        type: 'success',
                        hide: false,
                        styling: 'bootstrap3'
                    });
                }

                refreshTables();
            }
        });
    };

    window.refreshTables = function () {
        tableDevices.ajax.reload();
        tableListeners.ajax.reload();
    };
});