$(document).ready(function() {
    $("#gender_stats").click(function() {
        let female = (document.getElementById("female")).value
        let male = (document.getElementById("male")).value
        let myChart = echarts.init(document.getElementById("main"));
        let option = {
            series: [
                {
                    name: '性别统计',
                    type: 'pie',
                    radius: '55%',
                    data: [
                        {value: female, name: '女性'},
                        {value: male, name: '男性'},
                    ]
                }
            ]
        };
        myChart.setOption(option)
    })

    $("#age_stats").click(function() {
        let myChart = echarts.init(document.getElementById("side"));
        let teen = (document.getElementById("teen")).value
        let adult = (document.getElementById("adult")).value
        let elderly = (document.getElementById("elderly")).value
        let option = {
            series: [
                {
                    name: '年龄统计',
                    type: 'pie',
                    radius: '55%',
                    data: [
                        {value: teen, name: '青少年'},
                        {value: adult, name: '成年人'},
                        {value: elderly, name: '老年人'}
                    ]
                }
            ]
        };
        myChart.setOption(option)
    })
})
