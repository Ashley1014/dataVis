$(document).ready(function() {
    $("#gender_stats").click(function() {
        let female = (document.getElementById("female")).value
        let male = (document.getElementById("male")).value
        let myChart = echarts.init(document.getElementById("main"));
        let option = {
            series: [
                {
                    name: 'gender stats',
                    type: 'pie',
                    radius: '55%',
                    data: [
                        {value: female, name: 'female'},
                        {value: male, name: 'male'},
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
                    name: 'age stats',
                    type: 'pie',
                    radius: '55%',
                    data: [
                        {value: teen, name: 'teenager'},
                        {value: adult, name: 'adult'},
                        {value: elderly, name: 'elderly'}
                    ]
                }
            ]
        };
        myChart.setOption(option)
    })
})
