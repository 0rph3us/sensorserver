#!/bin/bash

mkdir -p assets/js/

echo 'Download jquery'
curl --silent 'http://code.jquery.com/jquery-2.1.4.min.js' > assets/js/jquery.min.js

echo 'Download highcharts'
curl --silent 'https://code.highcharts.com/highcharts.js' > assets/js/highcharts.js
curl --silent 'https://code.highcharts.com/highcharts-more.js' > assets/js/highcharts-more.js

