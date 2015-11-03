Highcharts.setOptions({
    global: {
        useUTC: false
    },
    exporting: {
        enabled: false
    },
    rangeSelector: {
        buttons: [{
            count: 1,
            type: 'month',
            text: '1M'
        }, {
            count: 5,
            type: 'month',
            text: '5M'
        }, {
            type: 'all',
            text: 'Alles'
        }],
        inputEnabled: false,
        selected: 0
    },
    lang: {
        decimalPoint: ',',
        thousandsSep: '.',
        loading: 'Daten werden geladen...',
        months: ['Januar', 'Februar', 'März', 'April', 'Mai', 'Juni', 'Juli', 'August', 'September', 'Oktober', 'November', 'Dezember'],
        weekdays: ['Sonntag', 'Montag', 'Dienstag', 'Mittwoch', 'Donnerstag', 'Freitag', 'Samstag'],
        shortMonths: ['Jan', 'Feb', 'Mär', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Okt', 'Nov', 'Dez'],
        exportButtonTitle: "Exportieren",
        printButtonTitle: "Drucken",
        rangeSelectorFrom: "Von",
        rangeSelectorTo: "Bis",
        rangeSelectorZoom: "Zeitraum",
        downloadPNG: 'Download als PNG-Bild',
        downloadJPEG: 'Download als JPEG-Bild',
        downloadPDF: 'Download als PDF-Dokument',
        downloadSVG: 'Download als SVG-Bild',
        resetZoom: "Zoom zurücksetzen",
        resetZoomTitle: "Zoom zurücksetzen"
    }
});


var options = {
    colors: ['#7cb5ec', '#434348', '#90ed7d', '#f7a35c', '#8085e9', '#f15c80', '#e4d354', '#8085e8', '#8d4653', '#91e8e1'],
    chart: {
        renderTo: 'content',
        zoomType: 'xy',
    },
    rangeSelector: {
        enabled: true,
        buttons: [{
            type: 'hour',
            count: 1,
            text: '1h'
        }, {
            type: 'month',
            count: 3,
            text: '3m'
        }, {
            type: 'month',
            count: 6,
            text: '6m'
        }, {
            type: 'ytd',
            text: 'YTD'
        }, {
            type: 'year',
            count: 1,
            text: '1y'
        }, {
            type: 'all',
            text: 'All'
        }]
    },
    title: {
        text: ''
    },
    xAxis: {
        type: 'datetime',
        plotBands: {{ .plotBands }},
        dateTimeLabelFormats: {
            hour: '%H:%M',
        }
    },
    yAxis: [{ // Primary yAxis
        title: {
            text: 'Luftfeuchtigkeit',
            style: {
                color: Highcharts.getOptions().colors[0]
            }
        },
        opposite: false,
        labels: {
            format: '{value} %',
            style: {
                color: Highcharts.getOptions().colors[0]
            }
        }

    }, { // Secondary yAxis
        gridLineWidth: 0,
        title: {
            text: 'Temperatur',
            style: {
                color: Highcharts.getOptions().colors[1]
            }
        },
        opposite: true,
        labels: {

            format: '{value}°C',
            style: {
                color: Highcharts.getOptions().colors[1]
            }
        }
    }, { // Tertiary yAxis
        gridLineWidth: 0,
        title: {
            text: 'Luftdruck auf Meereshöhe',
            style: {
                color: Highcharts.getOptions().colors[2]
            }
        },
        labels: {
            format: '{value} hPa',
            style: {
                color: Highcharts.getOptions().colors[2]
            }
        },
        opposite: true
    }],
    tooltip: {
        shared: true
    },
    legend: {
    },
    series: [{
        name: 'Luftfeuchtigkeit',
        type: 'spline',
        yAxis: 0,
        color: Highcharts.getOptions().colors[0],
        marker: {
            enabled: false
        },
        data: [ {{ $temp := .humidity }} {{ if eq .type "single" }}
                {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.Value}}], {{end}}
                {{ else }}
                {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.AvgValue}}], {{end}}
                {{end}}
        ],
        tooltip: {
            valueDecimals: 1,
            valueSuffix: ' %'
        }
    }, {{ if eq .type "min-max" }} {
            name: 'Range',
            data: [ {{ $temp := .humidity }}
                    {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.MinValue}}, {{$value.MaxValue}}], {{end}}
            ],
            type: 'arearange',
            yAxis: 0,
            lineWidth: 0,
            linkedTo: ':previous',
            color: Highcharts.getOptions().colors[0],
            fillOpacity: 0.3,
            zIndex: 0
    }, {{end}} {
        name: 'Temperatur',
        type: 'spline',
        yAxis: 1,
        color: Highcharts.getOptions().colors[1],
        marker: {
            enabled: false
        },
        data: [ {{ $temp := .tmp_dth22 }} {{ if eq .type "single" }}
                {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.Value}}], {{end}} {{ else }}
                {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.AvgValue}}], {{end}}
                {{end}}
        ],
        tooltip: {
            valueDecimals: 1,
            valueSuffix: ' °C'
        }
    }, {{ if eq .type "min-max" }} {
            name: 'Range',
            data: [ {{ $temp := .tmp_dth22 }}
                    {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.MinValue}}, {{$value.MaxValue}}], {{end}}
            ],
            type: 'arearange',
            yAxis: 1,
            lineWidth: 0,
            linkedTo: ':previous',
            color: Highcharts.getOptions().colors[1],
            fillOpacity: 0.3,
            zIndex: 0
    }, {{end}} {
        name: 'Luftdruck auf Meereshöhe',
        type: 'spline',
        visible: false,
        yAxis: 2,
        color: Highcharts.getOptions().colors[2],
        data: [ {{ $temp := .p_sea }}  {{ if eq .type "single" }}
                {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.Value}}], {{end}} {{ else }}
                {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.AvgValue}}], {{end}}
                {{end}}
        ],
        marker: {
            enabled: false
        },
        //dashStyle: 'shortdot',
        tooltip: {
            valueDecimals: 2,
            valueSuffix: ' hPa'
        }
    }{{ if eq .type "min-max" }}, {
            name: 'Range',
            data: [ {{ $temp := .p_sea }}
                    {{ range $n, $value := $temp }}[{{$value.Timestamp}}000, {{$value.MinValue}}, {{$value.MaxValue}}], {{end}}
            ],
            type: 'arearange',
            yAxis: 2,
            lineWidth: 0,
            linkedTo: ':previous',
            color: Highcharts.getOptions().colors[2],
            fillOpacity: 0.3,
            zIndex: 0
    }{{end}}]
};

$(function() {
    $('#content').highcharts(options);
});
